package admin

import "github.com/MQEnergy/mqshop/internal/app/model"

// LoginReq ...
type LoginReq struct {
	Account  string `form:"account" json:"account" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type WxLoginReq struct {
	Code string `form:"code" json:"code" validate:"required"`
}

// IndexReq ...
type IndexReq struct {
	Page    int    `form:"page" json:"page" validate:"required"`
	Limit   int    `form:"limit" json:"limit" validate:"required"`
	Keyword string `form:"keyword" json:"keyword"`
}

type CAdmin struct {
	model.CAdmin
	RoleList []*model.CRole `gorm:"-" json:"role_list"`
	IsSuper  int            `gorm:"-" json:"is_super"`
}

// CreateReq ...
type CreateReq struct {
	Account    string `json:"account" validate:"required"`    // 账号
	Password   string `json:"password"`                       // 密码
	Phone      string `form:"phone" json:"phone"`             // 手机号
	Avatar     string `form:"avatar" json:"avatar"`           // 头像
	Salt       string `form:"salt" json:"salt"`               // 密码加盐
	RealName   string `form:"real_name" json:"real_name"`     // 真实姓名
	RegisterIp string `form:"register_ip" json:"register_ip"` // 注册ip
	RoleIds    string `form:"role_ids" json:"role_ids"`       // 角色IDs
	Status     uint8  `form:"status" json:"status"`           // 状态 1：正常 2：禁用
}

// UpdateReq ...
type UpdateReq struct {
	Id       int64  `form:"id" json:"id" validate:"required"`
	Account  string `form:"account" json:"account" validate:"min=1,max=20"` // 账号
	Phone    string `form:"phone" json:"phone"`                             // 手机号
	Avatar   string `form:"avatar" json:"avatar"`                           // 头像
	RealName string `form:"real_name" json:"real_name"`                     // 真实姓名
	RoleIds  string `form:"role_ids" json:"role_ids" validate:"required"`   // 角色IDs
	Status   uint8  `form:"status" json:"status"`                           // 状态 1：正常 2：禁用
}

// DeleteReq ...
type DeleteReq struct {
	Id int64 `form:"id" json:"id" validate:"required"` //
}

// InfoReq ...
type InfoReq struct {
	Id int64 `query:"id" form:"id" json:"id" validate:"required"` //
}

// RoleDistributionReq 分配角色请求参数
type RoleDistributionReq struct {
	Id      int64  `form:"id" json:"id" validate:"required"` // 管理员ID
	RoleIds string `form:"role_ids" json:"role_ids" validate:"required"`
}

// ChangePassReq 修改密码请求参数
type ChangePassReq struct {
	OldPass    string `form:"old_pass" json:"old_pass" validate:"required"`
	NewPass    string `form:"new_pass" json:"new_pass" validate:"required"`
	RepeatPass string `form:"repeat_pass" json:"repeat_pass" validate:"required,eqfield=NewPass,min=6,max=20"`
}

// ChangePasswordReq 修改密码请求参数
type ChangePasswordReq struct {
	Uuid       string `form:"uuid" json:"uuid"`
	NewPass    string `form:"new_pass" json:"new_pass" validate:"required"`
	RepeatPass string `form:"repeat_pass" json:"repeat_pass" validate:"required,eqfield=NewPass,min=6,max=20"`
}

// ForgetPassReq 忘记密码请求参数
type ForgetPassReq struct {
	Type       int    `form:"type" json:"type" validate:"required"`       // 1: 验证用户名 2：验证验证码 3：设置密码并提交
	Account    string `form:"account" json:"account" validate:"required"` // 账号名称
	Phone      string `form:"phone" json:"phone"`                         // 号码
	Code       string `form:"code" json:"code"`                           // 验证码
	NewPass    string `form:"new_pass" json:"new_pass"`                   // 新密码
	RepeatPass string `form:"repeat_pass" json:"repeat_pass"`             // 重复密码
}

// SmsSendReq 发送验证码请求参数
type SmsSendReq struct {
	Phone string `form:"phone" json:"phone" validate:"required"` // 手机号
}

// WxQrRegisterReq ...
type WxQrRegisterReq struct {
	Code     string `form:"code" json:"code" validate:"required"`
	State    int    `fomr:"state" json:"state" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}
