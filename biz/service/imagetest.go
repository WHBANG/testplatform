package service

import (
	"context"
	"errors"

	bproto "git.supremind.info/testplatform/biz/proto"
	"git.supremind.info/testplatform/biz/service/db"
	"git.supremind.info/testplatform/biz/service/proto"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type ImageHandlerRouter interface {
	UpdateImageHandler(c *gin.Context)
	DeleteImageHander(c *gin.Context)
	InsertImageHandler(c *gin.Context)
	FindImageHandler(c *gin.Context)
	LikefindImageHandler(c *gin.Context)
	//LikefindByImageHandler(c *gin.Context)
	//FindByUserIDHandler(c *gin.Context)
	//LikefindByProductionHandler(c *gin.Context)
	//LikeFindByUpdateDate(c *gin.Context)
	//LikeFindByCreateDate(c *gin.Context)
	BatchDeleteImageHandler(c *gin.Context)
}

type HandlerRouter struct {
	mgnt db.ImageMgnt
}

var h HandlerRouter

/*
func showPage(page, size int, images []bproto.ImageInfo, c *gin.Context) {
	pagenum := page
	if page == 0 {
		pagenum = 1
	}
	sizenum := size
	if size == 0 {
		sizenum = 10
	}
	n := len(images)
	if n == 0 {
		DefaultImageRet(c, images, errors.New("null"))
		return
	}
	fmt.Println(pagenum, sizenum, n)
	maxpage := int(math.Ceil(float64(n) / float64(sizenum)))
	if pagenum > maxpage || pagenum <= 0 {
		DefaultImageRet(c, nil, errors.New("pagenum error"))
		return
	}
	begin := (pagenum - 1) * sizenum
	var cur_size int
	if pagenum == maxpage && 0 != n%sizenum {
		cur_size = n % sizenum
		//list := images[begin,begin+cur_size]
	} else {
		cur_size = sizenum
	}
	list := images[begin : begin+cur_size]
	imgs := make([]bproto.ImageInfo, cur_size)
	for i, v := range list {
		fmt.Println(i, "\n", v)
		image, _ := json.Marshal(v)
		json.Unmarshal(image, &imgs[i])
	}
	var data proto.GetImageRes
	data.Data = imgs
	data.Page = pagenum
	data.Size = sizenum
	data.Total = n
	DefaultImageRet(c, data, nil)
}
*/
func showPage(page, size, total int, images []bproto.ImageInfo, c *gin.Context) {
	if total == 0 {
		proto.DefaultRet(c, nil, errors.New("total is zero"))
		return
	}
	var data proto.GetImageRes
	data.Data = images
	data.Page = page
	data.Size = size
	data.Total = total
	proto.DefaultRet(c, data, nil)

}

// @Summary 更新镜像
// @Description 根据ID查找到对应的镜像，并将其更新为新输入的镜像信息
// @Accept json
// @Param id path string true "id"
// @Param example body proto.UpdateImageReq true "UpdateImageReq"
// @Success 200 {object}  proto.CommonRes{data=proto.ImageInfo}
// @Router /v1/image/{id} [PUT]
func (handler *HandlerRouter) UpdateImageHandler(c *gin.Context) {

	id := bson.ObjectIdHex(c.Param("id"))
	var structReq proto.UpdateImageReq
	err := c.BindJSON(&structReq)
	if err != nil {
		log.Error("Update Data Error: ", err)
		proto.DefaultRet(c, nil, err)
		return
	}

	var mapReq map[string]interface{}
	err = proto.ParseMapFromStruct(structReq, &mapReq)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	mapReq["status"] = bproto.Created
	img, err := h.mgnt.Update(id, mapReq)
	proto.DefaultRet(c, img, err)
}

// @Summary 删除镜像信息
// @Description 根据输入的镜像ID删除数据库中对应的镜像信息
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.ImageInfo}
// @Router /v1/image/{id} [delete]
func (handler *HandlerRouter) DeleteImageHandler(c *gin.Context) {
	id := c.Param("id")
	image, err := h.mgnt.Delete(bson.ObjectIdHex(id))
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, image, err)
}

// @Summary 批量删除
// @Description 根据输入的镜像ID数组删除数据库中镜像信息
// @Accept json
// @Param list query []string true "id" collectionFormat(multi)
// @Success 200 {object}  proto.CommonRes{data=proto.ImageInfo}
// @Router /v1/image/ [delete]
func (handler *HandlerRouter) BatchDeleteImageHandler(c *gin.Context) {
	var imageList []string
	imageList = c.QueryArray("list")
	len := len(imageList)
	if len == 0 {
		err := errors.New("Please select the data you want to delete!")
		proto.DefaultRet(c, nil, err)
		return
	}
	list_id := make([]bson.ObjectId, len)
	for i, v := range imageList {
		list_id[i] = bson.ObjectIdHex(v)
	}
	image, err := h.mgnt.BatchDelete(list_id)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	proto.DefaultRet(c, image, err)
}

