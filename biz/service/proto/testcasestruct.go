package proto

import (
	bproto "git.supremind.info/testplatform/biz/proto"
)

type InsertTestCaseReq struct {
	TestData    interface{} `json:"test_data"`
	UserID      int         `json:"user_id"`
	Description string      `json:"description"`
}

type UpdateTestCaseReq struct {
	TestData    interface{} `json:"test_data,omitempty"`
	UserID      int         `json:"user_id,omitempty"`
	Description string      `json:"description,omitempty"`
}

type GetTestCaseReq struct {
	Page        int         `form:"page" json:"page"`
	Size        int         `form:"size" json:"size"`
	TestData    interface{} `json:"test_data,omitempty"`
	UserID      int         `form:"user_id" json:"user_id,omitempty"`
	Description string      `form:"description" json:"description,omitempty"`
}

type GetTestCaseRes struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
	Data  []bproto.TestCaseInfo
}
