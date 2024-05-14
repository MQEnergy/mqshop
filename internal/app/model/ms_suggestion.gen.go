// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSuggestion = "ms_suggestion"

// Suggestion 问题反馈表
type Suggestion struct {
	ID        int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	MemberID  int64  `gorm:"column:member_id;type:bigint unsigned;not null;comment:用户ID" json:"member_id"`                      // 用户ID
	Content   string `gorm:"column:content;type:varchar(255);not null;comment:问题描述" json:"content"`                             // 问题描述
	Email     string `gorm:"column:email;type:varchar(32);not null;comment:email" json:"email"`                                 // email
	Status    int8   `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:状态 1：未处理 2：已处理 3：已删除" json:"status"` // 状态 1：未处理 2：已处理 3：已删除
	CreatedAt *int64 `gorm:"column:created_at;type:bigint unsigned" json:"created_at"`
}

// TableName Suggestion's table name
func (*Suggestion) TableName() string {
	return TableNameSuggestion
}