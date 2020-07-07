package db

import (
	"time"

	"git.supremind.info/testplatform/biz/proto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	imageCollName = "image"
)

type ImageMgnt interface {
	Insert(*proto.ImageInfo) (*proto.ImageInfo, error)
	Update(id bson.ObjectId, updateInfo map[string]interface{}) (*proto.ImageInfo, error)
}

type MongoImage struct {
	image *mgo.Collection
}

func NewMongoImage(dbc *mgo.Database) (ImageMgnt, error) {
	img := &MongoImage{
		image: dbc.C(imageCollName),
	}
	img.image.EnsureIndexKey("image", "user_id", "product", "status", "created_at")

	return img, nil
}

func (d *MongoImage) Insert(image *proto.ImageInfo) (*proto.ImageInfo, error) {
	if image.ID == "" {
		image.ID = bson.NewObjectId()
	}
	now := time.Now()
	image.CreatedAt = now
	image.UpdatedAt = now

	err := d.image.Insert(*image)
	if err != nil {
		log.Error("insert error:", err)
	}
	return image, err
}
func (d *MongoImage) Update(id bson.ObjectId, updateInfo map[string]interface{}) (*proto.ImageInfo, error) {
	now := time.Now()
	updateInfo["updated_at"] = now

	err := d.image.UpdateId(id, map[string]interface{}{
		"$set": updateInfo,
	})
	if err != nil {
		log.Error("update error:", err)
		return nil, err
	}
	image := &proto.ImageInfo{}
	err = d.image.FindId(id).One(image)
	return image, err
}
