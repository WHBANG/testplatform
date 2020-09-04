package db

import (
	"time"

	bproto "git.supremind.info/testplatform/biz/proto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
