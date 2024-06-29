package backend

import (
	"errors"
	"strings"

	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/app/pkg/pagination"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/request/product"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
)

type BrandService struct {
	service.Service
}

var Brand = &BrandService{}

// Index ...
func (s *BrandService) Index(params product.IndexReq) (*pagination.PaginateResp, error) {
	parsePage := pagination.New().ParsePage(params.Page, params.Limit)
	productBrands, count, err := dao.ProductBrand.FindByPage(parsePage.GetOffset(), parsePage.GetLimit())
	if err != nil {
		return nil, err
	}
	parsePage.Total = count
	parsePage.GetLastPage()
	parsePage.List = productBrands
	return &parsePage, nil
}

// View ...
func (s *BrandService) View(params product.ViewReq) (*model.ProductBrand, error) {
	return dao.ProductBrand.Where(dao.ProductBrand.ID.Eq(params.ID)).First()
}

// Status ...
func (s *BrandService) Status(params product.StatusReq) error {
	if _, err := dao.ProductBrand.Where(dao.ProductBrand.ID.Eq(params.ID)).Update(dao.ProductBrand.Status, params.Status); err != nil {
		return err
	}
	return nil
}

// Delete ...
func (s *BrandService) Delete(params product.MultiDeleteReq) error {
	idList := make([]int64, 0)
	ids := strings.Split(params.IDs, ",")
	if len(ids) == 0 {
		return errors.New("请选择需要删除的记录")
	}
	for _, id := range ids {
		idList = append(idList, cast.ToInt64(id))
	}
	if _, err := dao.ProductBrand.Where(dao.ProductBrand.ID.In(idList...)).Delete(); err != nil {
		return err
	}
	return nil
}

// Create ...
func (s *BrandService) Create(params product.BrandCreateReq) error {
	var info model.ProductBrand
	if _, err := dao.ProductBrand.Where(dao.ProductBrand.BrandName.Eq(params.BrandName)).First(); err == nil {
		return errors.New("品牌名称已存在")
	}
	if err := copier.Copy(&info, params); err != nil {
		return err
	}
	return dao.ProductBrand.Create(&info)
}

// Update ...
func (s *BrandService) Update(params product.BrandUpdateReq) error {
	var info model.ProductBrand
	if _, err := dao.ProductBrand.Where(dao.ProductBrand.ID.Eq(params.ID)).First(); err != nil {
		return err
	}
	if _, err := dao.ProductBrand.Where(dao.ProductBrand.BrandName.Eq(params.BrandName), dao.ProductBrand.ID.Neq(params.ID)).First(); err == nil {
		return errors.New("品牌名称已存在")
	}
	if err := copier.Copy(&info, params); err != nil {
		return err
	}
	if _, err := dao.ProductBrand.Updates(&info); err != nil {
		return err
	}
	return nil
}

// List ...
func (s *BrandService) List() ([]*model.ProductBrand, error) {
	return dao.ProductBrand.Find()
}
