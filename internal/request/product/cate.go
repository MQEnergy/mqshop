package product

import "github.com/MQEnergy/mqshop/internal/app/model"

type CateItem struct {
	model.ProductCategory
}

type CateCreateReq struct {
	CateName  string `form:"cate_name" json:"cate_name" validate:"required"`
	CateDesc  string `form:"cate_desc" json:"cate_desc"`
	CateUrl   string `form:"cate_url" json:"cate_url"`
	SortOrder int    `form:"sort_order" json:"sort_order"`
	IsHot     int8   `form:"is_hot" json:"is_hot"`
	IsIndex   int8   `form:"is_index" json:"is_index"`
	Status    int8   `form:"status" json:"status"`
	ParentID  int64  `form:"parent_id" json:"parent_id"`
}

type CateUpdateReq struct {
	ID        int64  `form:"id" json:"id" validate:"required"`
	CateName  string `form:"cate_name" json:"cate_name" validate:"required"`
	CateDesc  string `form:"cate_desc" json:"cate_desc"`
	CateUrl   string `form:"cate_url" json:"cate_url"`
	SortOrder int    `form:"sort_order" json:"sort_order"`
	IsHot     int8   `form:"is_hot" json:"is_hot"`
	IsIndex   int8   `form:"is_index" json:"is_index"`
	Status    int8   `form:"status" json:"status"`
	ParentID  int64  `form:"parent_id" json:"parent_id"`
}
