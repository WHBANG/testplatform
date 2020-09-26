package service

import (
	"context"
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"git.supremind.info/product/visionmind/com/flow"
	client "git.supremind.info/product/visionmind/sdk/vmr/go_sdk"
	bproto "git.supremind.info/testplatform/biz/proto"
	"git.supremind.info/testplatform/biz/service/db"
	"git.supremind.info/testplatform/biz/service/proto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	Match    string = "事件类型匹配"
	NotFound string = "未产生事件"
	NotMatch string = "事件类型不匹配"
	Failed   string = "获取任务失败"
)

type TaskHandlerImp interface {
	/*
		InsertTaskJsonHandler(c *gin.Context)
		DeleteTaskJsonHandler(c *gin.Context)
		UpdateTaskJsonHandler(c *gin.Context)
		FindTaskJsonHandler(c *gin.Context)
		LikeFindTaskJsonHandler(c *gin.Context)
		BatchDeleteTaskJsonHandler(c *gin.Context)
	*/
	//获取布控列表
	GetTaskListHandler(c *gin.Context)
	//创建布控任务
	CreateTaskHandler(c *gin.Context)
	//获取单个布控任务
	GetTaskHandler(c *gin.Context)
	//删除单个布控任务
	DelTaskHandler(c *gin.Context)
	//启动布控任务
	StartTaskHandler(c *gin.Context)
	//停止布控任务
	StopTaskHandler(c *gin.Context)
	//更新单个布控任务
	UpdateTaskHandler(c *gin.Context)
}

type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type TaskConf struct {
	NamePrefix     string `json:"name_prefix"`
	GlobalDeviceID string `json:"global_device_id"`
	MaxChannel     int    `json:"max_channel"`
}

type TaskHandlerRouter struct {
	enginetask   db.EngineTaskMgnt
	testcaseMgnt db.TestCaseMgnt
	TaskHost     *TargetHost
	eventHost    *TargetHost
	Client       *Client
	Conf         *TaskConf
}

var k TaskHandlerRouter

// @Summary 创建布控任务
// @Description 新增一个布控任务到VMR中
// @Accept json
// @Param example body proto.GetCreateTaskReq true "GetCreateTaskReq"
// @Success 200 {object}  proto.CommonRes{data=EngineTaskInfo}
// @Router /v1/task/create/{region} [post]
func (this *TaskHandlerRouter) CreateTaskHandler(c *gin.Context) {

	taskReq := proto.GetCreateTaskReq{}
	err := c.BindJSON(&taskReq)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	imgUrls := taskReq.TestData.Files.Images
	videoUrls := taskReq.TestData.Files.Videos
	caseMeta := bproto.MetaData{}
	caseMeta.Case = &taskReq.TestData.Case
	caseMeta.Event = taskReq.TestData.Event
	caseMeta.Files = &taskReq.TestData.Files
	caseMeta.Task = &taskReq.TestData.Task
	region := c.Param("region")

	taskID, err := this.CreateTask(&caseMeta, imgUrls, videoUrls, region)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	//插入数据库
	en := bproto.EngineTaskInfo{}
	en.TaskID = taskID
	testcase_id := taskReq.TestCaseID.String()
	start := strings.Index(testcase_id, "(")
	//parts := strings.SplitN(taskReq.TestCaseID.String(), '"', 3)
	en.TestCaseID = testcase_id[start+2 : start+24]
	en.EngineID = region
	en.Status = "on"
	en.TaskStatuss = bproto.TaskStopped
	data, err := k.enginetask.Insert(&en)
	proto.DefaultRet(c, data, err)

}

