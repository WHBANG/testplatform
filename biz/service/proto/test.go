package proto

import (
	bproto "git.supremind.info/testplatform/biz/proto"
)

const (
	DefaultErrorCode = 1
)

type CommonRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CreateTestReq struct {
}

type CreateEngineReq struct {
	bproto.EngineDeployInfo
}

type UpdateEgnineReq struct {
	bproto.EngineDeployInfo
}
type StartEngineReq struct {
	ID string `json:"id"`
}
