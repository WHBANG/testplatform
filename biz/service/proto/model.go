package proto

import (
	bproto "git.supremind.info/testplatform/biz/proto"
)

/*
	ModelID   int    `json:"model_id" bson:"model_id"`
	ModelName string `json:"model_name" bson:"model_name"`
	ModelType string `json:"model_type" bson:"model_type"`
	ModelURL  string `json:"model_url" bson:"model_url"`
*/
type UpdateModelReq struct {
	ModelID      int    `json:"model_id"`
	ModelType    string `json:"model_type,omitempty"`
	ModelName    string `json:"model_name,omitempty"`
	ModelURL     string `json:"model_url,omitempty"`
	ModelVersion string `json:"model_version"`
	ModelUser    string `json:"model_user"`
	Description  string `json:"description"`
}

type InsertModelReq struct {
	ModelID      int    `json:"model_id"`
	ModelType    string `json:"model_type"`
	ModelName    string `json:"model_name"`
	ModelURL     string `json:"model_url"`
	ModelVersion string `json:"model_version"`
	ModelUser    string `json:"model_user"`
	Description  string `json:"description"`
}

type DeleteModelReq struct {
	ModelID int `json:"model_id"`
}

type GetModelsByModelTypeReq struct {
	ModelType string `json:"model_type"`
}

type GetModelRes struct {
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	Total int            `json:"total"`
	Data  []bproto.Model `json:"data"`
}

type ModelInfo struct {
	ModelID      int       `json:"model_id" bson:"model_id"`
	ModelName    string    `json:"model_name" bson:"model_name"`
	ModelType    string    `json:"model_type" bson:"model_type"`
	ModelURL     string    `json:"model_url" bson:"model_url"`
	Accuracy     string    `bson:"accuracy" json:"accuracy"`
	CompleteTime LocalTime `bson:"complate_time" json:"complete_time"`
	Username     string    `bson:"username" json:"username"`
	Speed        string    `bson:"speed" json:"speed"`
}

type ProductInfo struct {
	ProductName string      `json:"product_name" bson:"product_name"`
	ProductID   int         `json:"product_id" bson:"product_id"`
	Models      []ModelInfo `json:"models" bson:"models"`
	//Models interface{} `bson:"models"`
}

type GetModelsRes []ProductInfo