// @Summary 新增镜像信息
// @Description 将输入的镜像信息添加到数据库中
// @Accept json
// @Param example body proto.InsertImageReq true "InsertImageReq"
// @Success 200 {object}  proto.CommonRes{data=proto.ImageInfo}
// @Router /v1/image/{id} [post]
func (Handler *HandlerRouter) InsertImageHandler(c *gin.Context) {

	var insertReq proto.InsertImageReq
	err := c.BindJSON(&insertReq)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		return
	}
	var image bproto.ImageInfo
	image.ID = insertReq.ID
	image.Image = insertReq.Image
	image.Product = insertReq.Product
	image.UserID = insertReq.UserID
	image.Description = insertReq.Description
	image.Models = insertReq.Models
	image.Status = bproto.Created
	img, err := h.mgnt.Insert(&image)
	if err != nil {
		log.Error("insert:insert err:", err)
		proto.DefaultRet(c, nil, err)
	} else {
		proto.DefaultRet(c, img, nil)
	}
}

// @Summary 查找镜像
// @Description 根据镜像ID在数据库中查找对应的镜像信息
// @Accept json
// @Param id path string true "id"
// @Success 200 {object}  proto.CommonRes{data=proto.ImageInfo}
// @Router /v1/image/ [get]
func (handler *HandlerRouter) FindImageHandler(c *gin.Context) {

	log.Println("find")
	id := bson.ObjectIdHex(c.Param("id")) //获取url后面的参数
	image, err := h.mgnt.Find(id)
	if err != nil {
		proto.DefaultRet(c, nil, err)
	} else {
		proto.DefaultRet(c, image, err)
	}
}

// @Summary 模糊查询
// @Description 根据镜像的字段信息进行模糊查询
// @Accept json
// @Param example query proto.GetImageReq false "GetImageReq"
// @Success 200 {object}  proto.CommonRes{data=proto.GetImageRes}
// @Router /v1/image/like/ [get]
func (handler *HandlerRouter) LikefindImageHandler(c *gin.Context) {

	log.Println("find like")
	getReq := proto.GetImageReq{}
	err := c.ShouldBindQuery(&getReq)
	if err != nil {
		log.Error("paramerter error:", err)
		proto.DefaultRet(c, nil, err)
		return
	}

	size := getReq.Size
	page := getReq.Page
	if size <= 0 {
		size = 10
	}
	if page <= 0 {
		page = 1
	}
	/*
		queryBytes, err := json.Marshal(proto.ImageQuery{
			Image:       getReq.Image,
			UserID:      getReq.UserID,
			Product:     getReq.Product,
			Description: getReq.Description,
		})
		if err != nil {
			log.Error("marshal error:", err)
			DefaultImageRet(c, nil, err)
		}
	*/
	var query []bson.M
	query1 := bson.M{"image": bson.M{"$regex": bson.RegEx{Pattern: getReq.Image, Options: "im"}}}
	query2 := bson.M{"product": bson.M{"$regex": bson.RegEx{Pattern: getReq.Product, Options: "im"}}}
	query3 := bson.M{"description": bson.M{"$regex": bson.RegEx{Pattern: getReq.Description, Options: "im"}}}
	query = []bson.M{query1, query2, query3}
	if getReq.UserID != 0 {
		query4 := bson.M{"user_id": getReq.UserID}
		query = append(query, query4)
	}
	q := bson.M{"$and": query}
	/*
		query = bson.M{"$and": []bson.M{bson.M{"image": bson.M{"$regex": bson.RegEx{Pattern: getReq.Image, Options: "im"}}},
				bson.M{"product": bson.M{"$regex": bson.RegEx{Pattern: getReq.Product, Options: "im"}}},
				bson.M{"user_id": getReq.UserID},
				bson.M{"description": bson.M{"$regex": bson.RegEx{Pattern: getReq.Description, Options: "im"}}}}}

	*/
	/*
		err = json.Unmarshal(queryBytes, &query)
		if err != nil {
			log.Error("unmarshal error:", err)
			DefaultImageRet(c, nil, err)
			return
		}
	*/
	images, total, err := h.mgnt.Likefind(q, page, size)
	if err != nil {
		proto.DefaultRet(c, nil, err)
	} else {
		showPage(page, size, total, images, c)
	}
}

func ImageHandler(group *gin.RouterGroup) {

	var svc HandlerRouter
	group.DELETE("/image/:id", svc.DeleteImageHandler)
	group.PUT("/image/:id", svc.UpdateImageHandler)
	group.POST("/image", svc.InsertImageHandler)
	group.GET("/image/:id", svc.FindImageHandler)
	group.GET("/image", svc.LikefindImageHandler)
	group.DELETE("/image", svc.BatchDeleteImageHandler)
}

func ImageHandlerSvc(ctx context.Context, imageMgnt db.ImageMgnt, group *gin.RouterGroup) {
	h.mgnt = imageMgnt
	ImageHandler(group)
}
