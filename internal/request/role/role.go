package role

// IndexReq ...
type IndexReq struct {
	Page    int    `query:"page" form:"page" json:"page" validate:"required"`
	Limit   int    `query:"limit" form:"limit" json:"limit" validate:"required"`
	Keyword string `query:"page" form:"keyword" json:"keyword"` // 搜索关键词
}

// CreateOrUpdateReq 创建和编辑请求参数
type CreateOrUpdateReq struct {
	Id     int    `form:"id" json:"id"`
	Name   string `form:"name" json:"name" validate:"required"` // 角色名称
	Desc   string `form:"desc" json:"desc"`                     // 角色描述
	Status int8   `form:"status" json:"status"`                 // 状态：1正常(默认) 2停用
}

// DeleteReq ...
type DeleteReq struct {
	Id int `form:"id" json:"id" validate:"required"` //
}

// ViewReq ...
type ViewReq struct {
	Id int `query:"id" form:"id" json:"id" validate:"required"` //
}

// InviteReq ...
type InviteReq struct {
	RoleId int `form:"role_id" json:"role_id" validate:"required"` //
}
