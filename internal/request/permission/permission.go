package permission

type ViewReq struct {
	RoleId int `query:"role_id" form:"role_id" json:"role_id" validate:"required"`
}

type MenuItem struct {
	Id       uint64 `form:"id" json:"id" validate:"required"`
	BUrl     string `form:"b_url" json:"b_url" validate:"required"`
	ParentId uint64 `form:"parent_id" json:"parent_id" validate:"required"`
}

type CreateReq struct {
	RoleId  int    `form:"role_id" json:"role_id" validate:"required"`
	MenuIds string `form:"menu_ids" json:"menu_ids" validate:"required"`
}
