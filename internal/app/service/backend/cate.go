package backend

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/request/product"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
)

type CateService struct {
	service.Service
}

var Cate = &CateService{}

// CateNode Define a struct to hold the hierarchical data
type CateNode struct {
	ID           int64  `json:"id"`
	CateName     string `json:"cate_name"`      // 分类名称
	BannerCateID int64  `json:"banner_cate_id"` // 轮播图分类ID
	CateURL      string `json:"cate_url"`       // 分类图标
	CateDesc     string `json:"cate_desc"`      // 分类描述
	SortOrder    int    `json:"sort_order"`     // 排序
	ParentID     int64  `json:"parent_id"`      // 父级ID
	Path         string `json:"path"`           // ID路径 1-2-3-4...
	IsHot        int8   `json:"is_hot"`         // 是否热门 1：热门 0：非热门
	IsIndex      int8   `json:"is_index"`       // 是否首页 1：首页 0：非首页
	Status       int8   `json:"status"`         // 状态 1：正常 0：下架
	CreatedAt    int64  `json:"created_at"`     //
	Label        string `json:"label"`
	Level        int    `json:"level"`
}

// CateItem ...
type CateItem struct {
	ID           int64       `json:"id"`
	CateName     string      `json:"cate_name"`      // 分类名称
	BannerCateID int64       `json:"banner_cate_id"` // 轮播图分类ID
	CateURL      string      `json:"cate_url"`       // 分类图标
	CateDesc     string      `json:"cate_desc"`      // 分类描述
	SortOrder    int         `json:"sort_order"`     // 排序
	ParentID     int64       `json:"parent_id"`      // 父级ID
	Path         string      `json:"path"`           // ID路径 1-2-3-4...
	IsHot        int8        `json:"is_hot"`         // 是否热门 1：热门 0：非热门
	IsIndex      int8        `json:"is_index"`       // 是否首页 1：首页 0：非首页
	Status       int8        `json:"status"`         // 状态 1：正常 0：下架
	CreatedAt    int64       `json:"created_at"`     //
	Children     []*CateItem `json:"children"`
}

// SortCategoriesTree sorts the categories
func SortCategoriesTree(categories []*model.ProductCategory, parentID int64, level int) []*CateNode {
	var sorted []*CateNode
	for _, category := range categories {
		if category.ParentID == parentID {
			menuItem := &CateNode{
				ID:           category.ID,
				CateName:     category.CateName,
				BannerCateID: category.BannerCateID,
				CateURL:      category.CateURL,
				CateDesc:     category.CateDesc,
				SortOrder:    category.SortOrder,
				ParentID:     category.ParentID,
				Path:         category.Path,
				IsHot:        category.IsHot,
				IsIndex:      category.IsIndex,
				Status:       category.Status,
				CreatedAt:    category.CreatedAt,
				Label:        fmt.Sprintf("%s%s", strings.Repeat("- ", level), category.CateName),
				Level:        level,
			}
			sorted = append(sorted, menuItem)
			sorted = append(sorted, SortCategoriesTree(categories, category.ID, level+1)...)
		}
	}
	return sorted
}

// BuildHierarchy Function to build the hierarchical structure
func BuildHierarchy(resources []*model.ProductCategory) []*CateItem {
	menuMap := make(map[int64]*CateItem)
	var roots []*CateItem

	for _, source := range resources {
		node := &CateItem{
			ID:           source.ID,
			CateName:     source.CateName,
			BannerCateID: source.BannerCateID,
			CateURL:      source.CateURL,
			CateDesc:     source.CateDesc,
			SortOrder:    source.SortOrder,
			ParentID:     source.ParentID,
			Path:         source.Path,
			IsHot:        source.IsHot,
			IsIndex:      source.IsIndex,
			Status:       source.Status,
			CreatedAt:    source.CreatedAt,
			Children:     []*CateItem{},
		}
		menuMap[source.ID] = node
	}
	// Build the hierarchical structure by linking parent and child nodes
	for _, source := range resources {
		node := menuMap[source.ID]
		if source.ParentID == 0 {
			roots = append(roots, node)
		} else {
			parentNode := menuMap[source.ParentID]
			if parentNode != nil {
				parentNode.Children = append(parentNode.Children, node)
			}
		}
	}
	return roots
}

