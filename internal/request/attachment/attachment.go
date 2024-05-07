package attachment

import "github.com/MQEnergy/mqshop/pkg/upload"

type UploadReq struct {
	upload.File
	FilePath string `form:"file_path"`
}
