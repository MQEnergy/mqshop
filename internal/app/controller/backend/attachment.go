package backend

import (
	"fmt"
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend"
	"github.com/MQEnergy/mqshop/internal/request/attachment"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/response"
	"github.com/MQEnergy/mqshop/pkg/upload"
	"github.com/gofiber/fiber/v2"
)

type AttachmentController struct {
	controller.Controller
}

var Attachment = &AttachmentController{}

// Index ...
func (c *AttachmentController) Index(ctx *fiber.Ctx) error {
	var params attachment.IndexReq
	if err := c.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	fmt.Println(params)
	attachmentList, err := backend.Attachment.Index(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", attachmentList)
}

// Upload 上传资源
func (c *AttachmentController) Upload(ctx *fiber.Ctx) error {
	reqParams := &attachment.UploadReq{}
	if err := c.Validate(ctx, reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	fileName, err := ctx.FormFile("file")
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	reqParams.FileName = fileName
	fileHeader, err := upload.New(0, []string{}).UploadToLocal(vars.Config, reqParams.FileName, reqParams.FilePath)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	userId := ctx.GetRespHeader("uid")
	if err := backend.Attachment.Create(userId, fileHeader); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", fileHeader)
}
