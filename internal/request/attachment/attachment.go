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
