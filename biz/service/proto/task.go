package proto

import (
	"git.supremind.info/product/visionmind/com/flow"
	bproto "git.supremind.info/testplatform/biz/proto"
	"gopkg.in/mgo.v2/bson"
)

type GetTaskRes struct {
	Data interface{} `bson:"data" json:"data"`
	Code int         `bson:"code" json:"code"`
}

type GetCreateTaskReq struct {
	TestCaseID  bson.ObjectId `json:"_id"`
	TestData    TestData      `json:"test_data"`
	UserID      int           `json:"user_id"`
	Description string        `json:"description"`
	Product     string        `json:"product"`
}

type TestData struct {
	Case  bproto.MetaCase    `json:"case"`
	Task  flow.Task          `json:"task"`
	Event []bproto.EventData `json:"event"`
	Files bproto.FileCase    `json:"files"`
}

type TaskListReq struct {
	TaskID     string            `bson:"task_id" json:"task_id"`
	TaskStatus bproto.TaskStatus `bson:"task_status" json:"task_status"`
}