// 创建任务
func (this *TaskHandlerRouter) CreateTask(caseMeta *bproto.MetaData, imgUrls, videoUrls []string, region string) (id string, err error) {

	id = "-999"
	if len(imgUrls) < 1 || len(videoUrls) < 1 {
		return id, errors.New("截图或视频数量小于1")
	}
	var (
		retryTimes              = 0
		device                  *flow.Device
		videoURLs, snapshotURLs []string
	)
	for _, video := range videoUrls {
		videoRawUrl, err := url.Parse(video)
		if err != nil {
			return id, err
		}
		videoRawUrl.RawQuery = videoRawUrl.Query().Encode()
		videoURLs = append(videoURLs, videoRawUrl.String())
	}
	for _, img := range imgUrls {
		snapshotRawUrl, err := url.Parse(img)
		if err != nil {
			return id, err
		}
		snapshotRawUrl.RawQuery = snapshotRawUrl.Query().Encode()
		snapshotURLs = append(snapshotURLs, snapshotRawUrl.String())
	}

	for retryTimes < 3 {
		flowClient := client.NewFlowClient(this.TaskHost.Host)
		device, err = CreateSubDevice(flowClient, caseMeta.Case.Name, videoURLs[0], this.Conf.MaxChannel,
			this.Conf.GlobalDeviceID, this.Conf.NamePrefix)
		if err == nil {
			break
		}
		retryTimes++
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return
	}

	retryTimes = 0
	for retryTimes < 3 {
		id, err = this.createTask(caseMeta.Task, device.ID, snapshotURLs[0], region)
		if err == nil {
			break
		}
		retryTimes++
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return
	}

	return id, nil
}

func (this *TaskHandlerRouter) createTask(task *flow.Task, streamID, snapshotURL, region string) (id string, err error) {
	id = "-999"
	if task == nil {
		err = errors.New("task is null")
		log.Errorf("createTask: %+v", err)
		return
	}
	parts := strings.SplitN(this.Conf.GlobalDeviceID, ".", 2)

	if len(parts) < 2 {
		//err = errors.New("invalid vms device")
		//return
		streamID = parts[0] + "." + streamID
	} else {
		streamID = parts[0] + "." + streamID
	}

	task.ID = this.Conf.NamePrefix + "_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	task.Name += "_" + this.Conf.NamePrefix + strconv.FormatInt(time.Now().UnixNano(), 10)
	task.StreamID = streamID
	task.Snapshot = snapshotURL
	task.StreamON = "ON"
	task.AnalyzeConfig["enable_tracking_debug"] = true
	task.Region = region
	violations, ok := task.AnalyzeConfig["violations"].([]interface{})
	if !ok {
		return id, errors.New("case.json文件格式错误")
	}
	for i, _ := range violations {
		if v, ok := violations[i].(map[string]interface{}); ok {
			v["on"] = true
		}
	}

	task.AnalyzeConfig["tracking_threshold"] = 0.6
	taskClient := client.NewTaskMgmClient(this.TaskHost.Host)
	err = taskClient.CreateTask(task)
	//err = go_sdk.ITaskMgmClient.CreateTask
	if err != nil {
		log.Errorf("t.TaskClient.CreateTask(%+v):%+v", task, err)
		return
	}
	return task.ID, nil
}

/*
func (this *TaskHandlerRouter) InsertTaskJsonHandler(c *gin.Context) {
	var insertReq bproto.CreateTaskReq
	err := c.BindJSON(&insertReq)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	tc := &bproto.TaskInfo{
		ID:            insertReq.ID,
		Type:          insertReq.Type,
		Region:        insertReq.Region,
		StreamList:    insertReq.StreamList,
		WorkField:     insertReq.WorkField,
		Status:        insertReq.Status,
		UpdatedAt:     time.Now(),
		Extra:         insertReq.Extra,
		AnalyzeConfig: insertReq.AnalyzeConfig,
		Name:          insertReq.Name,
		Snapshot:      insertReq.Snapshot,
		StreamID:      insertReq.StreamID,
		StreamON:      insertReq.StreamON,
		CreatedAt:     time.Now(),
	}
	data, err := k.task.Insert(tc)
	if err != nil {
		log.Fatal("insert error!", err)
		proto.DefaultRet(c, nil, err)
	} else {
		proto.DefaultRet(c, data, nil)
	}
}
*/
func (this *TaskHandlerRouter) GetEngineTaskListHandler(c *gin.Context) {

	engineID := c.Param("engine_id")
	log.Println(engineID)
	enginetasks, err := this.enginetask.FindAllByEngineId(engineID)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	size := len(enginetasks)
	log.Println(size)
	if size == 0 {
		err = errors.New("Not Data")
		proto.DefaultRet(c, nil, nil)
		return
	}

	tasks := make([]proto.TaskListReq, size)
	for i, v := range enginetasks {
		tasks[i].TaskID = v.TaskID
		tasks[i].TaskStatus = v.TaskStatuss
	}
	proto.DefaultRet(c, tasks, nil)
}

