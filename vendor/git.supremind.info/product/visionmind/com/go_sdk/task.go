package go_sdk

import (
	"git.supremind.info/product/visionmind/com/flow"
)

type ITaskMgmClient interface {
	CreateTask(task *flow.Task) (err error)
	DeleteTask(taskId string) (err error)
	UpdateTask(taskId string, req flow.UpdateTaskReq) (task flow.Task, err error)
	StartTask(taskId string) (err error)
	StopTask(taskId string) (err error)
	GetTask(taskId string) (task flow.Task, err error)
	SearchTask(req flow.SearchReqeust) (tasks []flow.Task, err error)
}