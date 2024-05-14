package rbac

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/MQEnergy/mqshop/internal/app/service/sredis"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
	"github.com/MQEnergy/mqshop/pkg/config"
	"github.com/MQEnergy/mqshop/pkg/wecom"

	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/request/admin"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/helper"
	"github.com/MQEnergy/mqshop/pkg/jwtauth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	service.Service
}

var Auth = &AuthService{}

// Login
// @Description: 登录
// @receiver s
// @param reqParams
// @return interface{}
// @return error
// @author cx
func (s *AuthService) Login(ctx *fiber.Ctx, reqParams admin.LoginReq) (fiber.Map, error) {
	var (
		isSuper = 0 // 是否超级管理员 1：是 0：不是
		u       = dao.CAdmin
	)
	adminInfo, err := u.Where(u.Account.Eq(reqParams.Account)).First()
	if err != nil {
		return nil, errors.New("账号或密码不正确")
	}
	if adminInfo.Status != 1 {
		return nil, errors.New("账号不存在或被限制登录")
	}
	if adminInfo.RoleIds == "" {
		return nil, errors.New("账号未分配角色，无法登录")
	}
	if adminInfo.Password != helper.GeneratePasswordHash(reqParams.Password, adminInfo.Salt) {
		return nil, errors.New("账号或密码不正确")
	}
	token, err := jwtauth.New(vars.Config).WithClaims(jwt.MapClaims{
		"id":       adminInfo.ID,
		"uuid":     adminInfo.UUID,
		"role_ids": adminInfo.RoleIds,
	}).GenerateToken()
	if err != nil {
		return nil, errors.New("登录失败")
	}
	if err := vars.Redis.Set(
		context.Background(),
		sredis.BuildAuthRedisKey(adminInfo.UUID),
		token,
		vars.Config.GetDuration("jwt.expire")*time.Second,
	).Err(); err != nil {
		return nil, errors.New("用户信息持久化失败")
	}
	if helper.InAnySlice[string](strings.Split(adminInfo.RoleIds, ","), vars.Config.GetString("server.superRoleId")) {
		isSuper = 1
	}
	// 更新login_ip login_time
	adminInfo.LoginIP = ctx.IP()
	adminInfo.LoginTime = time.Now().Unix()
	dao.CAdmin.Save(adminInfo)

	return fiber.Map{
		"token": token,
		"info": fiber.Map{
			"id":       adminInfo.ID,
			"uuid":     adminInfo.UUID,
			"account":  adminInfo.Account,
			"avatar":   adminInfo.Avatar,
			"role_ids": adminInfo.RoleIds,
			"is_super": isSuper,
		},
	}, nil
}

// Logout
// @Description: 退出登录
// @receiver s
// @param uuid
// @return error
// @author cx
func (s *AuthService) Logout(uuid string) error {
	_, err := vars.Redis.Del(context.Background(), sredis.BuildAuthRedisKey(uuid)).Result()
	if err != nil {
		return errors.New("会话过期，请重新登录")
	}
	return nil
}

// ChangePass
// @Description: 修改密码
// @receiver s
// @param requestParams
// @param info
// @return error
// @author cx
func (s *AuthService) ChangePass(requestParams admin.ChangePassReq, uuid string) error {
	u := dao.CAdmin
	adminInfo, err := u.Where(u.UUID.Eq(uuid)).First()
	if err != nil {
		return errors.New("用户信息不存在")
	}
	if adminInfo.Password != helper.GeneratePasswordHash(requestParams.OldPass, adminInfo.Salt) {
		return errors.New("原密码不正确")
	}
	salt := helper.GenerateUuid(32)
	if _, err := u.Where(u.UUID.Eq(uuid)).Updates(model.CAdmin{
		Password: helper.GeneratePasswordHash(requestParams.NewPass, salt),
		Salt:     salt,
	}); err != nil {
		return errors.New("修改密码失败")
	}
	// 删除redis
	s.Logout(uuid)
	return nil
}

