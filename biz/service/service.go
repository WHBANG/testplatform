package service

import (
	"context"

	"git.supremind.info/testplatform/biz/service/db"
	"github.com/gin-gonic/gin"
)

type Config struct {
	VMRClient
}

type VMRClient struct {
	FlowHost   string `json:"flow_host"`
	WebGeneral string `json:"web_general"`
	IsHTTPS    bool   `json:"is_https"`
	CAPath     string `json:"ca_path"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type TestPlatformSvc struct {
	config       *Config
	imageMgnt    db.ImageMgnt
	testcaseMgnt db.TestCaseMgnt
	task         db.TaskMgnt
	model        db.ModelMgnt
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

	taskMgnt, err := db.NewMongoTask(session, db.GetDBName())
	if err != nil {
		return nil, err
	}

	modelMgnt, err := db.NewMongoModel(session, db.GetDBName())
	if err != nil {
		return nil, err
	}

	svc := &TestPlatformSvc{
		config:       config,
		imageMgnt:    imageMgnt,
		testcaseMgnt: testcaseMgnt,
		task:         taskMgnt,
		model:        modelMgnt,
	}

	ImageHandlerSvc(ctx, imageMgnt, group)
	TestCaseHandlerSvc(ctx, testcaseMgnt, group)
	TaskHandlerSvc(ctx, taskMgnt, group, config.VMRClient)
	ModelHandlerSvc(ctx, modelMgnt, group)

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
