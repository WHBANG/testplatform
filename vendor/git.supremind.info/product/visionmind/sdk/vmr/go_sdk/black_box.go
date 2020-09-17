package client

import (
	"fmt"
	"time"

	"git.supremind.info/product/visionmind/com/black_box"
	"git.supremind.info/product/visionmind/com/go_sdk"
)

type BlackBoxClient struct {
	Host string
}

var _ go_sdk.IBlackBoxClient = &BlackBoxClient{}

func NewBlackBoxClient(host string) go_sdk.IBlackBoxClient {
	return &BlackBoxClient{
		Host: host,
	}
}

func (cli *BlackBoxClient) SetEvent(event *black_box.Event) (err error) {
	var (
		url  = fmt.Sprintf("http://%s/v1/event/collection", cli.Host)
		resp interface{}
	)

	return PostJson(url, *event, &resp, time.Second)
}

func (cli *BlackBoxClient) GetEvent(req *black_box.EventRequest) (resp black_box.EventResponse, err error) {
	var (
		url = fmt.Sprintf("http://%s/v1/event/collection?%s", cli.Host, structToValues(req).Encode())
	)

	resp = black_box.EventResponse{}
	err = GetJson(url, &resp, time.Second)

	return
}