// ForgetPass
// @Description: 忘记密码
// @receiver s
// @param params
// @return *model.CAdmin
// @return error
// @author cx
func (s *AuthService) ForgetPass(params admin.ForgetPassReq) (*model.CAdmin, error) {
	return nil, nil
}

// GetQrCodeUrl
// @Description: 获取二维码图片地址
// @receiver s
// @param config
// @return string
// @return error
// @author cx
func (s *AuthService) GetQrCodeUrl(state string, config *config.Config) (string, error) {
	c, err := wecom.New(
		wecom.WithCorpID(config.GetString("wecom.corpId")),
		wecom.WithAgentID(config.GetInt("wecom.agentId")),
		wecom.WithSecret(config.GetString("wecom.corpSecret")),
		wecom.WithCallbackURL(config.GetString("wecom.messageCallBack")),
		wecom.WithHttpDebug(true),
		wecom.WithOAuth(work.OAuth{Callback: config.GetString("wecom.oauthCallback"), Scopes: []string{"snsapi_base"}}),
		wecom.WithCache(kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{config.GetString("redis.host") + ":" + config.GetString("redis.port")},
			Password: config.GetString("redis.password"),
			DB:       config.GetInt("redis.dbNum"),
		})),
	)
	if err != nil {
		return "", err
	}
	return c.GetQrConnectURL(state)
}

// WxLogin
// @Description: 企业微信登录
// @receiver s
// @param ctx
// @param config
// @param params
// @return fiber.Map
// @return error
// @author cx
func (s *AuthService) WxLogin(ctx *fiber.Ctx, config *config.Config, params admin.WxLoginReq) (fiber.Map, error) {
	isSuper := 0 // 是否超级管理员 1：是 0：不是

	c, err := wecom.New(
		wecom.WithCorpID(config.GetString("wecom.corpId")),
		wecom.WithAgentID(config.GetInt("wecom.agentId")),
		wecom.WithSecret(config.GetString("wecom.corpSecret")),
		wecom.WithCallbackURL(config.GetString("wecom.messageCallBack")),
		wecom.WithHttpDebug(true),
		wecom.WithOAuth(work.OAuth{Callback: config.GetString("wecom.oauthCallback"), Scopes: []string{"snsapi_base"}}),
		wecom.WithCache(kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{config.GetString("redis.host") + ":" + config.GetString("redis.port")},
			Password: config.GetString("redis.password"),
			DB:       config.GetInt("redis.dbNum"),
		})),
	)
	if err != nil {
		return nil, err
	}
	userInfo, err := c.GetUserInfo(params.Code)
	if err != nil {
		return nil, err
	}
	// 登录
	u := dao.CAdmin
	adminInfo, err := u.Where(u.Account.Eq(userInfo.UserID)).First()
	if err != nil {
		return nil, errors.New("账号不存在或被限制")
	}
	if adminInfo.Status != 1 {
		return nil, errors.New("账号不存在或被限制登录")
	}
	token, err := jwtauth.New(vars.Config).WithClaims(jwt.MapClaims{
		"id":       adminInfo.ID,
		"uuid":     adminInfo.UUID,
		"role_ids": adminInfo.RoleIds,
	}).GenerateToken()
	if err != nil {
		return nil, errors.New("登录失败")
	}
	if err := vars.Redis.Set(
		context.Background(),
		sredis.BuildAuthRedisKey(adminInfo.UUID),
		token,
		vars.Config.GetDuration("jwt.expire")*time.Second,
	).Err(); err != nil {
		return nil, errors.New("用户信息持久化失败")
	}
	if helper.InAnySlice[string](strings.Split(adminInfo.RoleIds, ","), vars.Config.GetString("server.superRoleId")) {
		isSuper = vars.Config.GetInt("server.superRoleId")
	}
	// 更新login_ip login_time
	adminInfo.LoginIP = ctx.IP()
	adminInfo.LoginTime = time.Now().Unix()
	dao.CAdmin.Save(adminInfo)

	return fiber.Map{
		"token": token,
		"info": fiber.Map{
			"id":       adminInfo.ID,
			"uuid":     adminInfo.UUID,
			"account":  adminInfo.Account,
			"avatar":   adminInfo.Avatar,
			"role_ids": adminInfo.RoleIds,
			"is_super": isSuper,
		},
	}, nil
}

