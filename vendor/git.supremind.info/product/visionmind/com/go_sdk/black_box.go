package go_sdk

import "git.supremind.info/product/visionmind/com/black_box"

type IBlackBoxClient interface {
	SetEvent(event *black_box.Event) (err error)
	GetEvent(req *black_box.EventRequest) (resp black_box.EventResponse, err error)
}