// @Summary 布控任务查询
// @Description 根据输入的task字段信息来查询布控任务信息
// @Accept json
// @Param id query string false "id"
// @Param type query string false "type"
// @Param violation_type query string false "violation_type"
// @Param stream_id query string false "stream_id"
// @Param status query string false "status"
// @Param stream_on query string false "stream_on"
// @Param create_time_begin query string false "create_time_begin"
// @Param create_time_end query string false "create_time_end"
// @Param name query string false "name"
// @Param simple query string false "simple"
// @Param ids query []string false "ids"
// @Param org_codes query []string false "org_codes"
// @Param device_detail query []string false "device_detail"
// @Success 200 {object}  proto.CommonRes{data=[]proto.TaskInfo}
// @Router /v1/task/ [get]
func (this *TaskHandlerRouter) GetTaskListHandler(c *gin.Context) {
	Forward(c, k.TaskHost)
}

// @Summary 获取单个布控任务
// @Description 根据id来获取单个布控任务
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.TaskInfo}
// @Router /v1/task/{id} [get]
func (this *TaskHandlerRouter) GetTaskHandler(c *gin.Context) {
	//Forward(c, k.TaskHost)
	//localhost:7000/v1/task?id=test_wh_1600931300677560000
	id := c.Query("id")
	param := "v1/task?id=" + id
	data, err := Forward01(k.TaskHost.Host, "get", param)
	proto.DefaultRet(c, data, err)
}

// @Summary 删除单个布控任务
// @Description 根据id来删除单个布控任务
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=interface{}}
// @Router /v1/task/del/{id} [post]
func (this *TaskHandlerRouter) DelTaskHandler(c *gin.Context) {
	//Forward(c, k.TaskHost)
	id := c.Param("id")
	param := "v1/task/del/" + id
	data, err := Forward01(k.TaskHost.Host, "post", param)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	err = k.enginetask.Delete(id)
	proto.DefaultRet(c, data, err)
}

// @Summary 启动布控任务
// @Description 根据id来启动单个布控任务
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=interface{}}
// @Router /v1/task/start/{id} [post]
func (this *TaskHandlerRouter) StartTaskHandler(c *gin.Context) {

	id := c.Param("id")
	param := "v1/task/start/" + id
	data, err := Forward01(k.TaskHost.Host, "post", param)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	info := make(map[string]interface{})
	info["task_status"] = string(bproto.TaskRunning)
	err = k.enginetask.UpdateStatus(id, info)
	proto.DefaultRet(c, data, err)
}

// @Summary 停止布控任务
// @Description 根据id来停止布控任务
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{interface{}}
// @Router /v1/task/stop/{id} [post]
func (this *TaskHandlerRouter) StopTaskHandler(c *gin.Context) {
	//Forward01(c, k.TaskHost)
	id := c.Param("id")
	param := "v1/task/stop/" + id
	data, err := Forward01(k.TaskHost.Host, "post", param)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	info := make(map[string]interface{})
	info["task_status"] = string(bproto.TaskStopped)
	err = k.enginetask.UpdateStatus(id, info)
	proto.DefaultRet(c, data, err)
}

/*
// func (this *TaskHandlerRouter) StopTaskHandler(c *gin.Context) {
// 	Forward(c, k.TaskHost)
// }
*/

// @Summary 更新布控任务
// @Description 根据id来更新布控任务
// @Accept json
// @Param id path string true "id"
// @Param example body proto.UpdateTaskReq true "UpdateTaskReq"
// @Success 200 {object}  proto.CommonRes{data=proto.TaskInfo}
// @Router /v1/task/update/{id} [post]
func (this *TaskHandlerRouter) UpdateTaskHandler(c *gin.Context) {
	Forward(c, k.TaskHost)
}

// @Summary 事件查询
// @Description 查询已存在的事件
// @Accept json
// @Param eventId query string false "eventId"
// @Param taskId query string false "taskId"
// @Param processReasons query []string false "事件原因：{391:已处理，392:忽略}"
// @Param processStatus query string false "事件状态：{0:未处理，1:已处理}"
// @Param eventTypes query []string false "eventTypes"
// @Param cameraIds query []string false "cameraIds"
// @Param __timestamp__ query string false "__timestamp__"
// @Success 200 {object}  proto.CommonRes{data=proto.GetJTEventRes}
// @Router /v1/jt_event/events [get]
func (this *TaskHandlerRouter) GetEventHandler(c *gin.Context) {
	Forward(c, k.eventHost)
}

