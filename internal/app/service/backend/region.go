package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/request/region"
)

type RegionService struct {
	service.Service
}

var Region = &RegionService{}

// GetProvinceList ...
func (s *RegionService) GetProvinceList() ([]*model.CProvinces, error) {
	return dao.CProvinces.Find()
}

// GetCityListByProvince ...
func (s *RegionService) GetCityListByProvince(params region.CityReq) ([]*model.CCities, error) {
	return dao.CCities.Where(dao.CCities.ProvinceCode.Eq(params.ProvinceCode)).Find()
}

// GetAreaListByCity ...
func (s *RegionService) GetAreaListByCity(params region.AreaReq) ([]*model.CAreas, error) {
	return dao.CAreas.Where(dao.CAreas.CityCode.Eq(params.CityCode)).Find()
}

// GetStreetListByArea ...
func (s *RegionService) GetStreetListByArea(params region.StreetReq) ([]*model.CStreets, error) {
	return dao.CStreets.Where(dao.CStreets.AreaCode.Eq(params.AreaCode)).Find()
}

// GetVillageListByStreet ...
func (s *RegionService) GetVillageListByStreet(params region.VillageReq) ([]*model.CVillages, error) {
	return dao.CVillages.Where(dao.CVillages.StreetCode.Eq(params.StreetCode)).Find()
}
