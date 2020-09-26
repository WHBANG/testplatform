package proto

import (
	bproto "git.supremind.info/testplatform/biz/proto"
)

type InsertAnalyzerTypeReq struct {
	AnalyzerType  string            `json:"analyzer_type" bson:"analyzer_type"`
	ModelNameList bproto.ModelNames `json:"model_name_list" bson:"model_name_list"`
}

type GetAnzlyzerTypeModelNameReq struct {
	AnalyzerType string `json:"analyzer_type" bson:"analyzer_type"`
}

type GetAnzlyzerTypeModelNameRes struct {
	ModelNameList bproto.ModelNames `json:"model_name_list" bson:"model_name_list"`
}
