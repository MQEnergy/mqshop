// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameBannerCate = "ms_banner_cate"

// BannerCate mapped from table <ms_banner_cate>
type BannerCate struct {
	ID        int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name      string `gorm:"column:name;type:varchar(255);not null;comment:名称" json:"name"`                       // 名称
	Desc      string `gorm:"column:desc;type:varchar(255);not null;comment:描述" json:"desc"`                       // 描述
	IsHome    int8   `gorm:"column:is_home;type:tinyint unsigned;not null;comment:是否放首页 1：放 0：不放" json:"is_home"` // 是否放首页 1：放 0：不放
	Status    int8   `gorm:"column:status;type:tinyint unsigned;not null;comment:状态 1：正常 0：下架" json:"status"`     // 状态 1：正常 0：下架
	CreatedAt *int64 `gorm:"column:created_at;type:bigint unsigned" json:"created_at"`
	UpdatedAt *int64 `gorm:"column:updated_at;type:bigint unsigned" json:"updated_at"`
}

// TableName BannerCate's table name
func (*BannerCate) TableName() string {
	return TableNameBannerCate
}
