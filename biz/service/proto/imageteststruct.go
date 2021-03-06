package proto

import (
	bproto "git.supremind.info/testplatform/biz/proto"
)

type InsertImageReq struct {
	Image       string           `json:"image"`
	UserID      int              `json:"user_id"`
	Description string           `json:"description"`
	Product     string           `json:"product"`
	Models      []bproto.Model   `json:"models" bson:"models"`
	Type        bproto.ImageType `json:"type,omitempty"`
}

type UpdateImageReq struct {
	Image       string           `json:"image,omitempty"`
	UserID      int              `json:"user_id,omitempty"`
	Description string           `json:"description,omitempty"`
	Product     string           `json:"product,omitempty"`
	Models      []bproto.Model   `json:"models,omitempty"`
	Type        bproto.ImageType `json:"type,omitempty"`
}

type ImageQuery struct {
	Image       string `json:"image,omitempty"  `
	UserID      int    `json:"user_id,omitempty" `
	Product     string `json:"product,omitempty"  `
	Description string `json:"description,omitempty" `
}

type GetImageReq struct {
	Page        int    `form:"page" json:"page"`
	Size        int    `form:"size" json:"size"`
	Image       string `form:"image" json:"image,omitempty"  `
	UserID      int    `form:"user_id" json:"user_id,omitempty" `
	Product     string `form:"product" json:"product,omitempty"  `
	Description string `form:"description" json:"description,omitempty" `
}

type GetImageRes struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
	Data  []bproto.ImageInfo
}

type CreateImageReq struct {
	AnalyzerIOBaseImage string `json:"analyzer_io_base_image"`
	ModelConfig         string `json:"model_config"`
	ImageName           string `json:"image_name" `
	Product      string `json:"product" `
	Description  string `json:"description"`
	AnalyzerType string `json:"analyzer_type"`
}
