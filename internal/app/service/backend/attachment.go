package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/app/pkg/pagination"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/request/attachment"
	"github.com/MQEnergy/mqshop/pkg/upload"
	"github.com/spf13/cast"
)

type AttachmentService struct {
	service.Service
}

var Attachment = &AttachmentService{}

// Index ...
func (s *AttachmentService) Index(params attachment.IndexReq) (*pagination.PaginateResp, error) {
	var (
		parsePage = pagination.New().ParsePage(params.Page, params.Limit)
	)
	attachments, count, err := dao.CAttachment.Order(dao.CAttachment.ID.Desc()).FindByPage(parsePage.GetOffset(), parsePage.GetLimit())
	if err != nil {
		return nil, err
	}
	parsePage.SetCount(count).SetList(attachments).GetLastPage()
	return &parsePage, nil
}

func (s *AttachmentService) Create(userId string, params *upload.FileHeader) error {
	attachmentInfo := model.CAttachment{
		UserID:           cast.ToInt64(userId),
		AttachName:       params.Filename,
		AttachOriginName: params.OriginName,
		AttachURL:        params.FilePath,
		AttachType:       1,
		AttachMimetype:   params.MimeType,
		AttachExtension:  params.Extension,
		AttachSize:       cast.ToString(params.FileSize),
		Status:           1,
	}
	return dao.CAttachment.Create(&attachmentInfo)
}