/*
func TaskJsonHandler(group *gin.RouterGroup) {

	group.POST("/taskjson", k.InsertTaskJsonHandler)
	group.DELETE("/taskjson/:id", k.DeleteTaskJsonHandler)
	group.DELETE("taskjson", k.BatchDeleteTaskJsonHandler)
	group.PUT("/taskjson/:id", k.UpdateTaskJsonHandler)
	group.GET("/taskjson/:id", k.FindTaskJsonHandler)
	group.GET("/taskjson", k.LikeFindTaskJsonHandler)

}
*/

//获取已经存在的所有事件
func (this *TaskHandlerRouter) GetEventResultHandler() ([]bproto.EventInfo, error) {
	taskList, err := k.enginetask.GetAll()
	if err != nil {
		return nil, err
	}
	if len(taskList) == 0 {
		return nil, nil
	}

	var eventResult []bproto.EventInfo
	for _, v := range taskList {
		taskId := v.TaskID
		taskcaseId := v.TestCaseID
		var tmp bproto.EventInfo
		tmp.TestCaseID = v.EngineID
		tmp.TestCaseID = taskcaseId
		tmp.TaskID = taskId
		eventList, err := GetEventHandler(taskId)
		if err != nil {
			tmp.Result = Failed
			eventResult = append(eventResult, tmp)
			continue
		}
		if len(eventList) == 0 {
			tmp.Result = NotFound
			eventResult = append(eventResult, tmp)
			continue
		}
		for _, e := range eventList {
			result, err := GetEventResult(e, taskcaseId)
			if err != nil {
				log.Println("Get Result Error:", e.ID, "-", err)
				continue
			}
			if result == Match {
				tmp.Result = result
				tmp.EventID = e.EventID
				eventResult = append(eventResult, tmp)
				break
			}

		}
		tmp.Result = NotMatch
		eventResult = append(eventResult, tmp)
	}
	return eventResult, nil
}

func GetEventHandler(taskID string) ([]proto.JTEventInfo, error) {
	//localhost:7000/v1/jt_event/events?taskId=jtsj_1598493538728
	param := "v1/jt_event?taskId=" + taskID
	data, err := Forward02(k.eventHost.Host, "get", param)
	return data, err
}

//事件比较
func GetEventResult(data proto.JTEventInfo, tastcaseID string) (string, error) {
	id := bson.ObjectIdHex(tastcaseID)
	testcase, err := k.testcaseMgnt.Find(id)
	if err != nil {
		return Failed, err
	}
	eventData := testcase.TestData.Event
	for _, v := range eventData {
		if v["eventType"] == data.EventType {
			return Match, nil
		}
	}
	return NotMatch, err
}

func EngineTaskHandler(group *gin.RouterGroup) {

	group.GET("/task/group/:engine_id", k.GetEngineTaskListHandler)

}

func TaskHandler(group *gin.RouterGroup) {

	group.Use(MiddleWare(k.Client))
	//group.GET("/task", k.GetTaskListHandler)
	group.POST("/task/create/:region", k.CreateTaskHandler)
	group.GET("/task", k.GetTaskHandler)
	group.POST("/task/del/:id", k.DelTaskHandler)
	group.POST("/task/start/:id", k.StartTaskHandler)
	group.POST("/task/stop/:id", k.StopTaskHandler)
	group.POST("/task/update/:id", k.UpdateTaskHandler)

}

func EventHandler(group *gin.RouterGroup) {
	group.Use(MiddleWare(k.Client))
	group.GET("/jt_event/events", k.GetEventHandler)
}

func TaskHandlerSvc(ctx context.Context, task db.EngineTaskMgnt, testcase db.TestCaseMgnt, group *gin.RouterGroup, conf VMRClient) {

	k.enginetask = task
	k.testcaseMgnt = testcase
	k.TaskHost = &TargetHost{
		Host:    conf.FlowHost,
		IsHTTPS: conf.IsHTTPS,
		CAPath:  conf.CAPath,
	}
	k.eventHost = &TargetHost{
		Host:    conf.WebGeneral,
		IsHTTPS: conf.IsHTTPS,
		CAPath:  conf.CAPath,
	}
	k.Client = &Client{
		Username: conf.Username,
		Password: conf.Password,
	}
	k.Conf = &TaskConf{
		NamePrefix:     conf.NamePrefix,
		GlobalDeviceID: conf.GlobalDeviceID,
		MaxChannel:     conf.MaxChannel,
	}
	go EngineTaskHandler(group)
	go TaskHandler(group)
	go EventHandler(group)
}
