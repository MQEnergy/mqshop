// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTag = "ms_tag"

// Tag 标签表
type Tag struct {
	ID        int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	TagName   string `gorm:"column:tag_name;type:varchar(255);not null;comment:标签名称" json:"tag_name"` // 标签名称
	CreatedAt *int64 `gorm:"column:created_at;type:bigint unsigned" json:"created_at"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
