// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMemberGoodsFavorite = "ms_member_goods_favorite"

// MemberGoodsFavorite 用户收藏商品表
type MemberGoodsFavorite struct {
	ID        int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	MemberID  int64  `gorm:"column:member_id;type:bigint unsigned;not null;comment:用户ID" json:"member_id"` // 用户ID
	GoodsID   int64  `gorm:"column:goods_id;type:bigint unsigned;not null;comment:商品ID" json:"goods_id"`   // 商品ID
	CreatedAt *int64 `gorm:"column:created_at;type:bigint unsigned" json:"created_at"`
}

// TableName MemberGoodsFavorite's table name
func (*MemberGoodsFavorite) TableName() string {
	return TableNameMemberGoodsFavorite
}
