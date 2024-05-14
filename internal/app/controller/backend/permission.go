package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend/rbac"
	"github.com/MQEnergy/mqshop/internal/request/permission"
	"github.com/MQEnergy/mqshop/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type PermissionController struct {
	controller.Controller
}

var Permission = &PermissionController{}

// View 查看角色权限
func (c *PermissionController) View(ctx *fiber.Ctx) error {
	var reqParams permission.ViewReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	list, menuIds, err := rbac.Permission.View(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", fiber.Map{"list": list, "menu_ids": menuIds})
}

// Create ...
func (c *PermissionController) Create(ctx *fiber.Ctx) error {
	var reqParams permission.CreateReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := rbac.Permission.Create(reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}

// Alias 获取登录后权限详情
func (c *PermissionController) Alias(ctx *fiber.Ctx) error {
	uuid := ctx.GetRespHeader("uuid")
	aliasList, err := rbac.Permission.Alias(uuid)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", aliasList)
}
