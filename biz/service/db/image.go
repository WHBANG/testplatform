package db

import (
	"time"

	bproto "git.supremind.info/testplatform/biz/proto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	imageCollName = "image"
)

type ImageMgnt interface {
	Insert(*bproto.ImageInfo) (*bproto.ImageInfo, error)
	Update(id bson.ObjectId, updateInfo map[string]interface{}) (*bproto.ImageInfo, error)
	Delete(id bson.ObjectId) (*bproto.ImageInfo, error)
	Find(id bson.ObjectId) (*bproto.ImageInfo, error)
	Likefind(query bson.M, page, size int) ([]bproto.ImageInfo, int, error)
	BatchDelete(id []bson.ObjectId) ([]bproto.ImageInfo, error)
	/*
		LikefindByImage(msg string) ([]proto.ImageInfo, error)
		FindByUserID(userID int) ([]proto.ImageInfo, error)
		LikefindByProduction(production string) ([]proto.ImageInfo, error)
		LikefindByTime(way, from, to string) ([]proto.ImageInfo, error)
		LikefindByCreateTime(from, to string) ([]proto.ImageInfo, error)
		LikefindByUpdateTime(from, to string) ([]proto.ImageInfo, error)
	*/
}

type MongoImage struct {
	collName string
	DB       string
	session  *mgo.Session
}

func NewMongoImage(s *mgo.Session, db string) (ImageMgnt, error) {
	img := &MongoImage{
		session:  s,
		DB:       db,
		collName: imageCollName,
	}
	s.DB(db).C(imageCollName).EnsureIndexKey("image", "user_id", "product", "status", "created_at")
	return img, nil
}

func (d *MongoImage) Insert(image *bproto.ImageInfo) (*bproto.ImageInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	if image.ID == "" {
		image.ID = bson.NewObjectId()
	}
	now := time.Now()
	image.CreatedAt = now
	image.UpdatedAt = now

	err := c.Insert(*image)
	if err != nil {
		log.Error("Insert Error: ", err)
		return nil, err
	}
	return image, err
}
func (d *MongoImage) Update(id bson.ObjectId, updateInfo map[string]interface{}) (*bproto.ImageInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	now := time.Now()
	updateInfo["updated_at"] = now

	err := c.UpdateId(id, map[string]interface{}{
		"$set": updateInfo,
	})
	if err != nil {
		log.Error("Update Error: ", err)
		return nil, err
	}
	image := &bproto.ImageInfo{}
	err = c.FindId(id).One(image)
	return image, err
}

func (d *MongoImage) Delete(id bson.ObjectId) (*bproto.ImageInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	image := &bproto.ImageInfo{}
	c.Find(bson.M{"_id": id}).One(image)
	err := c.Remove(bson.M{"_id": id})
	if err != nil {
		log.Error("Delete Error:", err)
		return nil, err
	}
	return image, err
}

func (d *MongoImage) BatchDelete(list []bson.ObjectId) ([]bproto.ImageInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	images := []bproto.ImageInfo{}
	c.Find(bson.M{"_id": bson.M{"$in": list}}).All(&images)
	_, err := c.RemoveAll(bson.M{"_id": bson.M{"$in": list}})
	if err != nil {
		log.Error("Delete Error: ", err)
		return nil, err
	}
	return images, err
}

func (d *MongoImage) Find(id bson.ObjectId) (*bproto.ImageInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	image := &bproto.ImageInfo{}
	err := c.Find(bson.M{"_id": id}).One(image)
	if err != nil {
		log.Error("Find Error:", err)
		return nil, err
	}
	return image, err
}

func (d *MongoImage) Likefind(query bson.M, page, size int) ([]bproto.ImageInfo, int, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	//var images []proto.ImageInfo
	/*
		err := c.Find(bson.M{"$or": []bson.M{bson.M{"image": bson.M{"$regex": bson.RegEx{Pattern: msg, Options: "im"}}},
			bson.M{"product": bson.M{"$regex": bson.RegEx{Pattern: msg, Options: "im"}}},
			bson.M{"description": bson.M{"$regex": bson.RegEx{Pattern: msg, Options: "im"}}}}}).All(&images)
	*/
	ilist := make([]bproto.ImageInfo, size)
	err := c.Find(query).Skip((page - 1) * size).Limit(size).Sort("-_id").All(&ilist)
	if err != nil {
		log.Error("Image Query Error:", err)
		return nil, 0, err
	}
	total, err := c.Find(query).Count()
	if err != nil {
		log.Error("Image Query Count Error:", err)
		return nil, 0, err
	}
	return ilist, total, err
}

/*
func (d *MongoImage) LikefindByImage(imagemsg string) ([]proto.ImageInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	var images []proto.ImageInfo
	err := c.Find(bson.M{"image": bson.M{"$regex": bson.RegEx{Pattern: imagemsg, Options: "im"}}}).All(&images)
	if err != nil {
		log.Error("like find error:", err)
	}
	return images, err
}

func (d *MongoImage) FindByUserID(user_id int) ([]proto.ImageInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	var images []proto.ImageInfo
	err := c.Find(bson.M{"user_id": user_id}).All(&images)
	if err != nil {
		log.Error("like find by user_id message error:", err)
	}
	return images, err
}

func (d *MongoImage) LikefindByProduction(productionmsg string) ([]proto.ImageInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	var images []proto.ImageInfo
	err := c.Find(bson.M{"product": bson.M{"$regex": bson.RegEx{Pattern: productionmsg, Options: "im"}}}).All(&images)
	if err != nil {
		log.Error("like find by drocuction message error:", err)
	}
	return images, err
}

func (d *MongoImage) LikefindByTime(way, from, to string) ([]proto.ImageInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	var images []proto.ImageInfo
	starttime := from + " 00:00:00"
	endtime := to + " 23:59:59"
	start, _ := time.ParseInLocation("2006-01-02 15:04:05", starttime, time.Local)
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", endtime, time.Local)
	err := c.Find(bson.M{way: bson.M{"$gte": start, "$lte": end}}).All(&images)
	return images, err
}

// db.image.find({"created_at":{"$gte":ISODate("2020-07-19T00:00:00Z"),"$lte":ISODate("2020-07-19T23:59:59Z")}})

func (d *MongoImage) LikefindByCreateTime(from, to string) ([]proto.ImageInfo, error) {
	images, err := d.LikefindByTime("created_at", from, to)
	if err != nil {
		log.Error("like find by create time error")
	}
	return images, err
}

func (d *MongoImage) LikefindByUpdateTime(from, to string) ([]proto.ImageInfo, error) {
	images, err := d.LikefindByTime("updated_at", from, to)
	if err != nil {
		log.Error("like find by update time error")
	}
	return images, err
}
*/
