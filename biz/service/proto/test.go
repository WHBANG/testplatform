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
	Image          string `json:"image" example:"reg.supremind.info/products/vas/vas-app/analyzer:cuda10-20200604-16108" `
	UserID         int    `json:"user_id"  example:"1"`
	Description    string `json:"description"  example:"test"`
	Product        string `json:"product"  example:"massiveflow"`
	AnalyzerConfig string `json:"analyzer_config"   example:"{}"` //配置文件，覆盖原有的
}

type UpdateEgnineReq struct {
	Image          string `json:"image,omitempty" example:"reg.supremind.info/products/vas/vas-app/analyzer:cuda10-20200604-16108"`
	UserID         int    `json:"user_id,omitempty"  example:"1"`
	Description    string `json:"description,omitempty" example:"test"`
	Product        string `json:"product,omitempty"  example:"massiveflow"`
	AnalyzerConfig string `json:"analyzer_config,omitempty" example:"{}"` //配置文件，覆盖原有的
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
