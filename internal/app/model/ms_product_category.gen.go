// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProductCategory = "ms_product_category"

// ProductCategory 产品分类表
type ProductCategory struct {
	ID           int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	CateName     string `gorm:"column:cate_name;type:varchar(64);not null;comment:分类名称" json:"cate_name"`                  // 分类名称
	BannerCateID int64  `gorm:"column:banner_cate_id;type:bigint unsigned;not null;comment:轮播图分类ID" json:"banner_cate_id"` // 轮播图分类ID
	CateURL      string `gorm:"column:cate_url;type:varchar(128);not null;comment:分类图标" json:"cate_url"`                   // 分类图标
	CateDesc     string `gorm:"column:cate_desc;type:varchar(128);not null;comment:分类描述" json:"cate_desc"`                 // 分类描述
	SortOrder    int    `gorm:"column:sort_order;type:int unsigned;not null;default:50;comment:排序" json:"sort_order"`      // 排序
	ParentID     int64  `gorm:"column:parent_id;type:bigint unsigned;not null;comment:父级ID" json:"parent_id"`              // 父级ID
	Path         string `gorm:"column:path;type:varchar(255);not null;comment:ID路径 1-2-3-4..." json:"path"`                // ID路径 1-2-3-4...
	IsHot        int8   `gorm:"column:is_hot;type:tinyint unsigned;not null;comment:是否热门 1：热门 0：非热门" json:"is_hot"`        // 是否热门 1：热门 0：非热门
	IsIndex      int8   `gorm:"column:is_index;type:tinyint unsigned;not null;comment:是否首页 1：首页 0：非首页" json:"is_index"`    // 是否首页 1：首页 0：非首页
	Status       int8   `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:状态 1：正常 0：下架" json:"status"` // 状态 1：正常 0：下架
	CreatedAt    *int64 `gorm:"column:created_at;type:bigint unsigned" json:"created_at"`
	UpdatedAt    *int64 `gorm:"column:updated_at;type:bigint unsigned" json:"updated_at"`
}

// TableName ProductCategory's table name
func (*ProductCategory) TableName() string {
	return TableNameProductCategory
}