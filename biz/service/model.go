package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	bproto "git.supremind.info/testplatform/biz/proto"
	"git.supremind.info/testplatform/biz/service/db"
	"git.supremind.info/testplatform/biz/service/proto"
	"github.com/gin-gonic/gin"
)

type ModelHandlerImp interface {
	UpdateModelsHandler(c *gin.Context)
	ResetModelsHandler(c *gin.Context)

	UpdateModelHandler(c *gin.Context)
	DeleteModelHander(c *gin.Context)
	InsertModelHandler(c *gin.Context)
	FindModelHandler(c *gin.Context)
	BatchDeleteModelHandler(c *gin.Context)
}

type ModelHandler struct {
	mgnt   db.ModelMgnt
	host   *TargetHost
	models []proto.InsertModelReq
}

var m ModelHandler

// @Summary 更新模型
// @Description 根据ID字段查找到对应的模型，并将其更新为新输入的模型信息
// @Accept json
// @Param id path string true "id"
// @Param example body proto.UpdateModelReq true "UpdateModelReq"
// @Success 200 {object}  proto.CommonRes{data=proto.ModelInfo}
// @Router /v1/model/{id} [PUT]
func (handler *ModelHandler) UpdateModelHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Update Data Error: ", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	var mapReq map[string]interface{}
	var structReq proto.UpdateModelReq
	err = c.BindJSON(&structReq)
	if err != nil {
		log.Error("Update Data Error: ", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	err = proto.ParseMapFromStruct(structReq, &mapReq)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	model, err := m.mgnt.Update(id, mapReq)
	proto.DefaultRet(c, model, err)
}

// @Summary 删除模型
// @Description 根据模型ID删除数据库中对应的模型
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.ModelInfo}
// @Router /v1/model/delete/{id} [delete]
func (handler *ModelHandler) DeleteModelHander(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("Delete Data Error: ", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	model, err := m.mgnt.Delete(id)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, model, err)
}

// @Summary 批量删除模型
// @Description 根据输入的模型ID数组删除模型
// @Accept json
// @Param list query []string true "id" collectionFormat(multi)
// @Success 200 {object}  proto.CommonRes{data=[]proto.ModelInfo}
// @Router /v1/model/ [delete]
func (handler *ModelHandler) BatchDeleteModelHandler(c *gin.Context) {

	var modelList []string
	modelList = c.QueryArray("list")
	len := len(modelList)
	if len == 0 {
		err := errors.New("Please Select The Data You Want To Delete!")
		proto.DefaultRet(c, nil, err)
		return
	}
	list_id := make([]int, len)
	for i, v := range modelList {
		id, err := strconv.Atoi(v)
		if err != nil {
			err := errors.New("Query Data Error: ")
			proto.DefaultRet(c, nil, err)
			return
		}
		list_id[i] = id
	}
	models, err := m.mgnt.BatchDelete(list_id)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, models, err)
}

// @Summary 新增模型
// @Description 添加模型到数据库中
// @Accept json
// @Param example body proto.InsertModelReq true "InsertModelReq"
// @Success 200 {object}  proto.CommonRes{data=proto.ModelInfo}
// @Router /v1/model [post]
func (Handler *ModelHandler) InsertModelHandler(c *gin.Context) {

	var insertReq proto.InsertModelReq
	err := c.BindJSON(&insertReq)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	model := bproto.Model{}
	model.ModelID = insertReq.ModelID
	model.ModelName = insertReq.ModelName
	model.ModelType = insertReq.ModelType
	model.ModelURL = insertReq.ModelURL
	mod, err := m.mgnt.Insert(&model)
	if err != nil {
		log.Error("Insert Error: ", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, mod, nil)

}

// @Summary 查找模型
// @Description 根据模型ID查找对应的模型
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.ModelInfo}
// @Router /v1/model/id/{id} [get]
func (handler *ModelHandler) FindModelHandler(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	model, err := m.mgnt.FindByModelID(id)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, model, err)

}

func GetReleaseModels() (data []bproto.Model, err error) {
	req, err := http.NewRequest("GET", m.host.Host, nil)
	if err != nil {
		log.Println("Load Model Data Error: ", err)
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Load Model Data Error: ", err)
		return nil, err
	}
	var resp proto.GetModelsRes
	if res != nil {
		defer res.Body.Close()

		buffer := bytes.NewBuffer(make([]byte, 1024*8))
		_, err := io.Copy(buffer, res.Body)
		if err != nil {
			log.Println("Buffer Capacity Error: ", err)
			return nil, err
		}
		respData := buffer.Bytes()
		str := strings.Trim(string(respData), "\x00 \n")
		err = json.Unmarshal([]byte(str[:len(str)]), &resp)
		if err != nil {
			log.Println("Json Unmarshal Error: ", err)
			return nil, err
		}
	}
	for i := 0; i < len(resp); i++ {
		models := resp[i].Models
		for j := 0; j < len(models); j++ {
			var tmp bproto.Model
			tmp.ModelURL = models[j].ModelURL
			tmp.ModelID = models[j].ModelID
			tmp.ModelName = models[j].ModelName
			tmp.ModelType = models[j].ModelType
			data = append(data, tmp)
		}
	}
	return data, nil
}

// @Summary 重置模型数据库
// @Description 重置模型数据库
// @Accept json
// @Success 200 {object}  proto.CommonRes{data=proto.GetModelsRes}
// @Router /v1/model/reset [get]
func (handler *ModelHandler) ResetModelsHandler(c *gin.Context) {
	data, err := GetReleaseModels()
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	_, err = m.mgnt.Reset(data)
	if err != nil {
		log.Println("Reset Model Data Error: ", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, data, nil)
}

func (handler *ModelHandler) InitModelsHandler() error {
	data, err := GetReleaseModels()
	err = m.mgnt.Init(data)
	if err != nil {
		log.Println("Reset Model Data Error: ", err)
		return err
	}
	return nil
}

// @Summary 更新模型数据库
// @Description 更新模型数据库
// @Accept json
// @Success 200 {object}  proto.CommonRes{data=proto.GetModelsRes}
// @Router /v1/model/update [get]
func (handler *ModelHandler) UpdateModelsHandler(c *gin.Context) {
	data, err := GetReleaseModels()
	err = m.mgnt.Init(data)
	if err != nil {
		log.Println("Reset Model Data Error: ", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, nil, nil)
}

func ModelHandlerRouter(group *gin.RouterGroup) {

	group.GET("/model/reset", m.ResetModelsHandler)
	group.GET("/model/update", m.UpdateModelsHandler)

	group.PUT("/model/:id", m.UpdateModelHandler)
	group.DELETE("/model/delete/:id", m.DeleteModelHander)
	group.DELETE("/model/batchDelete", m.BatchDeleteModelHandler)
	group.POST("/model", m.InsertModelHandler)
	group.GET("model/id/:id", m.FindModelHandler)

}

func ModelHandlerSvc(ctx context.Context, modelMgnt db.ModelMgnt, group *gin.RouterGroup) {
	m.mgnt = modelMgnt
	m.host = &TargetHost{
		Host:    "http://100.100.142.18/get_release_models",
		IsHTTPS: false,
		CAPath:  "",
	}
	err := m.InitModelsHandler()
	if err != nil {
		log.Error("Init Model Data Error: ", err)
	}
	ModelHandlerRouter(group)
}