// Index ...
func (s *CateService) Index(params product.IndexReq) ([]*CateItem, error) {
	categories, err := dao.ProductCategory.Order(dao.ProductCategory.ID.Asc(), dao.ProductCategory.SortOrder.Desc()).Find()
	if err != nil {
		return nil, err
	}
	return BuildHierarchy(categories), nil
}

// Create ...
func (s *CateService) Create(params product.CateCreateReq) error {
	var (
		q    = dao.Use(vars.DB)
		info = model.ProductCategory{
			CateName:     params.CateName,
			BannerCateID: 0,
			CateURL:      params.CateUrl,
			CateDesc:     params.CateDesc,
			SortOrder:    params.SortOrder,
			ParentID:     params.ParentID,
			Path:         "",
			IsHot:        params.IsHot,
			IsIndex:      params.IsIndex,
			Status:       params.Status,
		}
	)
	if _, err := dao.ProductCategory.Where(dao.ProductCategory.CateName.Eq(params.CateName)).First(); err == nil {
		return errors.New("分类名称已存在")
	}
	return q.Transaction(func(tx *dao.Query) error {
		if err := tx.ProductCategory.Save(&info); err != nil {
			return err
		}
		// 更新path路径
		if info.ParentID != 0 {
			categoryInfo, err := dao.ProductCategory.Where(dao.ProductCategory.ID.Eq(params.ParentID)).First()
			if err == nil {
				info.Path = categoryInfo.Path + "-" + cast.ToString(info.ID)
			} else {
				info.Path = cast.ToString(info.ID)
			}
		} else {
			info.Path = cast.ToString(info.ID)
		}
		if err := tx.ProductCategory.Save(&info); err != nil {
			return err
		}
		return nil
	})
}

// Update ...
func (s *CateService) Update(params product.CateUpdateReq) error {
	var (
		q    = dao.Use(vars.DB)
		info model.ProductCategory
	)
	if err := copier.Copy(&info, params); err != nil {
		return err
	}
	return q.Transaction(func(tx *dao.Query) error {
		if err := tx.ProductCategory.Save(&info); err != nil {
			return err
		}
		// 更新path路径
		if info.ParentID != 0 {
			categoryInfo, err := dao.ProductCategory.Where(dao.ProductCategory.ID.Eq(params.ParentID)).First()
			if err == nil {
				info.Path = categoryInfo.Path + "-" + cast.ToString(info.ID)
			} else {
				info.Path = cast.ToString(info.ID)
			}
		} else {
			info.Path = cast.ToString(info.ID)
		}
		if err := tx.ProductCategory.Save(&info); err != nil {
			return err
		}
		return nil
	})
}

// Delete ...
func (s *CateService) Delete(params product.MultiDeleteReq) error {
	idList := make([]int64, 0)
	ids := strings.Split(params.IDs, ",")
	if len(ids) == 0 {
		return errors.New("请选择需要删除的记录")
	}
	for _, id := range ids {
		idList = append(idList, cast.ToInt64(id))
	}
	childCateList, err := dao.ProductCategory.Where(dao.ProductCategory.ParentID.In(idList...)).Find()
	if err != nil {
		return err
	}
	for _, category := range childCateList {
		idList = append(idList, category.ID)
	}
	// 查询该分类下的商品
	productList, _ := dao.ProductGoods.Where(dao.ProductGoods.CateID.In(idList...)).Find()
	if len(productList) > 0 {
		return errors.New("该分类下有商品，需要先删除商品")
	}
	if _, err := dao.ProductCategory.Where(dao.ProductCategory.ID.In(idList...)).Delete(); err != nil {
		return err
	}
	return nil
}

// View ...
func (s *CateService) View(params product.ViewReq) (*model.ProductCategory, error) {
	return dao.ProductCategory.Where(dao.ProductCategory.ID.Eq(params.ID)).First()
}

// Status ...
func (s *CateService) Status(params product.StatusReq) error {
	if _, err := dao.ProductCategory.Where(dao.ProductCategory.ID.Eq(params.ID)).Update(dao.ProductCategory.Status, params.Status); err != nil {
		return err
	}
	return nil
}

// List ...
func (s *CateService) List() ([]*CateNode, error) {
	categories, err := dao.ProductCategory.Find()
	if err != nil {
		return nil, err
	}
	return SortCategoriesTree(categories, 0, 0), nil
}
