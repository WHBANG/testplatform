package db

import (
	"gopkg.in/mgo.v2"
)

type TestMgnt interface {
	// CreateEngine(info *bproto.EngineDeployInfo) (bproto.EngineDeployInfo, error)
	// UpdateEngine(id bson.ObjectId, updateInfo bson.M) (bproto.EngineDeployInfo, error)
	// RemoveEngine(id bson.ObjectId) error
	// GetEngineOne(id bson.ObjectId) (bproto.EngineDeployInfo, error)
	// GetEngine(query bson.M, page int, size int) ([]bproto.EngineDeployInfo, error)
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
