package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"git.supremind.info/testplatform/biz/analyzerclient"
	bproto "git.supremind.info/testplatform/biz/proto"
	"git.supremind.info/testplatform/biz/service/db"
	"git.supremind.info/testplatform/biz/service/proto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

type EngineTestSvc struct {
	testMgnt      db.TestMgnt
	engineIfc     analyzerclient.AnalyzerInterface
	configFileMap map[string]string
}

//todo add goroutine to check job status

func NewEngineTestSvc(ctx context.Context, group *gin.RouterGroup, engineIfc analyzerclient.AnalyzerInterface, configFileMap map[string]string) (*EngineTestSvc, error) {
	session, err := db.GetMgoDBSession()
	if err != nil {
		return nil, err
	}
	testMgnt, err := db.NewMongoTest(session, db.GetDBName())
	if err != nil {
		return nil, err
	}
	svc := &EngineTestSvc{
		testMgnt:      testMgnt,
		engineIfc:     engineIfc,
		configFileMap: configFileMap,
	}

	group.GET("/engine/:id", svc.GetEngine)
	group.GET("/engine", svc.GetEngineList)
	group.POST("/engine", svc.CreateEngine)
	group.PUT("/engine/:id", svc.UpdateEngine)
	group.DELETE("/engine/:id", svc.RemoveEngine)
	group.POST("/engine/:id/start", svc.StartEngine)
	group.POST("/engine/:id/stop", svc.StopEngine)

	group.POST("/test", svc.CreateTest)
	group.POST("/test/start", svc.StartTest) //启动测试
	group.POST("/test/stop", svc.StopTest)   //停止测试
	group.POST("/test/clean", svc.CleanTest) //清除

	group.DELETE("/engine", svc.BatchRemoveEngine)

	return svc, nil
}

func parseMapFromStr(strData string) (map[string]interface{}, error) {
	var data = make(map[string]interface{})
	err := json.Unmarshal([]byte(strData), &data)
	if err != nil {
		log.Error("parse input error:", err)
		return nil, err
	}
	return data, nil
}

func parseMapFromItf(itf interface{}) (map[string]interface{}, error) {
	bData, err := json.Marshal(itf)
	if err != nil {
		log.Error("parse input error:", err)
		return nil, err
	}
	var data = make(map[string]interface{})
	err = json.Unmarshal(bData, &data)
	if err != nil {
		log.Error("parse input error:", err)
		return nil, err
	}
	return data, nil
}

// @Summary 创建测试引擎
// @Description 创建测试引擎
// @Accept json
// @Param example body proto.CreateEngineReq true "CreateEngineReq"
// @Success 200 {object}  proto.CommonRes{data=proto.EngineDeployInfo}
// @Router /v1/engine [post]
func (s *EngineTestSvc) CreateEngine(c *gin.Context) {
	var req proto.CreateEngineReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		proto.DefaultRet(c, nil, err)
		return
	}
	// log.Info(s.configFileMap)
	newID := bson.NewObjectId()
	var analyzerConfig map[string]interface{}
	if req.AnalyzerConfig != "" {
		analyzerConfig, err = parseMapFromStr(req.AnalyzerConfig)
		if err != nil {
			err = fmt.Errorf("parse analyzer config file error:%s", err.Error())
			proto.DefaultRet(c, nil, err)
			return
		}
		if len(analyzerConfig) > 0 {
			analyzerConfig["region"] = newID
		}
	}
	if len(analyzerConfig) == 0 {
		if analyzerConfigStr, ok := s.configFileMap["analyzer.conf"]; ok {
			analyzerConfig, err = parseMapFromStr(analyzerConfigStr)
			if err != nil {
				err = fmt.Errorf("parse default analyzer config file error:%s", err.Error())
				proto.DefaultRet(c, nil, err)
				return
			}
			analyzerConfig["region"] = newID
		}
	}

	analyzerConfigStr, _ := json.Marshal(analyzerConfig)

	info, err := s.testMgnt.CreateEngine(&bproto.EngineDeployInfo{
		ID:             newID,
		Image:          req.Image,
		UserID:         req.UserID,
		Description:    req.Description,
		Product:        req.Description,
		Status:         bproto.EngineStatusNone,
		AnalyzerConfig: string(analyzerConfigStr),
	})
	if err != nil {
		log.Error("create engine error:", err)
	}
	proto.DefaultRet(c, *info, err)
}

