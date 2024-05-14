// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProductGoodsTagRelation = "ms_product_goods_tag_relation"

// ProductGoodsTagRelation 商品标签关联表
type ProductGoodsTagRelation struct {
	ID        int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	GoodsID   int64  `gorm:"column:goods_id;type:bigint unsigned;not null;comment:商品ID" json:"goods_id"` // 商品ID
	TagID     int64  `gorm:"column:tag_id;type:bigint unsigned;not null;comment:标签ID" json:"tag_id"`     // 标签ID
	CreatedAt *int64 `gorm:"column:created_at;type:bigint unsigned" json:"created_at"`
}

// TableName ProductGoodsTagRelation's table name
func (*ProductGoodsTagRelation) TableName() string {
	return TableNameProductGoodsTagRelation
}