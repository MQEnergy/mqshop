package product

// IndexReq ...
type IndexReq struct {
	Page   int    `form:"page" json:"page" validate:"required"`
	Limit  int    `form:"limit" json:"limit" validate:"required"`
	Search string `form:"search" json:"search"`
}

type ViewReq struct {
	ID int64 `form:"id" json:"id" validate:"required"`
}

type DeleteReq struct {
	ID int64 `form:"id" json:"id" validate:"required"`
}

// MultiDeleteReq ...
type MultiDeleteReq struct {
	IDs string `form:"ids" json:"ids" validate:"required"` // 1,2,3,4 ...
}

type StatusReq struct {
	ID     int64 `form:"id" json:"id" validate:"required"`
	Status int8  `form:"status" json:"status"`
}

type CreateReq struct {
	GoodsTitle    string  `form:"goods_title" json:"goods_title" validate:"required"`
	GoodsSubtitle string  `form:"goods_subtitle" json:"goods_subtitle" validate:"required"`
	CateID        int64   `form:"cate_id" json:"cate_id" validate:"required"`
	BrandID       int64   `form:"brand_id" json:"brand_id" validate:"required"`
	AttrCateID    int64   `form:"attr_cate_id" json:"attr_cate_id" validate:"required"`
	ThumbList     string  `form:"thumb_list" json:"thumb_list" validate:"required"`
	ImageList     string  `form:"image_list" json:"image_list" validate:"required"`
	OriginPrice   float64 `form:"origin_price" json:"origin_price" validate:"required"`
	PromotePrice  float64 `form:"promote_price" json:"promote_price" validate:"required"`
	FinalPrice    float64 `form:"final_price" json:"final_price" validate:"required"`
	TotalStock    int     `form:"total_stock" json:"total_stock" validate:"required"`
	IsHot         int8    `form:"is_hot" json:"is_hot" validate:"required"`
	IsNew         int8    `form:"is_new" json:"is_new" validate:"required"`
	IsRecommend   int8    `form:"is_recommend" json:"is_recommend" validate:"required"`
	SortOrder     int     `form:"sort_order" json:"sort_order" validate:"required"`
	Status        int8    `form:"status" json:"status" validate:"required"`
}

type UpdateReq struct {
	ID            int64   `form:"id" json:"id" validate:"required"`
	GoodsTitle    string  `form:"goods_title" json:"goods_title" validate:"required"`
	GoodsSubtitle string  `form:"goods_subtitle" json:"goods_subtitle" validate:"required"`
	CateID        int64   `form:"cate_id" json:"cate_id" validate:"required"`
	BrandID       int64   `form:"brand_id" json:"brand_id" validate:"required"`
	AttrCateID    int64   `form:"attr_cate_id" json:"attr_cate_id" validate:"required"`
	ThumbList     string  `form:"thumb_list" json:"thumb_list" validate:"required"`
	ImageList     string  `form:"image_list" json:"image_list" validate:"required"`
	OriginPrice   float64 `form:"origin_price" json:"origin_price" validate:"required"`
	PromotePrice  float64 `form:"promote_price" json:"promote_price" validate:"required"`
	FinalPrice    float64 `form:"final_price" json:"final_price" validate:"required"`
	TotalStock    int     `form:"total_stock" json:"total_stock" validate:"required"`
	IsHot         int8    `form:"is_hot" json:"is_hot" validate:"required"`
	IsNew         int8    `form:"is_new" json:"is_new" validate:"required"`
	IsRecommend   int8    `form:"is_recommend" json:"is_recommend" validate:"required"`
	SortOrder     int     `form:"sort_order" json:"sort_order" validate:"required"`
	Status        int8    `form:"status" json:"status" validate:"required"`
}

type AttrsReq struct {
	ID         int64 `form:"id" json:"id" validate:"required"`
	AttrCateID int64 `form:"attr_cate_id" json:"attr_cate_id" validate:"required"`
}
