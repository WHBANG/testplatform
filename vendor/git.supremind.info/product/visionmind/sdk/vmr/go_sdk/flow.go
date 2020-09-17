package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"git.supremind.info/product/visionmind/com/flow"
	"git.supremind.info/product/visionmind/com/go_sdk"
)

type FlowClient struct {
	Host string
}

var _ go_sdk.IFlowClient = &FlowClient{}

func NewFlowClient(host string) go_sdk.IFlowClient {
	return &FlowClient{
		Host: host,
	}
}

func (cli *FlowClient) TaskTry(ctx context.Context, leaseType string, region string, holder string) (*flow.Task, error) {
	url := fmt.Sprintf("http://%s/v1/task/lease/try/%s", cli.Host, holder)

	var task flow.Task
	req := flow.AcquireTaskReq{
		Type:   leaseType,
		Region: region,
	}
	err := PostJson(url, req, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (cli *FlowClient) TaskTryWithWeight(ctx context.Context, leaseType string, region string, holder string, weight int) (*flow.Task, error) {
	url := fmt.Sprintf("http://%s/v1/task/lease/try_with_weight/%s/%d", cli.Host, holder, weight)

	var task flow.Task
	req := flow.AcquireTaskReq{
		Type:   leaseType,
		Region: region,
	}
	err := PostJson(url, req, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (cli *FlowClient) TaskHold(ctx context.Context, id string, ver int64, holder string, extra ...string) (*flow.Task, error) {
	url := fmt.Sprintf("http://%s/v1/task/lease/hold/%s/%d/%s", cli.Host, id, ver, holder)

	task := flow.Task{}
	var status, errorMsg string
	if len(extra) > 0 && extra[0] != "" {
		status = extra[0]
	}
	if len(extra) > 1 && extra[1] != "" {
		errorMsg = extra[1]
	}

	err := PostJson(url, struct {
		Status   string `json:"status"`
		ErrorMsg string `json:"error_msg"`
	}{
		Status:   status,
		ErrorMsg: errorMsg,
	}, &task)
	if err != nil {
		return &task, err
	}

	return &task, nil
}

func (cli *FlowClient) TaskRelease(ctx context.Context, id string, ver int64, extra ...string) error {
	url := fmt.Sprintf("http://%s/v1/task/lease/release/%s/%d", cli.Host, id, ver)

	var errorMsg string
	if len(extra) > 0 && extra[0] != "" {
		errorMsg = extra[0]
	}

	err := PostJson(url, struct {
		ErrorMsg string `json:"error_msg"`
	}{
		ErrorMsg: errorMsg,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

func (cli *FlowClient) Push(ctx context.Context, topic string, body interface{}) (int64, error) {
	url := fmt.Sprintf("http://%s/v1/msg/%s", cli.Host, topic)

	ret := struct {
		ID int64 `json:"id"`
	}{}
	err := PostJson(url, body, &ret)
	if err != nil {
		return 0, err
	}

	return ret.ID, nil
}

func (cli *FlowClient) PullByID(ctx context.Context, topic string, id int64) (*flow.Message, error) {
	url := fmt.Sprintf("http://%s/v1/msg/id/%s/%d", cli.Host, topic, id)

	var msg flow.Message
	err := GetJson(url, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (cli *FlowClient) PullBatch(ctx context.Context, topic string, begid, endid int64) ([]flow.Message, error) {
	url := fmt.Sprintf("http://%s/v1/msg/batch/%s/%d/%d", cli.Host, topic, begid, endid)

	var msgs []flow.Message
	err := GetJson(url, &msgs)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (cli *FlowClient) Pull(ctx context.Context, topic, group string) (*flow.Message, error) {
	return cli.PullSince(ctx, topic, group, time.Time{})
}

func (cli *FlowClient) PullSince(ctx context.Context, topic, group string, since time.Time) (*flow.Message, error) {
	var sinceUnix int64
	if !since.IsZero() {
		sinceUnix = since.Unix()
	}
	url := fmt.Sprintf("http://%s/v1/msg/group/%s/%s?since=%d", cli.Host, topic, group, sinceUnix)

	var msg flow.Message
	err := GetJson(url, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (cli *FlowClient) ConfigCheck(ctx context.Context, id string) (*flow.Config, error) {
	url := fmt.Sprintf("http://%s/v1/config/check/%s/0", cli.Host, id)

	var conf flow.Config
	err := GetJson(url, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func (cli *FlowClient) ConfigWatch(ctx context.Context, conf *flow.Config, callback func(context.Context, *flow.Config) bool) {

	curConf := conf
	ticker := time.NewTicker(time.Minute)

	for {
		select {
		// 结束
		case <-ctx.Done():
			log.Println("ConfigWatch Done")
			return

		case <-ticker.C:
			url := fmt.Sprintf("http://%s/v1/config/check/%s/%d", cli.Host, curConf.ID, curConf.Ver)
			var tmpConf flow.Config
			err := GetJson(url, &tmpConf)
			if err == nil && tmpConf.Ver > curConf.Ver {
				log.Printf("ConfigWatch: %+v", tmpConf)
				if callback(ctx, &tmpConf) {
					curConf = &tmpConf
				}
			}
		}
	}
}

func (cli *FlowClient) GetDevices(ctx context.Context) ([]flow.Device, error) {
	url := fmt.Sprintf("http://%s/v1/device", cli.Host)

	var devices []flow.Device
	err := GetJson(url, &devices)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (cli *FlowClient) GetDevice(ctx context.Context, deviceID string) (*flow.Device, error) {
	url := fmt.Sprintf("http://%s/v1/device/%s", cli.Host, deviceID)

	var device flow.Device
	err := GetJson(url, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (cli *FlowClient) GetCameras(ctx context.Context, options go_sdk.CameraOptions) ([]flow.Camera, error) {
	url := fmt.Sprintf("http://%s/v1/device/camera?deleted_include=%t", cli.Host, options.DeleteInclude)

	var cameras []flow.Camera
	err := GetJson(url, &cameras)
	if err != nil {
		return nil, err
	}

	return cameras, nil
}

func (cli *FlowClient) GetDeviceStreamURL(ctx context.Context, deviceID string, streamID int, streamType string) (string, error) {
	url := fmt.Sprintf("http://%s/v1/device/%s/url/stream?stream_type=%s&stream_id=%d", cli.Host, deviceID, streamType, streamID)

	resp := struct {
		URL string `json:"url"`
	}{}
	err := GetJson(url, &resp)
	if err != nil {
		return "", err
	}

	return resp.URL, nil
}

func (cli *FlowClient) GetDeviceVChannelURL(ctx context.Context, deviceID string, target string) (url string, err error) {
	url = fmt.Sprintf("http://%s/v1/device/%s/url/vchannel?target=%s", cli.Host, deviceID, target)

	resp := struct {
		URL string `json:"url"`
	}{}
	err = GetJson(url, &resp)
	if err != nil {
		return
	}

	return resp.URL, nil
}

func (cli *FlowClient) GetDeviceSnap(ctx context.Context, deviceID string, streamType string) (*flow.Snap, error) {
	url := fmt.Sprintf("http://%s/v1/device/%s/snap?stream_type=%s", cli.Host, deviceID, streamType)

	var resp flow.Snap
	err := GetJson(url, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (cli *FlowClient) GetDownloadUrl(ctx context.Context, deviceID string, startTime, endTime time.Time, timeout time.Duration) (string, error) {
	if !startTime.Before(endTime) {
		return "", errors.New("invalid startTime and endTime")
	}
	if timeout < time.Second {
		return "", errors.New("timeout must be at least 1 second")
	}

	resp := struct {
		URL string `json:"url"`
	}{}
	start := startTime.Format("20060102150405")
	end := endTime.Format("20060102150405")
	url := fmt.Sprintf("http://%s/v1/device/%s/url/download?start_time=%s&end_time=%s&timeout=%d",
		cli.Host, deviceID, start, end, timeout/time.Second)
	err := GetJson(url, &resp)
	if err != nil {
		return "", err
	}

	return resp.URL, nil
}

func (cli *FlowClient) CreateDevice(ctx context.Context, device interface{}) (*flow.Device, error) {
	url := fmt.Sprintf("http://%s/v1/device", cli.Host)

	var resp flow.Device
	err := PostJson(url, device, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (cli *FlowClient) DeleteDevice(ctx context.Context, deviceIDs []string) ([]string, error) {
	url := fmt.Sprintf("http://%s/v1/device", cli.Host)

	params := struct {
		IDs []string `json:"ids"`
	}{
		IDs: deviceIDs,
	}

	var resp []string
	err := DeleteJson(url, params, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (cli *FlowClient) FetchTopics(ctx context.Context, filters []string) ([]string, error) {
	topics := strings.Join(filters, ",")
	url := fmt.Sprintf("http://%s/v1/msg/topics/%s", cli.Host, topics)
	var resp []string
	err := GetJson(url, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *FlowClient) UpdateConfig(ctx context.Context, configId string, data interface{}) (err error) {
	url := fmt.Sprintf("http://%s/v1/config/set/%s", cli.Host, configId)
	println(url)
	var resp commonResp

	err = PostJson(url, data, &resp)
	return

}

func PostJson(url string, body interface{}, result interface{}, timeout ...time.Duration) error {

	client := http.DefaultClient
	if len(timeout) > 0 {
		client = &http.Client{
			Timeout: timeout[0],
		}
	}

	reqbs, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(reqbs))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		err := fmt.Errorf("err, StatusCode = %d", resp.StatusCode)
		return err
	}

	if result != nil {
		respbs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var resp struct {
			Code    int             `json:"code"`
			Message string          `json:"message"`
			Data    json.RawMessage `json:"data"`
		}
		err = json.Unmarshal(respbs, &resp)
		if err != nil {
			return err
		}
		if resp.Message != "" {
			return errors.New(resp.Message)
		}

		if resp.Data != nil {
			err = json.Unmarshal(resp.Data, result)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetJson(url string, result interface{}, timeout ...time.Duration) error {

	client := http.DefaultClient
	if len(timeout) > 0 {
		client = &http.Client{
			Timeout: timeout[0],
		}
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		err := fmt.Errorf("err, StatusCode = %d", resp.StatusCode)
		return err
	}

	if result != nil {
		respbs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var resp struct {
			Code    int             `json:"code"`
			Message string          `json:"message"`
			Data    json.RawMessage `json:"data"`
		}
		err = json.Unmarshal(respbs, &resp)
		if err != nil {
			return err
		}
		if resp.Message != "" {
			return errors.New(resp.Message)
		}

		err = json.Unmarshal(resp.Data, result)

		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteJson(url string, body interface{}, result interface{}, timeout ...time.Duration) error {

	client := http.DefaultClient
	if len(timeout) > 0 {
		client = &http.Client{
			Timeout: timeout[0],
		}
	}

	reqbs, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(reqbs))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		err := fmt.Errorf("err, StatusCode = %d", resp.StatusCode)
		return err
	}

	if result != nil {
		respbs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var resp struct {
			Code    int             `json:"code"`
			Message string          `json:"message"`
			Data    json.RawMessage `json:"data"`
		}
		err = json.Unmarshal(respbs, &resp)
		if err != nil {
			return err
		}
		if resp.Message != "" {
			return errors.New(resp.Message)
		}

		if resp.Data != nil {
			err = json.Unmarshal(resp.Data, result)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
