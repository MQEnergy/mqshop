package product

type BrandCreateReq struct {
	BrandName string `form:"brand_name" json:"brand_name" validate:"required"`
	LogoUrl   string `form:"logo_url" json:"logo_url" validate:"required"`
	Desc      string `form:"desc" json:"desc"`
	SortOrder int    `form:"sort_order" json:"sort_order" validate:"required"`
	IsHot     int8   `form:"is_hot" json:"is_hot"`
	Status    int8   `form:"status" json:"status"`
}

type BrandUpdateReq struct {
	ID        int64  `form:"id" json:"id" validate:"required"`
	BrandName string `form:"brand_name" json:"brand_name" validate:"required"`
	LogoUrl   string `form:"logo_url" json:"logo_url" validate:"required"`
	Desc      string `form:"desc" json:"desc"`
	SortOrder int    `form:"sort_order" json:"sort_order"`
	IsHot     int8   `form:"is_hot" json:"is_hot"`
	Status    int8   `form:"status" json:"status"`
}
