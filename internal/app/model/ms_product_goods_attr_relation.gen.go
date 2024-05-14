// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProductGoodsAttrRelation = "ms_product_goods_attr_relation"

// ProductGoodsAttrRelation 商户属性关联表
type ProductGoodsAttrRelation struct {
	ID        int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	GoodsID   int64  `gorm:"column:goods_id;type:bigint unsigned;not null;comment:商品ID" json:"goods_id"`              // 商品ID
	AttrID    int64  `gorm:"column:attr_id;type:bigint unsigned;not null;comment:属性ID" json:"attr_id"`                // 属性ID
	AttrType  int8   `gorm:"column:attr_type;type:tinyint unsigned;not null;comment:属性类型 1：属性 2：参数" json:"attr_type"` // 属性类型 1：属性 2：参数
	AttrName  string `gorm:"column:attr_name;type:varchar(255);not null;comment:商品属性名称 如：颜色" json:"attr_name"`        // 商品属性名称 如：颜色
	AttrValue string `gorm:"column:attr_value;type:varchar(255);not null;comment:商品属性值 如：金色,银色,黄色" json:"attr_value"` // 商品属性值 如：金色,银色,黄色
	CreatedAt *int64 `gorm:"column:created_at;type:bigint unsigned" json:"created_at"`
	UpdatedAt *int64 `gorm:"column:updated_at;type:bigint unsigned" json:"updated_at"`
}

// TableName ProductGoodsAttrRelation's table name
func (*ProductGoodsAttrRelation) TableName() string {
	return TableNameProductGoodsAttrRelation
}