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
	Image          string      `json:"image"  `
	UserID         int         `json:"user_id"  example:"1"`
	Description    string      `json:"description"  `
	Product        string      `json:"product" `
	AnalyzerConfig interface{} `json:"analyzer_config"  ` //配置文件，覆盖原有的
}

type UpdateEgnineReq struct {
	Image          string      `json:"image,omitempty" `
	UserID         int         `json:"user_id,omitempty"  `
	Description    string      `json:"description,omitempty" `
	Product        string      `json:"product,omitempty"  `
	AnalyzerConfig interface{} `json:"analyzer_config,omitempty" ` //配置文件，覆盖原有的
}

type StartEngineReq struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type StopEngineReq struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type RemoveEngineReq struct {
	ID string `uri:"id" json:"id" binding:"required"`
}

type EngineQuery struct {
	Image   string `json:"image,omitempty"  `
	UserID  int    `json:"user_id,omitempty" `
	Product string `json:"product,omitempty"  `
	Status  string `json:"status,omitempty" ` //enginestatus
}

type GetEngineReq struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	Image   string `json:"image,omitempty"  `
	UserID  int    `json:"user_id,omitempty" `
	Product string `json:"product,omitempty"  `
	Status  string `json:"status,omitempty" ` //enginestatus
}

type GetEngineRes struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
	Data  []bproto.EngineDeployInfo
}
