package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend"
	"github.com/MQEnergy/mqshop/internal/request/product"
	"github.com/MQEnergy/mqshop/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	controller.Controller
}

var Product = &ProductController{}

// Index ...
func (s *ProductController) Index(ctx *fiber.Ctx) error {
	var params product.IndexReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	pagination, err := backend.Product.Index(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", pagination)
}

// Create ...
func (s *ProductController) Create(ctx *fiber.Ctx) error {
	var params product.CreateReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.Product.Create(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "创建成功", nil)
}

// Update ...
func (s *ProductController) Update(ctx *fiber.Ctx) error {
	var params product.UpdateReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.Product.Update(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "更新成功", nil)
}

// View ...
func (s *ProductController) View(ctx *fiber.Ctx) error {
	var params product.ViewReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	productGoods, err := backend.Product.View(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", productGoods)
}

// Delete ...
func (s *ProductController) Delete(ctx *fiber.Ctx) error {
	var params product.DeleteReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	err := backend.Product.Delete(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "删除成功", nil)
}

// Status ...
func (s *ProductController) Status(ctx *fiber.Ctx) error {
	var params product.StatusReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	err := backend.Product.Status(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "更新成功", nil)
}
