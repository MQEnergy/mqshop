package product

type AttrCate struct {
	ID        int64  `json:"id"`
	CateName  string `json:"cate_name"`
	Status    int8   `json:"status"`
	ParamsNum int64  `json:"params_num"` // 参数数量
	AttrNum   int64  `json:"attr_num"`   // 属性数量
	CreatedAt int64  `json:"created_at"`
}

type AttrCateCreateReq struct {
	CateName string `form:"cate_name" json:"cate_name" validate:"required"`
	Status   int8   `form:"status" json:"status"`
}

type AttrCateUpdateReq struct {
	ID       int64  `form:"id" json:"id" validate:"required"`
	CateName string `form:"cate_name" json:"cate_name" validate:"required"`
	Status   int8   `form:"status" json:"status"`
}

type AttrCateAttrParamListReq struct {
	ID       int64  `form:"id" json:"id" validate:"required"`
	AttrType int8   `form:"attr_type" json:"attr_type" validate:"required"`
	Page     int    `form:"page" json:"page" validate:"required"`
	Limit    int    `form:"limit" json:"limit" validate:"required"`
	Search   string `form:"search" json:"search"`
}

type AttrCateAttrParamCreateReq struct {
	CateID    int64  `form:"cate_id" json:"cate_id" validate:"required"`
	AttrType  int8   `form:"attr_type" json:"attr_type" validate:"required"`
	AttrName  string `form:"attr_name" json:"attr_name" validate:"required"`
	AttrValue string `form:"attr_value" json:"attr_value"`
	InputType int8   `form:"input_type" json:"input_type" validate:"required"`
	SortOrder int    `form:"sort_order" json:"sort_order" validate:"required"`
}

type AttrCateAttrParamUpdateReq struct {
	ID        int64  `form:"id" json:"id" validate:"required"`
	CateID    int64  `form:"cate_id" json:"cate_id" validate:"required"`
	AttrType  int8   `form:"attr_type" json:"attr_type" validate:"required"`
	AttrName  string `form:"attr_name" json:"attr_name" validate:"required"`
	AttrValue string `form:"attr_value" json:"attr_value"`
	InputType int8   `form:"input_type" json:"input_type" validate:"required"`
	SortOrder int    `form:"sort_order" json:"sort_order" validate:"required"`
}
