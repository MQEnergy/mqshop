package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend"
	"github.com/MQEnergy/mqshop/internal/request/product"
	"github.com/MQEnergy/mqshop/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type AttrCateController struct {
	controller.Controller
}

var AttrCate = &AttrCateController{}

// Index ...
func (s *AttrCateController) Index(ctx *fiber.Ctx) error {
	var params product.IndexReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	result, err := backend.AttrCate.Index(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "获取成功", result)
}

// View ...
func (s *AttrCateController) View(ctx *fiber.Ctx) error {
	var params product.ViewReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	result, err := backend.AttrCate.View(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "获取成功", result)
}

// Create ...
func (s *AttrCateController) Create(ctx *fiber.Ctx) error {
	var params product.AttrCateCreateReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.AttrCate.Create(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "创建成功", "")
}

// Update ...
func (s *AttrCateController) Update(ctx *fiber.Ctx) error {
	var params product.AttrCateUpdateReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.AttrCate.Update(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "更新成功", "")
}

// Delete ...
func (s *AttrCateController) Delete(ctx *fiber.Ctx) error {
	var params product.MultiDeleteReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.AttrCate.Delete(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "删除成功", "")
}

// List ...
func (s *AttrCateController) List(ctx *fiber.Ctx) error {
	result, err := backend.AttrCate.List()
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "获取成功", result)
}

// AttrParamsList ...
func (s *AttrCateController) AttrParamsList(ctx *fiber.Ctx) error {
	var params product.AttrCateAttrParamListReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	result, err := backend.AttrCate.AttrParamsList(params)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "获取成功", result)
}

// AttrParamsCreate ...
func (s *AttrCateController) AttrParamsCreate(ctx *fiber.Ctx) error {
	var params product.AttrCateAttrParamCreateReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.AttrCate.AttrParamsCreate(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "创建成功", "")
}

// AttrParamsUpdate ...
func (s *AttrCateController) AttrParamsUpdate(ctx *fiber.Ctx) error {
	var params product.AttrCateAttrParamUpdateReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.AttrCate.AttrParamsUpdate(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "更新成功", "")
}

// AttrParamsDelete ...
func (s *AttrCateController) AttrParamsDelete(ctx *fiber.Ctx) error {
	var params product.MultiDeleteReq
	if err := s.Validate(ctx, &params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := backend.AttrCate.AttrParamsDelete(params); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "删除成功", "")
}
