package resource

type IndexReq struct {
	Page int `query:"page" form:"page" json:"page"`
}

type ListReq struct {
	Id int64 `query:"id" form:"id" json:"id"`
}

type CreateOrUpdateReq struct {
	Id         int64  `form:"id" json:"id"`
	Name       string `form:"name" json:"name" validate:"required"`           // 菜单名称
	Alias_     string `form:"alias" json:"alias" validate:"required"`         // 菜单别名（前端加必传）
	Desc       string `form:"desc" json:"desc"`                               // 菜单描述
	FUrl       string `form:"f_url" json:"f_url" validate:"required"`         // 菜单前端URL
	BUrl       string `form:"b_url" json:"b_url" validate:"required"`         // 后端URL
	Icon       string `form:"icon" json:"icon"`                               // 菜单icon
	ParentId   uint64 `form:"parent_id" json:"parent_id"`                     // 父级ID
	Path       string `form:"path" json:"path"`                               // ID路径 1-2-3...
	MenuType   int8   `form:"menu_type" json:"menu_type" validate:"required"` // 菜单类型 1：模块 2：操作
	Status     uint8  `form:"status" json:"status"`                           // 状态：1正常(默认) 0停用
	SortOrder  uint16 `form:"sort_order" json:"sort_order"`                   // 排序
	ParentPath string `form:"parent_path" json:"parent_path"`                 // 父菜单的path路径
}

type DeleteReq struct {
	Id int64 `form:"id" json:"id" validate:"required"`
}

// Item ...
type Item struct {
	ID            int64   `form:"id" json:"id" validate:"required"`
	ParentID      int64   `form:"parent_id" json:"parent_id"`
	Label         string  `form:"label" json:"label"`
	Name          string  `form:"name" json:"name"`
	Alias_        string  `form:"alias" json:"alias"`
	Level         int     `form:"level" json:"level"`
	Path          string  `form:"path" json:"path"`
	MenuType      int8    `form:"menu_type" json:"menu_type"`
	Actions       []*Item `form:"actions" json:"actions"`
	CheckData     []int   `form:"checkData" json:"checkData"`
	Indeterminate bool    `form:"indeterminate" json:"indeterminate"`
	CheckedAll    bool    `form:"checkedAll" json:"checkedAll"`
}

// ViewReq ...
type ViewReq struct {
	Id int64 `form:"id" json:"id" validate:"required"`
}
