package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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

	return svc, nil
}

func DefaultRet(c *gin.Context, data interface{}, err error) {
	var res proto.CommonRes
	if err != nil {
		res.Code = proto.DefaultErrorCode
		res.Msg = err.Error()
	}
	res.Data = data
	c.JSON(http.StatusOK, res)
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

func (s *EngineTestSvc) CreateEngine(c *gin.Context) {
	var req proto.CreateEngineReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		DefaultRet(c, nil, err)
		return
	}

	newID := bson.NewObjectId()
	var analyzerConfig map[string]interface{}
	if req.AnalyzerConfig != nil {
		analyzerConfig, err = parseMapFromItf(req.AnalyzerConfig)
		if err != nil {
			err = fmt.Errorf("parse analyzer config file error:%s", err.Error())
			DefaultRet(c, nil, err)
			return
		}
		analyzerConfig["region"] = newID
	} else if analyzerConfigStr, ok := s.configFileMap["analyzer.conf"]; ok {
		analyzerConfig, err = parseMapFromStr(analyzerConfigStr)
		if err != nil {
			err = fmt.Errorf("parse default analyzer config file error:%s", err.Error())
			DefaultRet(c, nil, err)
			return
		}
		analyzerConfig["region"] = newID
	}

	info, err := s.testMgnt.CreateEngine(&bproto.EngineDeployInfo{
		ID:             newID,
		Image:          req.Image,
		UserID:         req.UserID,
		Description:    req.Description,
		Product:        req.Description,
		Status:         bproto.EngineStatusNone,
		AnalyzerConfig: analyzerConfig,
	})
	if err != nil {
		log.Error("create engine error:", err)
	}
	DefaultRet(c, *info, err)
}

func (s *EngineTestSvc) StartEngine(c *gin.Context) {
	var req proto.StartEngineReq
	var res interface{}
	var err error
	defer func() {
		DefaultRet(c, res, err)
	}()
	err = c.ShouldBindUri(&req)
	if err != nil {
		log.Error(err)
		return
	}
	info, err := s.testMgnt.GetEngineOne(bson.ObjectId(req.ID))
	if err != nil {
		log.Error("update mongo error:", err)
		return
	}
	configM := make(map[string]string)
	analyzerConfig, _ := json.Marshal(info.AnalyzerConfig)
	configM["analyzer.conf"] = string(analyzerConfig)

	jobDeployInfo := analyzerclient.AnalyzerDeployInfo{
		JobName:   req.ID,
		Image:     info.Image,
		ConfigMap: configM,
	}
	if _, ok := s.configFileMap["start.sh"]; ok {
		jobDeployInfo.Args = "sh /workspace/mnt/config/start.sh"
	}
	job, err := s.engineIfc.CreateAnalyzer(&jobDeployInfo)
	if err != nil {
		log.Error("CreateAnalyzer error:", err)
		return
	}
	info, err = s.testMgnt.UpdateEngine(bson.ObjectId(req.ID), bson.M{
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
	info, err = s.testMgnt.UpdateEngine(bson.ObjectId(req.ID), bson.M{
		"status":   bproto.EngineStatusStarted,
		"job_info": *job,
	})
	if err != nil {
		log.Error("update mongo error:", err)
		return
	}
	res = info
}

func (s *EngineTestSvc) StopEngine(c *gin.Context) {
	var req proto.StopEngineReq
	var res interface{}
	var err error
	defer func() {
		DefaultRet(c, res, err)
	}()

	err = c.ShouldBindUri(&req)
	if err != nil {
		log.Error(err)
		return
	}
	info, err := s.testMgnt.GetEngineOne(bson.ObjectId(req.ID))
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
	info, err = s.testMgnt.UpdateEngine(bson.ObjectId(req.ID), bson.M{
		"status": bproto.EngineStatusStoped,
	})
	if err != nil {
		log.Error("update mongo error:", err)
	}
	DefaultRet(c, *info, err)
}

func (s *EngineTestSvc) UpdateEngine(c *gin.Context) {
	var reqID string
	var ok bool
	if reqID, ok = c.Params.Get("id"); !ok {
		err := errors.New("No id specified")
		DefaultRet(c, nil, err)
		return
	}

	var req proto.UpdateEgnineReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		DefaultRet(c, nil, err)
		return
	}

	if req.AnalyzerConfig != nil {
		var analyzerConfig map[string]interface{}
		analyzerConfig, err = parseMapFromItf(req.AnalyzerConfig)
		if err != nil {
			err = fmt.Errorf("parse analyzer config file error:%s", err.Error())
			DefaultRet(c, nil, err)
			return
		}
		analyzerConfig["region"] = reqID
		req.AnalyzerConfig = analyzerConfig
	}
	updateStr, _ := json.Marshal(req)
	updateInfo, err := parseMapFromStr(string(updateStr))
	if err != nil {
		DefaultRet(c, nil, err)
		return
	}
	info, err := s.testMgnt.UpdateEngine(bson.ObjectId(reqID), updateInfo)
	if err != nil {
		log.Error("update mongo error:", err)
	}
	DefaultRet(c, *info, err)
}

func (s *EngineTestSvc) RemoveEngine(c *gin.Context) {
	var req proto.RemoveEngineReq
	err := c.ShouldBindUri(&req)
	if err != nil {
		log.Error(err)
		DefaultRet(c, nil, err)
		return
	}

	//todo check status

	err = s.testMgnt.RemoveEngine(bson.ObjectId(req.ID))
	if err != nil {
		log.Error("remove mongo error:", err)
	}

	DefaultRet(c, nil, err)
}

func (s *EngineTestSvc) GetEngine(c *gin.Context) {
	var reqID string
	var ok bool
	if reqID, ok = c.Params.Get("id"); !ok {
		err := errors.New("No id specified")
		DefaultRet(c, nil, err)
		return
	}

	info, err := s.testMgnt.GetEngineOne(bson.ObjectId(reqID))
	if err != nil {
		log.Error("get engine error:", err)
	}

	DefaultRet(c, *info, err)
}

func (s *EngineTestSvc) GetEngineList(c *gin.Context) {
	var req proto.GetEngineReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		DefaultRet(c, nil, err)
		return
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	queryBytes, err := json.Marshal(req.EngineQuery)
	if err != nil {
		log.Error("marshal error:", err)
		DefaultRet(c, nil, err)
		return
	}
	var query bson.M
	err = json.Unmarshal(queryBytes, &query)
	if err != nil {
		log.Error("unmarshal error:", err)
		DefaultRet(c, nil, err)
		return
	}
	dataList, num, err := s.testMgnt.GetEngine(query, req.Page, req.Size)
	if err != nil {
		log.Error("get engine list error:", err)
	}
	res := proto.GetEngineRes{
		Page:  req.Page,
		Size:  req.Size,
		Total: num,
		Data:  dataList,
	}
	DefaultRet(c, res, err)
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
