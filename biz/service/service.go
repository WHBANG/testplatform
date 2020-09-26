package service

import (
	"context"

	"git.supremind.info/testplatform/biz/service/db"
	"github.com/gin-gonic/gin"
)

type Config struct {
	VMRClient
}

/*
 		"flow_host":"100.100.142.75:10066",
        "web_general":"100.100.142.75:10080",
        "is_https":false,
        "ca_path":"",
        "username":"super",
        "password":"smai123",
        "global_device_id": "vms229.QN00115a0ecb88f9ab69",
        "name_prefix": "test_wh"
*/

type VMRClient struct {
	FlowHost       string `json:"flow_host"`
	WebGeneral     string `json:"web_general"`
	IsHTTPS        bool   `json:"is_https"`
	CAPath         string `json:"ca_path"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	GlobalDeviceID string `json:"global_device_id"`
	NamePrefix     string `json:"name_prefix"`
	MaxChannel     int    `json:"max_channel"`
}

type TestPlatformSvc struct {
	config       *Config
	imageMgnt    db.ImageMgnt
	testcaseMgnt db.TestCaseMgnt
	task         db.EngineTaskMgnt
	model        db.ModelMgnt
	analyzerType db.AnalyzerMgnt
}

func NewTestPlatformSvc(ctx context.Context, config *Config, group *gin.RouterGroup) (*TestPlatformSvc, error) {
	session, err := db.GetMgoDBSession()
	if err != nil {
		return nil, err
	}
	imageMgnt, err := db.NewMongoImage(session, db.GetDBName())
	if err != nil {
		return nil, err
	}

	testcaseMgnt, err := db.NewMongoCase(session, db.GetDBName())
	if err != nil {
		return nil, err
	}

	taskMgnt, err := db.NewMongoEngineTask(session, db.GetDBName())
	if err != nil {
		return nil, err
	}

	modelMgnt, err := db.NewMongoModel(session, db.GetDBName())
	if err != nil {
		return nil, err
	}
	analyzerTypeMgnt, err := db.NewMongoAnalyzerType(session, db.GetDBName())
	if err != nil {
		return nil, err
	}

	svc := &TestPlatformSvc{
		config:       config,
		imageMgnt:    imageMgnt,
		testcaseMgnt: testcaseMgnt,
		task:         taskMgnt,
		model:        modelMgnt,
		analyzerType: analyzerTypeMgnt,
	}

	ImageHandlerSvc(ctx, imageMgnt, group)
	TestCaseHandlerSvc(ctx, testcaseMgnt, group)
	TaskHandlerSvc(ctx, taskMgnt, testcaseMgnt, group, config.VMRClient)
	ModelHandlerSvc(ctx, modelMgnt, group)
	AnalyzerTypeHandlerSvc(ctx, analyzerTypeMgnt, group)

	//group.GET("/ping", svc.Ping)
	//group.POST("/image", svc.Ping)
	//group.PUT("/image/:id", svc.Ping)
	//group.GET("/image/:id", svc.Ping)

	return svc, nil
}

func (s *TestPlatformSvc) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
