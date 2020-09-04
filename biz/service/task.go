package service

import (
	"context"
	"time"

	bproto "git.supremind.info/testplatform/biz/proto"
	"git.supremind.info/testplatform/biz/service/db"
	"git.supremind.info/testplatform/biz/service/proto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TaskHandlerImp interface {
	InsertTaskJsonHandler(c *gin.Context)
	DeleteTaskJsonHandler(c *gin.Context)
	UpdateTaskJsonHandler(c *gin.Context)
	FindTaskJsonHandler(c *gin.Context)
	LikeFindTaskJsonHandler(c *gin.Context)
	BatchDeleteTaskJsonHandler(c *gin.Context)

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

type TaskHandlerRouter struct {
	task      db.TaskMgnt
	TaskHost  *TargetHost
	eventHost *TargetHost
	client    *Client
}

var k TaskHandlerRouter

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
func (this *TaskHandlerRouter) DeleteTaskJsonHandler(c *gin.Context) {

}
func (this *TaskHandlerRouter) UpdateTaskJsonHandler(c *gin.Context) {

}
func (this *TaskHandlerRouter) FindTaskJsonHandler(c *gin.Context) {

}
func (this *TaskHandlerRouter) LikeFindTaskJsonHandler(c *gin.Context) {

}
func (this *TaskHandlerRouter) BatchDeleteTaskJsonHandler(c *gin.Context) {
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
// @Success 200 {object}  proto.CommonTaskRes{data=[]proto.TaskInfo}
// @Router /v1/task/ [get]
func (this *TaskHandlerRouter) GetTaskListHandler(c *gin.Context) {
	Forward(c, k.TaskHost)
}

// @Summary 创建布控任务
// @Description 新增一个布控任务到VMR中
// @Accept json
// @Param example body proto.CreateTaskReq true "CreateTaskReq"
// @Success 200 {object}  proto.CommonTaskRes{}
// @Router /v1/task/ [post]
func (this *TaskHandlerRouter) CreateTaskHandler(c *gin.Context) {
	Forward(c, k.TaskHost)
}

// @Summary 获取单个布控任务
// @Description 根据id来获取单个布控任务
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.TaskInfo
// @Router /v1/task/{id} [get]
func (this *TaskHandlerRouter) GetTaskHandler(c *gin.Context) {
	Forward(c, k.TaskHost)
}

// @Summary 删除单个布控任务
// @Description 根据id来删除单个布控任务
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonTaskRes{}
// @Router /v1/task/del/{id} [post]
func (this *TaskHandlerRouter) DelTaskHandler(c *gin.Context) {
	Forward(c, k.TaskHost)
}

// @Summary 启动布控任务
// @Description 根据id来启动单个布控任务
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonTaskRes{}
// @Router /v1/task/start/{id} [post]
func (this *TaskHandlerRouter) StartTaskHandler(c *gin.Context) {
	Forward(c, k.TaskHost)
}

// @Summary 停止布控任务
// @Description 根据id来停止布控任务
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonTaskRes{}
// @Router /v1/task/stop/{id} [post]
func (this *TaskHandlerRouter) StopTaskHandler(c *gin.Context) {
	Forward(c, k.TaskHost)
}

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

func TaskJsonHandler(group *gin.RouterGroup) {

	group.POST("/taskjson", k.InsertTaskJsonHandler)
	group.DELETE("/taskjson/:id", k.DeleteTaskJsonHandler)
	group.DELETE("taskjson", k.BatchDeleteTaskJsonHandler)
	group.PUT("/taskjson/:id", k.UpdateTaskJsonHandler)
	group.GET("/taskjson/:id", k.FindTaskJsonHandler)
	group.GET("/taskjson", k.LikeFindTaskJsonHandler)

}

func TaskHandler(group *gin.RouterGroup) {

	group.Use(MiddleWare(k.client))
	group.GET("/task", k.GetTaskListHandler)
	group.POST("/task", k.CreateTaskHandler)
	group.GET("/task/:id", k.GetTaskHandler)
	group.POST("/task/del/:id", k.DelTaskHandler)
	group.POST("/task/start/:id", k.StartTaskHandler)
	group.POST("/task/stop/:id", k.StopTaskHandler)
	group.POST("/task/update/:id", k.UpdateTaskHandler)

}

func EventHandler(group *gin.RouterGroup) {
	group.Use(MiddleWare(k.client))
	group.GET("/jt_event/events", k.GetEventHandler)
}

func TaskHandlerSvc(ctx context.Context, task db.TaskMgnt, group *gin.RouterGroup, conf VMRClient) {

	k.task = task
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
	k.client = &Client{
		Username: conf.Username,
		Password: conf.Password,
	}
	go TaskJsonHandler(group)
	go TaskHandler(group)
	go EventHandler(group)
}
