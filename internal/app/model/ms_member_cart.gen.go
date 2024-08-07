// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMemberCart = "ms_member_cart"

// MemberCart 用户购物车
type MemberCart struct {
	ID        int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	MemberID  int64  `gorm:"column:member_id;type:bigint unsigned;not null;comment:用户ID" json:"member_id"`      // 用户ID
	GoodsID   int64  `gorm:"column:goods_id;type:bigint unsigned;not null;comment:商品ID" json:"goods_id"`        // 商品ID
	GoodsNum  int    `gorm:"column:goods_num;type:int unsigned;not null;default:1;comment:数量" json:"goods_num"` // 数量
	SkuNo     string `gorm:"column:sku_no;type:varchar(64);not null;comment:商品sku编码" json:"sku_no"`             // 商品sku编码
	CreatedAt *int64 `gorm:"column:created_at;type:bigint unsigned" json:"created_at"`
	UpdatedAt *int64 `gorm:"column:updated_at;type:bigint unsigned" json:"updated_at"`
}

// TableName MemberCart's table name
func (*MemberCart) TableName() string {
	return TableNameMemberCart
}
