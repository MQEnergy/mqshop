package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend/rbac"
	"github.com/MQEnergy/mqshop/internal/request/admin"
	"github.com/MQEnergy/mqshop/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	controller.Controller
}

var Admin = &AdminController{}

// Index 获取列表
func (c *AdminController) Index(ctx *fiber.Ctx) error {
	var reqParams admin.IndexReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	list, err := rbac.Admin.Index(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", list)
}

// Create ...
func (c *AdminController) Create(ctx *fiber.Ctx) error {
	var reqParams admin.CreateReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	reqParams.RegisterIp = ctx.IP()
	if err := rbac.Admin.Create(reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}

// Update ...
func (c *AdminController) Update(ctx *fiber.Ctx) error {
	var reqParams admin.UpdateReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := rbac.Admin.Update(reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}

// Delete ...
func (c *AdminController) Delete(ctx *fiber.Ctx) error {
	var reqParams admin.DeleteReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := rbac.Admin.Delete(reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}

// View 获取登录用户信息
func (c *AdminController) View(ctx *fiber.Ctx) error {
	uuid := ctx.GetRespHeader("uuid")
	info, err := rbac.Admin.View(uuid)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", info)
}

// Info 根据用户ID获取用户信息
func (c *AdminController) Info(ctx *fiber.Ctx) error {
	var reqParams admin.InfoReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	info, err := rbac.Admin.Info(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", info)
}

// RoleDistribution 角色分配
func (c *AdminController) RoleDistribution(ctx *fiber.Ctx) error {
	var reqParams admin.RoleDistributionReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := rbac.Admin.RoleDistribution(reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}

// ChangePassword 修改密码
func (c *AdminController) ChangePassword(ctx *fiber.Ctx) error {
	var reqParams admin.ChangePasswordReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := rbac.Admin.ChangePass(reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}
