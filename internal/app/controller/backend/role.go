package backend

import (
	"strconv"

	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend/rbac"
	"github.com/MQEnergy/mqshop/internal/request/role"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type RoleController struct {
	*controller.Controller
}

var Role = &RoleController{}

// Index 获取列表
func (c *RoleController) Index(ctx *fiber.Ctx) error {
	var reqParams role.IndexReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	list, err := rbac.Role.Index(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", list)
}

// CreateOrUpdate ...
func (c *RoleController) CreateOrUpdate(ctx *fiber.Ctx) error {
	var reqParams role.CreateOrUpdateReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := rbac.Role.CreateOrUpdate(reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}

// Delete 状态禁止
func (c *RoleController) Delete(ctx *fiber.Ctx) error {
	var reqParams role.DeleteReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	if err := rbac.Role.Delete(reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}

// View ...
func (c *RoleController) View(ctx *fiber.Ctx) error {
	var reqParams role.ViewReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	info, err := rbac.Role.View(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", info)
}

// List ...
func (c *RoleController) List(ctx *fiber.Ctx) error {
	list := rbac.Role.List()
	return response.SuccessJSON(ctx, "", list)
}

// Invite ...
func (c *RoleController) Invite(ctx *fiber.Ctx) error {
	var reqParams role.InviteReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	codeUrl, err := rbac.Auth.GetQrCodeUrl(strconv.Itoa(reqParams.RoleId), vars.Config)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", codeUrl)
}
