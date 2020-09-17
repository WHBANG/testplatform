package client

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"git.supremind.info/product/visionmind/com/flow"
	"git.supremind.info/product/visionmind/com/go_sdk"
)

type TaskMgmClient struct {
	Host string
}

var _ go_sdk.ITaskMgmClient = &TaskMgmClient{}

func NewTaskMgmClient(host string) go_sdk.ITaskMgmClient {
	return &TaskMgmClient{
		Host: host,
	}
}

type commonResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (cli *TaskMgmClient) CreateTask(task *flow.Task) (err error) {
	var (
		url  = fmt.Sprintf("http://%s/v1/task", cli.Host)
		resp interface{}
	)

	err = PostJson(url, *task, &resp)
	if err != nil {
		return err
	}
	return nil
}

func (cli *TaskMgmClient) DeleteTask(taskId string) (err error) {
	var (
		url  = fmt.Sprintf("http://%s/v1/task/del/%s", cli.Host, taskId)
		resp commonResp
		body interface{}
	)

	err = PostJson(url, body, &resp)
	if err != nil {
		return err
	}
	return nil
}

func (cli *TaskMgmClient) UpdateTask(taskId string, req flow.UpdateTaskReq) (task flow.Task, err error) {
	var (
		url = fmt.Sprintf("http://%s/v1/task/update/%s", cli.Host, taskId)
	)

	err = PostJson(url, req, &task)
	if err != nil {
		return
	}
	return
}

func (cli *TaskMgmClient) StartTask(taskId string) (err error) {
	var (
		url  = fmt.Sprintf("http://%s/v1/task/start/%s", cli.Host, taskId)
		resp commonResp
		body interface{}
	)

	err = PostJson(url, body, &resp)
	if err != nil {
		return err
	}
	return nil
}

func (cli *TaskMgmClient) StopTask(taskId string) (err error) {
	var (
		url  = fmt.Sprintf("http://%s/v1/task/stop/%s", cli.Host, taskId)
		resp commonResp
		body interface{}
	)

	err = PostJson(url, body, &resp)
	if err != nil {
		return err
	}
	return nil
}

func (cli *TaskMgmClient) GetTask(taskId string) (task flow.Task, err error) {
	var (
		url = fmt.Sprintf("http://%s/v1/task/%s", cli.Host, taskId)
	)

	err = GetJson(url, &task)
	if err != nil {
		return
	}

	return
}

func (cli *TaskMgmClient) SearchTask(req flow.SearchReqeust) (tasks []flow.Task, err error) {
	var (
		url = fmt.Sprintf("http://%s/v1/task?%s", cli.Host, structToValues(&req).Encode())
	)

	err = GetJson(url, &tasks)
	if err != nil {
		return nil, err
	}

	return
}

func structToValues(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	for i := 0; i < iVal.NumField(); i++ {
		jsonTagName := iVal.Type().Field(i).Tag.Get("json")
		jsonTagName = strings.Split(jsonTagName, ",")[0]
		if jsonTagName == "" {
			continue
		}
		switch iVal.Field(i).Kind() {
		case reflect.Slice:
			for j := 0; j < iVal.Field(i).Len(); j++ {
				values.Add(jsonTagName, fmt.Sprint(iVal.Field(i).Index(j)))
			}
		default:
			value := fmt.Sprint(iVal.Field(i))
			if value == "" {
				continue
			}
			values.Set(jsonTagName, value)
		}
	}
	return
}
