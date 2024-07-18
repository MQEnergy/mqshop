package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/app/pkg/pagination"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/request/product"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/helper"
	"github.com/goccy/go-json"
)

type ProductService struct {
	service.Service
}

var Product = &ProductService{}

// Index ...
func (s *ProductService) Index(params product.IndexReq) (*pagination.PaginateResp, error) {
	var (
		productList = make([]*product.ProductGoods, 0)
		parsePage   = pagination.New().ParsePage(params.Page, params.Limit)
	)

	productGoods, count, err := dao.ProductGoods.FindByPage(parsePage.GetOffset(), parsePage.GetLimit())
	if err != nil {
		return nil, err
	}
	for _, good := range productGoods {
		cateName := ""
		brandName := ""
		photoUrls := make([]string, 0)
		// 查找分类名称
		if cateInfo, err := dao.ProductCategory.Select(dao.ProductCategory.CateName).Where(dao.ProductCategory.ID.Eq(good.CateID)).First(); err == nil {
			cateName = cateInfo.CateName
		}
		// 查找品牌名称
		if brandInfo, err := dao.ProductBrand.Select(dao.ProductBrand.BrandName).Where(dao.ProductBrand.ID.Eq(good.BrandID)).First(); err == nil {
			brandName = brandInfo.BrandName
		}
		_ = json.Unmarshal([]byte(good.PhotoUrls), &photoUrls)
		productList = append(productList, &product.ProductGoods{
			ID:            good.ID,
			UUID:          good.UUID,
			GoodsTitle:    good.GoodsTitle,
			GoodsSubtitle: good.GoodsSubtitle,
			CateID:        good.CateID,
			CateName:      cateName,
			BrandID:       good.BrandID,
			BrandName:     brandName,
			AttrCateID:    good.AttrCateID,
			ThumbURL:      good.ThumbURL,
			PhotoUrls:     photoUrls,
			OriginPrice:   good.OriginPrice,
			PromotePrice:  good.PromotePrice,
			FinalPrice:    good.FinalPrice,
			TotalStock:    good.TotalStock,
			IsHot:         good.IsHot,
			IsNew:         good.IsNew,
			IsRecommend:   good.IsRecommend,
			SortOrder:     good.SortOrder,
			Status:        good.Status,
			CreatedAt:     good.CreatedAt,
		})
	}
	parsePage.SetCount(count).SetList(productList).GetLastPage()
	return &parsePage, nil
}

func (s *ProductService) Delete(params product.DeleteReq) error {
	//	查询
	productGoods, err := dao.ProductGoods.Where(dao.ProductGoods.ID.Eq(params.ID)).First()
	if err != nil {
		return err
	}
	//	开启事务
	q := dao.Use(vars.DB)
	return q.Transaction(func(tx *dao.Query) error {
		// 删除商品
		if _, err := dao.ProductGoods.Where(dao.ProductGoods.ID.Eq(params.ID)).Delete(); err != nil {
			return err
		}
		// 删除商品详情
		if _, err := dao.ProductGoodsInfo.Where(dao.ProductGoodsInfo.GoodsID.Eq(productGoods.ID)).Delete(); err != nil {
			return err
		}
		// 删除商品sku
		if _, err := dao.ProductGoodsSku.Where(dao.ProductGoodsSku.GoodsID.Eq(productGoods.ID)).Delete(); err != nil {
			return err
		}
		// 删除商户属性关联表
		if _, err := dao.ProductGoodsAttrRelation.Where(dao.ProductGoodsAttrRelation.GoodsID.Eq(productGoods.ID)).Delete(); err != nil {
			return err
		}
		// 删除标签关联表
		if _, err := dao.ProductGoodsTagRelation.Where(dao.ProductGoodsTagRelation.GoodsID.Eq(productGoods.ID)).Delete(); err != nil {
			return err
		}
		// 删除商品评论表
		if _, err := dao.ProductComment.Where(dao.ProductComment.GoodsID.Eq(productGoods.ID)).Delete(); err != nil {
			return err
		}
		return nil
	})
}

// Create
//
//	@Description: Todo
//	@receiver s
//	@param params
//	@return error
func (s *ProductService) Create(params product.CreateReq) error {
	// imageList := params.ImageList
	//	开启事务
	q := dao.Use(vars.DB)
	return q.Transaction(func(tx *dao.Query) error {
		// 创建商品
		if err := tx.ProductGoods.Create(&model.ProductGoods{
			UUID:          helper.GenerateBaseSnowId(0, nil),
			GoodsTitle:    params.GoodsTitle,
			GoodsSubtitle: params.GoodsSubtitle,
			CateID:        params.CateID,
			BrandID:       params.BrandID,
			AttrCateID:    params.AttrCateID,
			ThumbURL:      params.ThumbList,
			PhotoUrls:     params.ThumbList,
			OriginPrice:   params.OriginPrice,
			PromotePrice:  params.PromotePrice,
			FinalPrice:    params.FinalPrice,
			TotalStock:    params.TotalStock,
			IsHot:         params.IsHot,
			IsNew:         params.IsNew,
			IsRecommend:   params.IsRecommend,
			SortOrder:     params.SortOrder,
			Status:        params.Status,
		}); err != nil {
			return err
		}

		return nil
	})
}

// Update
//
//	@Description: Todo
//	@receiver s
//	@param params
//	@return error
func (s *ProductService) Update(params product.UpdateReq) error {
	// 开启事务
	q := dao.Use(vars.DB)
	return q.Transaction(func(tx *dao.Query) error {
		// 更新商品
		if _, err := tx.ProductGoods.Where(dao.ProductGoods.ID.Eq(params.ID)).Updates(&model.ProductGoods{
			GoodsTitle:    params.GoodsTitle,
			GoodsSubtitle: params.GoodsSubtitle,
			CateID:        params.CateID,
			BrandID:       params.BrandID,
			AttrCateID:    params.AttrCateID,
			ThumbURL:      params.ThumbList,
			PhotoUrls:     params.ThumbList,
			OriginPrice:   params.OriginPrice,
			PromotePrice:  params.PromotePrice,
			FinalPrice:    params.FinalPrice,
			TotalStock:    params.TotalStock,
			IsHot:         params.IsHot,
			IsNew:         params.IsNew,
			IsRecommend:   params.IsRecommend,
			SortOrder:     params.SortOrder,
			Status:        params.Status,
		}); err != nil {
			return err
		}

		return nil
	})
}

// View
//
//	@Description: Todo
//	@receiver s
//	@param id
//	@return *model.ProductGoods
//	@return error
func (s *ProductService) View(params product.ViewReq) (*model.ProductGoods, error) {
	return dao.ProductGoods.Where(dao.ProductGoods.ID.Eq(params.ID)).First()
}

// Status
//
//	@Description: Todo
//	@receiver s
//	@param params
//	@return error
func (s *ProductService) Status(params product.StatusReq) error {
	if _, err := dao.ProductGoods.Where(dao.ProductGoods.ID.Eq(params.ID)).Update(dao.ProductGoods.Status, params.Status); err != nil {
		return err
	}
	return nil
}
