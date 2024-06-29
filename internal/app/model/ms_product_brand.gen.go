// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProductBrand = "ms_product_brand"

// ProductBrand 商品品牌表
type ProductBrand struct {
	ID        int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	BrandName string `gorm:"column:brand_name;type:varchar(32);not null;comment:品牌名称" json:"brand_name"`                // 品牌名称
	LogoURL   string `gorm:"column:logo_url;type:varchar(255);not null;comment:品牌logo" json:"logo_url"`                 // 品牌logo
	Desc      string `gorm:"column:desc;type:varchar(255);not null;comment:品牌描述" json:"desc"`                           // 品牌描述
	SortOrder int    `gorm:"column:sort_order;type:int unsigned;not null;default:50;comment:排序" json:"sort_order"`      // 排序
	IsHot     int8   `gorm:"column:is_hot;type:tinyint unsigned;not null;comment:是否热门 1：热门 0：非热门" json:"is_hot"`        // 是否热门 1：热门 0：非热门
	Status    int8   `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:状态 1：正常 0：下架" json:"status"` // 状态 1：正常 0：下架
	CreatedAt int64  `gorm:"column:created_at;type:bigint unsigned;not null" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;type:bigint unsigned;not null" json:"updated_at"`
}

// TableName ProductBrand's table name
func (*ProductBrand) TableName() string {
	return TableNameProductBrand
}
