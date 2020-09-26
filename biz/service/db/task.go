package db

import (
	bproto "git.supremind.info/testplatform/biz/proto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
const (
	taskCollName = "tasks"
)

type TaskMgnt interface {
	Insert(*bproto.TaskInfo) (*bproto.TaskInfo, error)
}

type Task struct {
	collName string
	DB       string
	session  *mgo.Session
}

func NewMongoTask(c *mgo.Session, db string) (TaskMgnt, error) {
	coll := &Task{
		session:  c,
		DB:       db,
		collName: taskCollName,
	}
	c.DB(db).C(coll.collName).EnsureIndexKey("_id", "type", "name")
	return coll, nil
}

func (m *Task) Insert(tc *bproto.TaskInfo) (*bproto.TaskInfo, error) {
	session := m.session.Clone()
	defer session.Close()
	c := session.DB(m.DB).C(m.collName)

	if tc.IID == "" {
		tc.IID = bson.NewObjectId()
	}
	tc.CreatedAt = time.Now()
	tc.UpdatedAt = time.Now()
	err := c.Insert(*tc)
	if err != nil {
		log.Error("Insert Error: ", err)
		return nil, err
	}
	return tc, err
}
*/

const (
	engineTaskCollName = "enginetasks"
)

type EngineTaskMgnt interface {
	Insert(*bproto.EngineTaskInfo) (*bproto.EngineTaskInfo, error)
	Delete(taskID string) error
	FindAllByEngineId(engineID string) ([]bproto.EngineTaskInfo, error)
	UpdateStatus(taskID string, dataMap map[string]interface{}) error
	GetAll() ([]bproto.EngineTaskInfo, error)
}

type EngineTask struct {
	collName string
	DB       string
	session  *mgo.Session
}

func NewMongoEngineTask(c *mgo.Session, db string) (EngineTaskMgnt, error) {
	coll := &EngineTask{
		session:  c,
		DB:       db,
		collName: engineTaskCollName,
	}

	index := mgo.Index{
		Key:    []string{"task_id"},
		Unique: true,
	}
	c.DB(db).C(engineTaskCollName).EnsureIndex(index)
	return coll, nil
}

func (m *EngineTask) Insert(tc *bproto.EngineTaskInfo) (*bproto.EngineTaskInfo, error) {
	session := m.session.Clone()
	defer session.Close()
	c := session.DB(m.DB).C(m.collName)

	err := c.Insert(*tc)
	if err != nil {
		log.Error("Insert Error11: ", err)
		return nil, err
	}
	return tc, err
}

func (m *EngineTask) Delete(taskID string) error {
	session := m.session.Clone()
	defer session.Close()
	c := session.DB(m.DB).C(m.collName)

	//err := c.Remove(bson.M{"task_id": taskID})
	var updateInfo map[string]string
	updateInfo["status"] = "off"

	err := c.UpdateId(taskID, map[string]interface{}{
		"$set": updateInfo,
	})

	if err != nil {
		log.Error("Delete Error: ", err)
		return err
	}
	return nil
}

func (m *EngineTask) FindAllByEngineId(engineID string) ([]bproto.EngineTaskInfo, error) {
	session := m.session.Clone()
	defer session.Close()
	c := session.DB(m.DB).C(m.collName)

	enginetasks := []bproto.EngineTaskInfo{}
	err := c.Find(bson.M{"$and": []bson.M{bson.M{"engine_id": engineID}, bson.M{"status": "on"}}}).All(&enginetasks)
	//err := c.Find(bson.M{"engine_id": engineID}).All(&enginetasks)
	if err != nil {
		log.Error("Find Error: ", err)
		return nil, err
	}
	return enginetasks, nil
}

func (m *EngineTask) GetAll() ([]bproto.EngineTaskInfo, error) {
	session := m.session.Clone()
	defer session.Close()
	c := session.DB(m.DB).C(m.collName)

	enginetasks := []bproto.EngineTaskInfo{}
	err := c.Find(nil).All(&enginetasks)
	//err := c.Find(bson.M{"engine_id": engineID}).All(&enginetasks)
	if err != nil {
		log.Error("Get All Error: ", err)
		return nil, err
	}
	return enginetasks, nil
}

func (m *EngineTask) UpdateStatus(taskID string, dataMap map[string]interface{}) error {
	session := m.session.Clone()
	defer session.Close()
	c := session.DB(m.DB).C(m.collName)

	err := c.Update(bson.M{"task_id": taskID}, map[string]interface{}{
		"$set": dataMap,
	})
	if err != nil {
		log.Error("Find Error: ", err)
		return err
	}
	return nil
}
