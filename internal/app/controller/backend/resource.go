package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend/rbac"
	"github.com/MQEnergy/mqshop/internal/request/resource"
	"github.com/MQEnergy/mqshop/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type ResourceController struct {
	*controller.Controller
}

var Resource = &ResourceController{}

// Index 获取列表
func (c *ResourceController) Index(ctx *fiber.Ctx) error {
	var reqParams resource.IndexReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	list, err := rbac.Resource.Index(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", list)
}

func (c *ResourceController) ResourceList(ctx *fiber.Ctx) error {
	var reqParams resource.ListReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	list, err := rbac.Resource.MenuList(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", list)
}

// CreateOrUpdate 创建或更新
func (c *ResourceController) CreateOrUpdate(ctx *fiber.Ctx) error {
	var reqParams resource.CreateOrUpdateReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	err := rbac.Resource.CreateOrUpdate(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", nil)
}

// Delete ...
func (c *ResourceController) Delete(ctx *fiber.Ctx) error {
	var reqParams resource.DeleteReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	err := rbac.Resource.Delete(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}

// View ...
func (c *ResourceController) View(ctx *fiber.Ctx) error {
	var reqParams resource.ViewReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	info, err := rbac.Resource.View(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", info)
}
