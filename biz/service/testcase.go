package service

import (
	"context"
	"errors"

	bproto "git.supremind.info/testplatform/biz/proto"
	"git.supremind.info/testplatform/biz/service/db"
	"git.supremind.info/testplatform/biz/service/proto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

type TestCastHandlerImp interface {
	InsertTestCaseHandler(c *gin.Context)
	DeleteTestCaseHandler(c *gin.Context)
	UpdateTestCaseHandler(c *gin.Context)
	FindTestCaseHandler(c *gin.Context)
	LikeFindHandler(c *gin.Context)
	BatchDeleteTestCaseHandler(c *gin.Context)
}

type TestCastRouter struct {
	testcase db.TestCaseMgnt
}

var t TestCastRouter

func showTestCasePage(page, size, total int, images []bproto.TestCaseInfo, c *gin.Context) {
	if total == 0 {
		proto.DefaultRet(c, nil, errors.New("total is zero"))
		return
	}
	var data proto.GetTestCaseRes
	data.Data = images
	data.Page = page
	data.Size = size
	data.Total = total
	proto.DefaultRet(c, data, nil)

}

// @Summary 插入测试用例
// @Description 向数据库中插入新的测试用例
// @Accept json
// @Param example body proto.InsertTestCaseReq true "InsertTestCaseReq"
// @Success 200 {object}  proto.CommonRes{data=proto.TestCaseInfo}
// @Router /testcase/ [post]
func (handler *TestCastRouter) InsertTestCaseHandler(c *gin.Context) {
	var insertReq proto.InsertTestCaseReq
	err := c.BindJSON(&insertReq)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	var tc bproto.TestCaseInfo
	tc.UserID = insertReq.UserID
	tc.TestData = insertReq.TestData
	tc.Description = insertReq.Description
	data, err := t.testcase.Insert(&tc)
	if err != nil {
		log.Fatal("insert error!", err)
		proto.DefaultRet(c, nil, err)
	} else {
		proto.DefaultRet(c, data, nil)
	}
}

// @Summary 删除测试用例
// @Description 删除指定ID所对应的测试用例数据
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.TestCaseInfo}
// @Router /testcase/{id} [delete]
func (handler *TestCastRouter) DeleteTestCaseHandler(c *gin.Context) {
	id := c.Param("id")
	data, err := t.testcase.Delete(bson.ObjectIdHex(id))
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, data, nil)
}

// @Summary 批量删除
// @Description 根据id数组批量删除测试用例数据
// @Accept json
// @Param list query []string true "list" collectionFormat(multi)
// @Success 200 {object}  proto.CommonRes{data=proto.TestCaseInfo}
// @Router /testcase/ [delete]
func (handler *TestCastRouter) BatchDeleteTestCaseHandler(c *gin.Context) {
	var list []string
	list = c.QueryArray("list")
	log.Println(list)
	len := len(list)
	if len == 0 {
		err := errors.New("Please select the data you want to delete!")
		proto.DefaultRet(c, nil, err)
		return
	}
	list_id := make([]bson.ObjectId, len)
	for i, v := range list {
		list_id[i] = bson.ObjectIdHex(v)
	}
	log.Println(list)
	data, err := t.testcase.BatchDelete(list_id)
	if err != nil {
		proto.DefaultRet(c, nil, err)
	} else {
		proto.DefaultRet(c, data, err)
	}
}

// @Summary 更新测试用例
// @Description 更新指定ID所对应的测试用例的信息
// @Accept json
// @Param id path string true "id"
// @Param example body proto.UpdateTestCaseReq true "updateTestCaseReq"
// @Success 200 {object}  proto.CommonRes{data=proto.TestCaseInfo}
// @Router /testcase/{id} [put]
func (handler *TestCastRouter) UpdateTestCaseHandler(c *gin.Context) {
	id := bson.ObjectIdHex(c.Param("id"))
	var updateReq proto.UpdateTestCaseReq
	err := c.BindJSON(&updateReq)
	if err != nil {
		log.Println("update data error", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	var mapReq map[string]interface{}
	err = proto.ParseMapFromStruct(updateReq, &mapReq)
	if err != nil {
		log.Println("parseMapFromStruct error", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	data, err := t.testcase.Update(id, mapReq)
	if err != nil {
		log.Println("update error", err)
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, data, err)
}

// @Summary 查找测试用例
// @Description 根据ID在数据库中查找对应的测试用例信息
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.TestCaseInfo}
// @Router /testcase/ [get]
func (handler *TestCastRouter) FindTestCaseHandler(c *gin.Context) {
	id := c.Param("id")
	_id := bson.ObjectIdHex(id)
	data, err := t.testcase.Find(_id)
	proto.DefaultRet(c, data, err)
}

// @Summary 模糊查询测试用例
// @Description 根据测试用例的字段信息进行模糊查询
// @Accept json
// @Param example query proto.GetTestCaseReq false "GetTestCaseReq"
// @Success 200 {object}  proto.CommonRes{data=proto.GetTestCaseRes}
// @Router /testcase/like/ [get]
func (handler *TestCastRouter) LikeFindTestCaseHandler(c *gin.Context) {
	getReq := proto.GetTestCaseReq{}
	err := c.ShouldBindQuery(&getReq)
	if err != nil {
		log.Error("paramter error!:", err)
		proto.DefaultRet(c, nil, err)
	}
	size := getReq.Size
	page := getReq.Page
	if size <= 0 {
		size = 10
	}
	if page <= 0 {
		page = 1
	}
	var query []bson.M
	//query1 := bson.M{"test_data": getReq.TestData}
	query1 := bson.M{"description": bson.M{"$regex": bson.RegEx{Pattern: getReq.Description, Options: "im"}}}
	//query2 := bson.M{"product": bson.M{"$regex": bson.RegEx{Pattern: getReq.Product, Options: "im"}}}
	query = []bson.M{query1}
	if getReq.TestData != nil {
		query3 := bson.M{"test_data": getReq.TestData}
		query = append(query, query3)
	}
	if getReq.UserID != 0 {
		query4 := bson.M{"user_id": getReq.UserID}
		query = append(query, query4)
	}
	q := bson.M{"$and": query}
	data, total, err := t.testcase.LikeFind(q, page, size)
	if err != nil {
		proto.DefaultRet(c, nil, err)
	} else {
		showTestCasePage(page, size, total, data, c)
	}
}

func TestCaseHandler(group *gin.RouterGroup) {
	var svc TestCastRouter
	group.POST("/testcase", svc.InsertTestCaseHandler)
	group.DELETE("/testcase/:id", svc.DeleteTestCaseHandler)
	group.DELETE("testcase", svc.BatchDeleteTestCaseHandler)
	group.PUT("/testcase/:id", svc.UpdateTestCaseHandler)
	group.GET("/testcase/:id", svc.FindTestCaseHandler)
	group.GET("/testcase", svc.LikeFindTestCaseHandler)
}

func TestCaseHandlerSvc(ctx context.Context, testcase db.TestCaseMgnt, group *gin.RouterGroup) {

	t.testcase = testcase
	TestCaseHandler(group)

}
