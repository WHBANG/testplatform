package service

import (
	"context"

	"git.supremind.info/testplatform/biz/service/db"
	"github.com/gin-gonic/gin"
)

type Config struct {
}

type TestPlatformSvc struct {
	config    *Config
	imageMgnt db.ImageMgnt
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
	svc := &TestPlatformSvc{
		config:    config,
		imageMgnt: imageMgnt,
	}

	group.GET("/ping", svc.Ping)
	group.POST("/image", svc.Ping)
	group.PUT("/image/:id", svc.Ping)
	group.GET("/image/:id", svc.Ping)

	return svc, nil
}

func (s *TestPlatformSvc) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
