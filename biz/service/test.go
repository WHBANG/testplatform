package service

import (
	"context"
	"net/http"

	"git.supremind.info/testplatform/biz/analyzerclient"
	"git.supremind.info/testplatform/biz/service/db"
	"git.supremind.info/testplatform/biz/service/proto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TestSvc struct {
	imageMgnt db.ImageMgnt
	engineIfc analyzerclient.AnalyzerInterface
}

func NewTestSvc(ctx context.Context, group *gin.RouterGroup, engineIfc analyzerclient.AnalyzerInterface) (*TestSvc, error) {
	dbConn, err := db.GetMgoDB()
	if err != nil {
		return nil, err
	}
	imageMgnt, err := db.NewMongoImage(dbConn)
	if err != nil {
		return nil, err
	}
	svc := &TestSvc{
		imageMgnt: imageMgnt,
		engineIfc: engineIfc,
	}

	group.GET("/engine/:id", svc.GetEngine)
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

func (s *TestSvc) CreateEngine(c *gin.Context) {
	var req proto.CreateEngineReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		DefaultRet(c, nil, err)
		return
	}

}

func (s *TestSvc) StartEngine(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *TestSvc) StopEngine(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *TestSvc) UpdateEngine(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *TestSvc) RemoveEngine(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *TestSvc) GetEngine(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *TestSvc) CreateTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *TestSvc) StartTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *TestSvc) StopTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func (s *TestSvc) CleanTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
