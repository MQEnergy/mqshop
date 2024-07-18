package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend"
	"github.com/MQEnergy/mqshop/internal/request/region"
	"github.com/MQEnergy/mqshop/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type RegionController struct {
	controller.Controller
}

var Region = &RegionController{}

// Province ...
func (s *RegionController) Province(ctx *fiber.Ctx) error {
	provinceList, err := backend.Region.GetProvinceList()
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", provinceList)
}

// City ...
func (s *RegionController) City(ctx *fiber.Ctx) error {
	var params region.CityReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	provinceList, err := backend.Region.GetCityListByProvince(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", provinceList)
}

// Area ...
func (s *RegionController) Area(ctx *fiber.Ctx) error {
	var params region.AreaReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	provinceList, err := backend.Region.GetAreaListByCity(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", provinceList)
}

// Street ...
func (s *RegionController) Street(ctx *fiber.Ctx) error {
	var params region.StreetReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	provinceList, err := backend.Region.GetStreetListByArea(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", provinceList)
}

// Village ...
func (s *RegionController) Village(ctx *fiber.Ctx) error {
	var params region.VillageReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	provinceList, err := backend.Region.GetVillageListByStreet(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", provinceList)
}