// @Summary 启动引擎
// @Description 启动引擎
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.EngineDeployInfo}
// @Router /v1/engine/{id}/start [post]
func (s *EngineTestSvc) StartEngine(c *gin.Context) {
	var req proto.StartEngineReq
	var res interface{}
	var err error
	defer func() {
		proto.DefaultRet(c, res, err)
	}()
	err = c.ShouldBindUri(&req)
	if err != nil {
		log.Error(err)
		return
	}
	info, err := s.testMgnt.GetEngineOne(bson.ObjectIdHex(req.ID))
	if err != nil {
		log.Error("update mongo error:", err)
		return
	}
	configM := make(map[string]string)
	// analyzerConfig, _ := json.Marshal(info.AnalyzerConfig)
	configM["analyzer.conf"] = info.AnalyzerConfig
	jobDeployInfo := analyzerclient.AnalyzerDeployInfo{
		JobName: req.ID,
		Image:   info.Image,
	}
	if _, ok := s.configFileMap["start.sh"]; ok {
		jobDeployInfo.Args = "sh /workspace/mnt/config/start.sh"
		// jobDeployInfo.Args = "tail -f /dev/null"
		configM["start.sh"] = s.configFileMap["start.sh"]
	}

	jobDeployInfo.ConfigMap = configM
	job, err := s.engineIfc.CreateAnalyzer(&jobDeployInfo)
	if err != nil {
		log.Error("CreateAnalyzer error:", err)
		return
	}
	info, err = s.testMgnt.UpdateEngine(bson.ObjectIdHex(req.ID), bson.M{
		"status": bproto.EngineStatusCreated,
	})
	if err != nil {
		log.Error("update mongo error:", err)
		return
	}
	err = s.engineIfc.StartAnalyzer(&analyzerclient.JobInfo{
		Name:    job.Name,
		Creator: job.Creator,
	})
	if err != nil {
		log.Error("StartAnalyzer error:", err)
		return
	}
	info, err = s.testMgnt.UpdateEngine(bson.ObjectIdHex(req.ID), bson.M{
		"status":   bproto.EngineStatusStarted,
		"job_info": *job,
	})
	if err != nil {
		log.Error("update mongo error:", err)
		return
	}
	res = *info
}

// @Summary 停止引擎
// @Description 停止引擎
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.EngineDeployInfo}
// @Router /v1/engine/{id}/stop [post]
func (s *EngineTestSvc) StopEngine(c *gin.Context) {
	var req proto.StopEngineReq
	var res interface{}
	var err error
	defer func() {
		proto.DefaultRet(c, res, err)
	}()

	err = c.ShouldBindUri(&req)
	if err != nil {
		log.Error(err)
		return
	}
	info, err := s.testMgnt.GetEngineOne(bson.ObjectIdHex(req.ID))
	if err != nil {
		log.Error("update mongo error:", err)
		return
	}
	err = s.engineIfc.StopAnalyzer(&analyzerclient.JobInfo{
		Name:    info.JobInfo.Name,
		Creator: info.JobInfo.Creator,
	})
	if err != nil {
		log.Error("StopAnalyzer error:", err)
		// return
	}
	err = s.engineIfc.RemoveAnalyzer(&analyzerclient.JobInfo{
		Name:    info.JobInfo.Name,
		Creator: info.JobInfo.Creator,
	})
	if err != nil {
		log.Error("RemoveAnalyzer error:", err)
		// return
	}
	info, err = s.testMgnt.UpdateEngine(bson.ObjectIdHex(req.ID), bson.M{
		"status": bproto.EngineStatusStoped,
	})
	if err != nil {
		log.Error("update mongo error:", err)
	}
	proto.DefaultRet(c, *info, err)
}

