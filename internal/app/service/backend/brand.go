package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/service"
)

type BrandService struct {
	service.Service
}

var Brand = &BrandService{}

// Index ...
func (s *BrandService) Index() error {
	// Todo implement ...
	return nil
}
