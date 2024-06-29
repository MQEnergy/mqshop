package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend"
	"github.com/MQEnergy/mqshop/internal/request/product"
	"github.com/MQEnergy/mqshop/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type BrandController struct {
	controller.Controller
}

var Brand = &BrandController{}

// Index ...
func (s *BrandController) Index(ctx *fiber.Ctx) error {
	var params product.IndexReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	brandList, err := backend.Brand.Index(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", brandList)
}

// View ...
func (s *BrandController) View(ctx *fiber.Ctx) error {
	var params product.ViewReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	brand, err := backend.Brand.View(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", brand)
}

// Status ...
func (s *BrandController) Status(ctx *fiber.Ctx) error {
	var params product.StatusReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.Brand.Status(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "操作成功", nil)
}

// Delete ...
func (s *BrandController) Delete(ctx *fiber.Ctx) error {
	var params product.MultiDeleteReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.Brand.Delete(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "操作成功", nil)
}

// Create ...
func (s *BrandController) Create(ctx *fiber.Ctx) error {
	var params product.BrandCreateReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.Brand.Create(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "操作成功", nil)
}

// Update ...
func (s *BrandController) Update(ctx *fiber.Ctx) error {
	var params product.BrandUpdateReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.Brand.Update(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "操作成功", nil)
}

// List ...
func (s *BrandController) List(ctx *fiber.Ctx) error {
	brandList, err := backend.Brand.List()
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", brandList)
}
