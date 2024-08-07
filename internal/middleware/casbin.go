package middleware

import (
	_ "embed"
	"strings"

	"github.com/MQEnergy/mqshop/configs"

	"gorm.io/gorm"

	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/helper"
	"github.com/MQEnergy/mqshop/pkg/response"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/fiber/v2"
)

// CasbinMiddleware casbin middleware
func CasbinMiddleware(db *gorm.DB, prefix, tableName string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if db == nil {
			return ctx.Next()
		}
		if strings.HasSuffix(ctx.Path(), "/backend/auth/login") ||
			strings.HasSuffix(ctx.Path(), "/backend/auth/logout") ||
			strings.HasPrefix(ctx.Path(), "/backend/auth/forget-pass") {
			return ctx.Next()
		}

		if tableName == "" {
			tableName = "casbin_rule"
		}
		adapter, _ := gormadapter.NewFilteredAdapterByDB(db, prefix, tableName)
		rc, _ := model.NewModelFromString(configs.RbacModelConf)
		e, _ := casbin.NewEnforcer(rc, adapter)
		e.AddFunction("ParamsObjMatch", ParamsObjMatchFunc)
		e.AddFunction("ParamsActMatch", ParamsActMatchFunc)
		_ = e.LoadPolicy()

		//	获取当前请求的url
		obj := ctx.Path()
		act := ctx.Method()
		roleIds := ctx.GetRespHeader("role_ids")
		if roleIds == "" {
			return response.UnauthorizedException(ctx, "该用户还未分配权限")
		}
		roleList := strings.Split(roleIds, ",")
		if helper.InAnySlice[string](roleList, vars.Config.GetString("server.superRoleId")) {
			return ctx.Next()
		}
		flag := false
		for _, sub := range roleList {
			//	判断策略中是否存在
			if ok, _ := e.Enforce(sub, obj, act); ok {
				flag = true
				break
			}
		}
		if !flag {
			return response.ForbiddenException(ctx, "该用户无此权限")
		}
		return ctx.Next()
	}
}

// ParamsActMatchFunc 自定义规则函数 method
func ParamsActMatchFunc(args ...interface{}) (interface{}, error) {
	rAct := args[0].(string)
	pAct := args[1].(string)
	pActArr := strings.Split(pAct, ",")
	if len(pActArr) > 1 {
		return helper.InAnySlice[string](pActArr, rAct), nil
	}
	return pActArr[0] == rAct, nil
}

// ParamsObjMatchFunc 自定义规则函数 path
func ParamsObjMatchFunc(args ...interface{}) (interface{}, error) {
	rObj := args[0].(string)
	pObj := args[1].(string)
	rObj1 := strings.Split(rObj, "?")[0]
	return util.KeyMatch2(rObj1, pObj), nil
}
