package backend

import (
	"errors"
	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/app/pkg/pagination"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/request/product"
	"github.com/spf13/cast"
	"strings"
)

type AttrCateService struct {
	service.Service
}

var AttrCate = &AttrCateService{}

// Index ...
func (s *AttrCateService) Index(params product.IndexReq) (*pagination.PaginateResp, error) {
	var (
		attrList = make([]*product.AttrCate, 0)
	)

	parsePage := pagination.New().ParsePage(params.Page, params.Limit)
	result, count, err := dao.ProductGoodsAttrCate.FindByPage(parsePage.GetOffset(), parsePage.GetLimit())
	if err != nil {
		return nil, err
	}
	for _, cate := range result {
		attrNum, _ := dao.ProductGoodsAttr.Where(dao.ProductGoodsAttr.CateID.Eq(cate.ID), dao.ProductGoodsAttr.AttrType.Eq(1)).Count()
		paramsNum, _ := dao.ProductGoodsAttr.Where(dao.ProductGoodsAttr.CateID.Eq(cate.ID), dao.ProductGoodsAttr.AttrType.Eq(2)).Count()
		attrList = append(attrList, &product.AttrCate{
			ID:        cate.ID,
			CateName:  cate.CateName,
			Status:    cate.Status,
			ParamsNum: paramsNum,
			AttrNum:   attrNum,
			CreatedAt: cate.CreatedAt,
		})
	}
	parsePage.Total = count
	parsePage.GetLastPage()
	parsePage.List = attrList
	return &parsePage, nil
}

// View ...
func (s *AttrCateService) View(params product.ViewReq) (*model.ProductGoodsAttrCate, error) {
	return dao.ProductGoodsAttrCate.Where(dao.ProductGoodsAttrCate.ID.Eq(params.ID)).First()
}

// Create ...
func (s *AttrCateService) Create(params product.AttrCateCreateReq) error {
	if _, err := dao.ProductGoodsAttrCate.Where(dao.ProductGoodsAttrCate.CateName.Eq(params.CateName)).First(); err == nil {
		return errors.New("分类名称已存在")
	}
	return dao.ProductGoodsAttrCate.Create(&model.ProductGoodsAttrCate{
		CateName: params.CateName,
		Status:   params.Status,
	})
}

// Update ...
func (s *AttrCateService) Update(params product.AttrCateUpdateReq) error {
	if _, err := dao.ProductGoodsAttrCate.Where(dao.ProductGoodsAttrCate.ID.Eq(params.ID)).First(); err != nil {
		return errors.New("分类不存在")
	}
	if _, err := dao.ProductGoodsAttrCate.Where(dao.ProductGoodsAttrCate.CateName.Eq(params.CateName), dao.ProductGoodsAttrCate.ID.Neq(params.ID)).First(); err == nil {
		return errors.New("分类名称已存在")
	}
	_, err := dao.ProductGoodsAttrCate.Debug().Where(dao.ProductGoodsAttrCate.ID.Eq(params.ID)).Updates(map[string]interface{}{
		"cate_name": params.CateName,
		"status":    params.Status,
	})
	return err
}

// Delete ...
func (s *AttrCateService) Delete(params product.MultiDeleteReq) error {
	idList := make([]int64, 0)
	ids := strings.Split(params.IDs, ",")
	if len(ids) == 0 {
		return errors.New("请选择需要删除的记录")
	}
	for _, id := range ids {
		idList = append(idList, cast.ToInt64(id))
	}
	if _, err := dao.ProductGoodsAttrCate.Where(dao.ProductGoodsAttrCate.ID.In(idList...)).First(); err != nil {
		return errors.New("分类不存在")
	}
	_, err := dao.ProductGoodsAttrCate.Where(dao.ProductGoodsAttrCate.ID.In(idList...)).Delete()
	return err
}

// List ...
func (s *AttrCateService) List() ([]*model.ProductGoodsAttrCate, error) {
	return dao.ProductGoodsAttrCate.Where(dao.ProductGoodsAttrCate.Status.Eq(1)).Find()
}
