package backend

import (
	"github.com/MQEnergy/mqshop/internal/app/controller"
	"github.com/MQEnergy/mqshop/internal/app/service/backend/rbac"
	"github.com/MQEnergy/mqshop/internal/request/admin"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	controller.Controller
}

var Auth = &AuthController{}

// Login 用户登录
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var reqParams admin.LoginReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	info, err := rbac.Auth.Login(ctx, reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", info)
}

// Logout 退出登录
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	uuid := ctx.GetRespHeader("uuid")
	if err := rbac.Auth.Logout(uuid); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}

// ChangePass 修改密码
func (c *AuthController) ChangePass(ctx *fiber.Ctx) error {
	var reqParams admin.ChangePassReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	uuid := ctx.GetRespHeader("uuid")
	if err := rbac.Auth.ChangePass(reqParams, uuid); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", "")
}

// WxLogin 企业微信登录
func (c *AuthController) WxLogin(ctx *fiber.Ctx) error {
	var reqParams admin.WxLoginReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	info, err := rbac.Auth.WxLogin(ctx, vars.Config, reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", info)
}

// ForgetPass 忘记密码
func (c *AuthController) ForgetPass(ctx *fiber.Ctx) error {
	var reqParams admin.ForgetPassReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	adminInfo, err := rbac.Auth.ForgetPass(reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", adminInfo)
}

// WxQrRegister 验证码发送
func (c *AuthController) WxQrRegister(ctx *fiber.Ctx) error {
	var reqParams admin.WxQrRegisterReq
	if err := c.Validate(ctx, &reqParams); err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	info, err := rbac.Auth.WxQrRegister(ctx, vars.Config, reqParams)
	if err != nil {
		return response.BadRequestException(ctx, err.Error())
	}
	return response.SuccessJSON(ctx, "", info)
}
