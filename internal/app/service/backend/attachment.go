package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/pkg/pagination"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/request/attachment"
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
	attachments, count, err := dao.CAttachment.FindByPage(parsePage.GetOffset(), parsePage.GetLimit())
	if err != nil {
		return nil, err
	}
	parsePage.SetCount(count).SetList(attachments).GetLastPage()
	return &parsePage, nil
}
