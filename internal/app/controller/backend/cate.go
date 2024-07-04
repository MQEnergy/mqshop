package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend"
	"github.com/MQEnergy/mqshop/internal/request/product"
	"github.com/MQEnergy/mqshop/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type CateController struct {
	controller.Controller
}

var Cate = &CateController{}

// Index ...
func (s *CateController) Index(ctx *fiber.Ctx) error {
	var params product.IndexReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	productGoods, err := backend.Cate.Index(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", productGoods)
}

// Create ...
func (s *CateController) Create(ctx *fiber.Ctx) error {
	var params product.CateCreateReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.Cate.Create(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "创建成功", "")
}

// Update ...
func (s *CateController) Update(ctx *fiber.Ctx) error {
	var params product.CateUpdateReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.Cate.Update(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "更新成功", "")
}

// View ...
func (s *CateController) View(ctx *fiber.Ctx) error {
	var params product.ViewReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	productGoods, err := backend.Cate.View(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", productGoods)
}

// Delete ...
func (s *CateController) Delete(ctx *fiber.Ctx) error {
	var params product.MultiDeleteReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.Cate.Delete(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "删除成功", "")
}

// Status ...
func (s *CateController) Status(ctx *fiber.Ctx) error {
	var params product.StatusReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.Cate.Status(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "更新成功", "")
}

// List ...
func (s *CateController) List(ctx *fiber.Ctx) error {
	var params product.CateListReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	productGoods, err := backend.Cate.List(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", productGoods)
}
