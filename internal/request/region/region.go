package region

type CityReq struct {
	ProvinceCode string `query:"province_code" form:"province_code" json:"province_code" validate:"required"`
}

type AreaReq struct {
	CityCode string `query:"city_code" form:"city_code" json:"city_code" validate:"required"`
}

type StreetReq struct {
	AreaCode string `query:"area_code" form:"area_code" json:"area_code" validate:"required"`
}

type VillageReq struct {
	StreetCode string `query:"street_code" form:"street_code" json:"street_code" validate:"required"`
}