// WxQrRegister
// @Description: 企业微信授权注册
// @receiver s
// @param ctx
// @param config
// @param params
// @return fiber.Map
// @return error
// @author cx
func (s *AuthService) WxQrRegister(ctx *fiber.Ctx, config *config.Config, params admin.WxQrRegisterReq) (fiber.Map, error) {
	isSuper := 0 // 是否超级管理员 1：是 0：不是

	c, err := wecom.New(
		wecom.WithCorpID(config.GetString("wecom.corpId")),
		wecom.WithAgentID(config.GetInt("wecom.agentId")),
		wecom.WithSecret(config.GetString("wecom.corpSecret")),
		wecom.WithCallbackURL(config.GetString("wecom.messageCallBack")),
		wecom.WithHttpDebug(true),
		wecom.WithOAuth(work.OAuth{Callback: config.GetString("wecom.oauthCallback"), Scopes: []string{"snsapi_privateinfo"}}),
		wecom.WithCache(kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{config.GetString("redis.host") + ":" + config.GetString("redis.port")},
			Password: config.GetString("redis.password"),
			DB:       config.GetInt("redis.dbNum"),
		})),
	)
	if err != nil {
		return nil, err
	}
	userInfo, err := c.GetUserInfo(params.Code)
	if err != nil {
		slog.Info("GetUserInfo", err.Error())
		return nil, errors.New("当前账号已经注册")
	}
	userByID, err := c.GetUserByID(userInfo.UserID)
	if err != nil {
		slog.Info("userByID", err.Error())
		return nil, errors.New("当前账号已经注册")
	}
	// 登录
	u := dao.CAdmin
	if _, err := u.Where(u.Account.Eq(userByID.UserID)).First(); err == nil {
		return nil, errors.New("当前账号已经注册")
	}
	// 注册
	salt := helper.GenerateUuid(32)
	uuid := helper.GenerateBaseSnowId(0, nil)
	password := helper.GeneratePasswordHash(params.Password, salt)
	adminInfo := &model.CAdmin{
		UUID:         uuid,
		Account:      userByID.UserID,
		Password:     password,
		Phone:        userByID.Mobile,
		Avatar:       userByID.Avatar,
		Salt:         salt,
		RealName:     userByID.Name,
		RegisterTime: time.Now().Unix(),
		RegisterIP:   ctx.IP(),
		LoginTime:    0,
		LoginIP:      "",
		RoleIds:      strconv.Itoa(params.State),
		Status:       1,
	}
	if err := u.Save(adminInfo); err != nil {
		return nil, errors.New("注册失败 err: " + err.Error())
	}
	// 生成token
	token, err := jwtauth.New(vars.Config).WithClaims(jwt.MapClaims{
		"id":       adminInfo.ID,
		"uuid":     adminInfo.UUID,
		"role_ids": adminInfo.RoleIds,
	}).GenerateToken()
	if err != nil {
		return nil, errors.New("登录失败")
	}
	if err := vars.Redis.Set(
		context.Background(),
		sredis.BuildAuthRedisKey(adminInfo.UUID),
		token,
		vars.Config.GetDuration("jwt.expire")*time.Second,
	).Err(); err != nil {
		return nil, errors.New("用户信息持久化失败")
	}
	if helper.InAnySlice[string](strings.Split(adminInfo.RoleIds, ","), vars.Config.GetString("server.superRoleId")) {
		isSuper = vars.Config.GetInt("server.superRoleId")
	}
	// 更新login_ip login_time
	adminInfo.LoginIP = ctx.IP()
	adminInfo.LoginTime = time.Now().Unix()
	dao.CAdmin.Save(adminInfo)

	return fiber.Map{
		"token": token,
		"info": fiber.Map{
			"id":       adminInfo.ID,
			"uuid":     adminInfo.UUID,
			"account":  adminInfo.Account,
			"avatar":   adminInfo.Avatar,
			"role_ids": adminInfo.RoleIds,
			"is_super": isSuper,
		},
	}, nil
}
