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
	engine     *mgo.Collection
	engineTest *mgo.Collection
}

func NewMongoTest(dbc *mgo.Database) (TestMgnt, error) {
	c := &MongoTest{
		engine:     dbc.C("engine"),
		engineTest: dbc.C("engine_test"),
	}
	c.engine.EnsureIndexKey("image", "user_id", "product", "status", "created_at")
	c.engineTest.EnsureIndexKey("engine_id", "user_id", "product", "test_status", "created_at")
	return c, nil
}

func (c *MongoTest) CreateEngine(info *bproto.EngineDeployInfo) (*bproto.EngineDeployInfo, error) {
	if info.ID == "" {
		info.ID = bson.NewObjectId()
	}
	now := time.Now()
	info.CreatedAt = now
	info.UpdatedAt = now

	err := c.engine.Insert(*info)
	if err != nil {
		log.Error("insert error:", err)
	}

	return info, nil
}
func (c *MongoTest) UpdateEngine(id bson.ObjectId, updateInfo bson.M) (*bproto.EngineDeployInfo, error) {
	now := time.Now()
	updateInfo["updated_at"] = now

	err := c.engine.UpdateId(id, map[string]interface{}{
		"$set": updateInfo,
	})
	if err != nil {
		log.Error("update error:", err)
		return nil, err
	}
	info := &bproto.EngineDeployInfo{}
	err = c.engine.FindId(id).One(info)
	return info, err
}

func (c *MongoTest) RemoveEngine(id bson.ObjectId) error {
	//todo check engine status
	err := c.engine.RemoveId(id)
	return err

}
func (c *MongoTest) GetEngineOne(id bson.ObjectId) (*bproto.EngineDeployInfo, error) {
	info := &bproto.EngineDeployInfo{}
	err := c.engine.FindId(id).One(info)
	return info, err

}
func (c *MongoTest) GetEngine(query bson.M, page int, size int) ([]bproto.EngineDeployInfo, int, error) {
	list := make([]bproto.EngineDeployInfo, size)
	err := c.engine.Find(query).Skip((page - 1) * size).Limit(size).Sort("-_id").All(&list)
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}
	total, err := c.engine.Find(query).Count()
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}
	return list, total, err
}
