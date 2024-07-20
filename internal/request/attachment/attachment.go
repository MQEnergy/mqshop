package attachment

import "github.com/MQEnergy/mqshop/pkg/upload"

type IndexReq struct {
	Page   int    `form:"page" json:"page" validate:"required"`
	Limit  int    `form:"limit" json:"limit" validate:"required"`
	Search string `form:"search" json:"search"`
}

type UploadReq struct {
	upload.File
	FilePath string `form:"file_path"`
}

type CreateReq struct {
	UserID           int64  `form:"user_id" json:"user_id"`
	AttachName       string `form:"attach_name" json:"attach_name" validate:"required"`
	AttachOriginName string `form:"attach_origin_name" json:"attach_origin_name" validate:"required"`
	AttachURL        string `form:"attach_url" json:"attach_url" validate:"required"`
	AttachType       int8   `form:"attach_type" json:"attach_type" validate:"required"`
	AttachMinetype   string `form:"attach_minetype" json:"attach_minetype" validate:"required"`
	AttachExtension  string `form:"attach_extension" json:"attach_extension" validate:"required"`
	AttachSize       string `form:"attach_size" json:"attach_size" validate:"required"`
	Status           int8   `form:"status" json:"status" validate:"required"`
}
