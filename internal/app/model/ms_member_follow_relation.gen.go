// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMemberFollowRelation = "ms_member_follow_relation"

// MemberFollowRelation 用户关注关联表
type MemberFollowRelation struct {
	ID             int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	FollowMemberID int64  `gorm:"column:follow_member_id;type:bigint unsigned;not null;comment:被关注者ID" json:"follow_member_id"` // 被关注者ID
	MemberID       int64  `gorm:"column:member_id;type:bigint unsigned;not null;comment:关注者ID" json:"member_id"`                // 关注者ID
	CreatedAt      *int64 `gorm:"column:created_at;type:bigint unsigned" json:"created_at"`
	UpdatedAt      *int64 `gorm:"column:updated_at;type:bigint unsigned" json:"updated_at"`
}

// TableName MemberFollowRelation's table name
func (*MemberFollowRelation) TableName() string {
	return TableNameMemberFollowRelation
}
