package service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"

	bproto "git.supremind.info/testplatform/biz/proto"
	"git.supremind.info/testplatform/biz/service/db"
	"git.supremind.info/testplatform/biz/service/proto"
	"github.com/gin-gonic/gin"
)

type AnalyzerTypeImp interface {
	GetAnalyzerTypeHandler(c *gin.Context)
	FindAnalyzerTypeHandler(c *gin.Context)
	InsertAnalyzerTypeHandler(c *gin.Context)
	DeleteAnalyzerTypeHandler(c *gin.Context)
}

type AnalyzerTypeHandler struct {
	mgnt db.AnalyzerMgnt
}

var a AnalyzerTypeHandler

// @Summary 删除
// @Description 根据ID删除analyzerType
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.AnalyzerTypeInfo}
// @Router /v1/analyzertype/{id} [delete]
func (handler *AnalyzerTypeHandler) DeleteAnalyzerTypeHandler(c *gin.Context) {

	id := bson.ObjectIdHex(c.Param("id"))
	t, err := a.mgnt.Delete(id)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, t, err)
}

// @Summary 新增
// @Description 添加analyzerType到数据库中
// @Accept json
// @Param example body proto.InsertAnalyzerTypeReq true "InsertAnalyzerTypeReq"
// @Success 200 {object}  proto.CommonRes{data=proto.AnalyzerTypeInfo}
// @Router /v1/analyzertype [post]
func (Handler *AnalyzerTypeHandler) InsertAnalyzerTypeHandler(c *gin.Context) {

	var insertReq proto.InsertAnalyzerTypeReq
	err := c.BindJSON(&insertReq)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	an := bproto.AnalyzerTypeInfo{}
	an.AnalyzerType = insertReq.AnalyzerType
	an.ModelNameList = insertReq.ModelNameList
	mod, err := a.mgnt.Insert(&an)
	if err != nil {
		log.Error("Insert Error: ", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, mod, nil)

}

// @Summary 查找
// @Description 根据ID查找对应的analyzertype
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.AnalyzerTypeInfo}
// @Router /v1/analyzertype/{id} [get]
func (handler *AnalyzerTypeHandler) FindAnalyzerTypeHandler(c *gin.Context) {

	id := bson.ObjectIdHex(c.Param("id"))
	model, err := a.mgnt.Find(id)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, model, err)

}

// @Summary 获取所有
// @Description 获取所有的analyzertype
// @Accept json
// @Success 200 {object}  proto.CommonRes{data=[]proto.AnalyzerTypeInfo}
// @Router /v1/analyzertype/ [get]
func (handler *AnalyzerTypeHandler) GetAnalyzerTypeHandler(c *gin.Context) {
	data, err := a.mgnt.FindAll()
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, data, nil)
}

func AnalyzerTypeHandlerRouter(group *gin.RouterGroup) {

	group.GET("/analyzertype", a.GetAnalyzerTypeHandler)
	group.GET("/analyzertype/:id", a.FindAnalyzerTypeHandler)
	group.DELETE("/analyzertype/:id", a.DeleteAnalyzerTypeHandler)
	group.POST("analyzertype", a.InsertAnalyzerTypeHandler)

}

func AnalyzerTypeHandlerSvc(ctx context.Context, mgnt db.AnalyzerMgnt, group *gin.RouterGroup) {
	a.mgnt = mgnt
	AnalyzerTypeHandlerRouter(group)
}