// @Summary 更改引擎
// @Description 更改引擎
// @Accept json
// @Param id path string true "id"
// @Param example body proto.UpdateEgnineReq true "UpdateEgnineReq"
// @Success 200 {object}  proto.CommonRes{data=proto.EngineDeployInfo}
// @Router /v1/engine/{id} [put]
func (s *EngineTestSvc) UpdateEngine(c *gin.Context) {
	var reqID string
	var ok bool
	if reqID, ok = c.Params.Get("id"); !ok || !bson.IsObjectIdHex(reqID) {
		err := errors.New("invalid id specified")
		proto.DefaultRet(c, nil, err)
		return
	}

	var req proto.UpdateEgnineReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		proto.DefaultRet(c, nil, err)
		return
	}

	if req.AnalyzerConfig != "" {
		var analyzerConfig map[string]interface{}
		analyzerConfig, err = parseMapFromStr(req.AnalyzerConfig)
		if err != nil {
			err = fmt.Errorf("parse analyzer config file error:%s", err.Error())
			proto.DefaultRet(c, nil, err)
			return
		}
		if len(analyzerConfig) > 0 {
			analyzerConfig["region"] = reqID
			configStr, _ := json.Marshal(analyzerConfig)
			req.AnalyzerConfig = string(configStr)
		} else {
			req.AnalyzerConfig = ""
		}
	}
	updateStr, _ := json.Marshal(req)
	updateInfo, err := parseMapFromStr(string(updateStr))
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}

	info, err := s.testMgnt.UpdateEngine(bson.ObjectIdHex(reqID), updateInfo)
	if err != nil {
		log.Error("update mongo error:", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, *info, err)
}

// @Summary 移除引擎
// @Description 移除引擎
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.EngineDeployInfo}
// @Router /v1/engine/{id} [delete]
func (s *EngineTestSvc) RemoveEngine(c *gin.Context) {
	var req proto.RemoveEngineReq
	err := c.ShouldBindUri(&req)
	if err != nil {
		log.Error(err)
		proto.DefaultRet(c, nil, err)
		return
	}

	//todo check status

	err = s.testMgnt.RemoveEngine(bson.ObjectIdHex(req.ID))
	if err != nil {
		log.Error("remove mongo error:", err)
	}

	proto.DefaultRet(c, nil, err)
}

// @Summary 批量移除引擎
// @Description 批量移除引擎
// @Accept json
// @Param list query []string true "list" collectionFormat(multi)
// @Success 200 {object}  proto.CommonRes{data=proto.EngineDeployInfo}
// @Router /v1/engine [delete]
func (s *EngineTestSvc) BatchRemoveEngine(c *gin.Context) {
	var req []string
	req = c.QueryArray("list")
	len := len(req)
	if len == 0 {
		err := errors.New("Please select the data you want to delete!")
		proto.DefaultRet(c, nil, err)
		return
	}

	//todo check status

	list := make([]bson.ObjectId, len)
	for i, v := range req {
		list[i] = bson.ObjectIdHex(v)
	}
	err := s.testMgnt.BatchRemoveEngine(list)
	if err != nil {
		log.Error("remove mongo error:", err)
	}

	proto.DefaultRet(c, nil, err)
}

// @Summary 获取引擎
// @Description 获取单个引擎
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.EngineDeployInfo}
// @Router /v1/engine/{id} [get]
func (s *EngineTestSvc) GetEngine(c *gin.Context) {
	var reqID string
	var ok bool
	if reqID, ok = c.Params.Get("id"); !ok {
		err := errors.New("No id specified")
		proto.DefaultRet(c, nil, err)
		return
	}

	info, err := s.testMgnt.GetEngineOne(bson.ObjectIdHex(reqID))
	if err != nil {
		log.Error("get engine error:", err)
		proto.DefaultRet(c, nil, err)
		return
	}

	proto.DefaultRet(c, *info, err)
}

// @Summary 获取引擎列表
// @Description 获取引擎列表
// @Accept json
// @Param example query proto.GetEngineReq true "GetEngineReq"
// @Success 200 {object}  proto.CommonRes{data=proto.GetEngineRes}
// @Router /v1/engine [get]
func (s *EngineTestSvc) GetEngineList(c *gin.Context) {
	var req proto.GetEngineReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Error("paramerter error:", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	queryBytes, err := json.Marshal(proto.EngineQuery{
		Image:   req.Image,
		UserID:  req.UserID,
		Product: req.Product,
		Status:  req.Status,
	})
	if err != nil {
		log.Error("marshal error:", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	var query bson.M
	err = json.Unmarshal(queryBytes, &query)
	if err != nil {
		log.Error("unmarshal error:", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	dataList, num, err := s.testMgnt.GetEngine(query, req.Page, req.Size)
	if err != nil {
		log.Error("get engine list error:", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	res := proto.GetEngineRes{
		Page:  req.Page,
		Size:  req.Size,
		Total: num,
		Data:  dataList,
	}
	proto.DefaultRet(c, res, err)
}

func (s *EngineTestSvc) CreateTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *EngineTestSvc) StartTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *EngineTestSvc) StopTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func (s *EngineTestSvc) CleanTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
