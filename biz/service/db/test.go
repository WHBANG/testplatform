package db

import (
	"time"

	bproto "git.supremind.info/testplatform/biz/proto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TestMgnt interface {
	CreateEngine(info *bproto.EngineDeployInfo) (*bproto.EngineDeployInfo, error)
	UpdateEngine(id bson.ObjectId, updateInfo bson.M) (*bproto.EngineDeployInfo, error)
	RemoveEngine(id bson.ObjectId) error
	GetEngineOne(id bson.ObjectId) (*bproto.EngineDeployInfo, error)
	GetEngine(query bson.M, page int, size int) ([]bproto.EngineDeployInfo, int, error)
}

type MongoTest struct {
	DB                 string
	session            *mgo.Session
	engineCollName     string
	engineTestCollName string
}

func NewMongoTest(s *mgo.Session, db string) (TestMgnt, error) {

	c := &MongoTest{
		DB:                 db,
		session:            s,
		engineCollName:     "engine",
		engineTestCollName: "engine_test",
	}
	s.DB(db).C(c.engineCollName).EnsureIndexKey("image", "user_id", "product", "status", "created_at")
	s.DB(db).C(c.engineTestCollName).EnsureIndexKey("engine_id", "user_id", "product", "test_status", "created_at")
	return c, nil
}

func (d *MongoTest) CreateEngine(info *bproto.EngineDeployInfo) (*bproto.EngineDeployInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	if info.ID == "" {
		info.ID = bson.NewObjectId()
	}
	now := time.Now()
	info.CreatedAt = now
	info.UpdatedAt = now

	err := c.Insert(*info)
	if err != nil {
		log.Error("insert error:", err)
	}

	return info, nil
}
func (d *MongoTest) UpdateEngine(id bson.ObjectId, updateInfo bson.M) (*bproto.EngineDeployInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(imageCollName)

	now := time.Now()
	updateInfo["updated_at"] = now

	err := c.UpdateId(id, map[string]interface{}{
		"$set": updateInfo,
	})
	if err != nil {
		log.Error("update error:", err)
		return nil, err
	}
	info := &bproto.EngineDeployInfo{}
	err = c.FindId(id).One(info)
	return info, err
}

func (d *MongoTest) RemoveEngine(id bson.ObjectId) error {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(d.engineCollName)

	//todo check engine status
	err := c.RemoveId(id)
	return err

}
func (d *MongoTest) GetEngineOne(id bson.ObjectId) (*bproto.EngineDeployInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(d.engineCollName)

	info := &bproto.EngineDeployInfo{}
	err := c.FindId(id).One(info)
	return info, err

}
func (d *MongoTest) GetEngine(query bson.M, page int, size int) ([]bproto.EngineDeployInfo, int, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(d.engineCollName)

	list := make([]bproto.EngineDeployInfo, size)
	err := c.Find(query).Skip((page - 1) * size).Limit(size).Sort("-_id").All(&list)
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}
	total, err := c.Find(query).Count()
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}
	return list, total, err
}
