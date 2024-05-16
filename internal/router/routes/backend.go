package routes

import (
	"github.com/MQEnergy/mqshop/internal/app/controller/backend"
	"github.com/MQEnergy/mqshop/internal/middleware"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/database"
	"github.com/gofiber/fiber/v2"
)

// InitBackendGroup 初始化后台接口路由
func InitBackendGroup(r fiber.Router, handles ...fiber.Handler) {
	prefix := vars.Config.Get("database.mysql.sources." + database.DefaultAlias + ".prefix")
	backendHandles := append(handles, middleware.CasbinMiddleware(vars.DB, prefix.(string), "c_casbin_rule"))
	router := r.Group("backend", backendHandles...)
	{
		// 登录
		router.Post("/auth/login", backend.Auth.Login)
		// 退出登录
		router.Post("/auth/logout", backend.Auth.Logout)
		// 修改密码（修改自己的）
		router.Post("/auth/change-pass", backend.Auth.ChangePass)
		// 上传附件
		router.Post("/attachment/upload", backend.Attachment.Upload)
		// 忘记密码
		router.Post("/auth/forget-pass", backend.Auth.ForgetPass)

		// ============================== RBAC ==============================
		// ============================== 角色 ==============================
		// 创建角色
		router.Post("/role/create", backend.Role.CreateOrUpdate)
		// 更新角色
		router.Post("/role/update", backend.Role.CreateOrUpdate)
		// 删除角色
		router.Post("/role/delete", backend.Role.Delete)
		// 获取角色列表
		router.Get("/role/index", backend.Role.Index)
		// 获取角色详情
		router.Get("/role/view", backend.Role.View)
		// 获取全部角色列表
		router.Get("/role/list", backend.Role.List)
		// 邀请关联角色
		router.Post("/role/invite", backend.Role.Invite)

		// ============================== 资源 ==============================
		// 创建资源
		router.Post("/resource/create", backend.Resource.CreateOrUpdate)
		// 更新资源
		router.Post("/resource/update", backend.Resource.CreateOrUpdate)
		// 删除资源
		router.Post("/resource/delete", backend.Resource.Delete)
		// 获取资源列表
		router.Get("/resource/index", backend.Resource.Index)
		// 获取资源列表
		router.Get("/resource/list", backend.Resource.ResourceList)
		// 获取资源详情
		router.Get("/resource/view", backend.Resource.View)

		// ============================== 权限 ==============================
		// 创建权限
		router.Post("/permission/create", backend.Permission.Create)
		// 获取登录后权限详情（前端权限）
		router.Get("/permission/info", backend.Permission.Alias)
		// 获取权限详情
		router.Get("/permission/view", backend.Permission.View)

		// ============================== 人员 ==============================
		// 获取登录用户信息
		router.Get("/admin/view", backend.Admin.View)
		// 创建人员
		router.Post("/admin/create", backend.Admin.Create)
		// 更新人员
		router.Post("/admin/update", backend.Admin.Update)
		// 删除人员
		router.Post("/admin/delete", backend.Admin.Delete)
		// 获取人员列表
		router.Get("/admin/index", backend.Admin.Index)
		// 根据ID获取人员信息
		router.Get("/admin/info", backend.Admin.Info)
		// 分配角色列表
		router.Post("/admin/role-distribution", backend.Admin.RoleDistribution)
		// 修改密码 (修改其他人的)
		router.Post("/admin/change-pass", backend.Admin.ChangePassword)
		// ============================== RBAC ==============================
	}

	// casbin中间件可根据不同的数据库进行单独配置 示例如下：
	// demo数据库中存在casbin_rule
	//prefix := vars.Config.Get("database.mysql.sources.demo.prefix")
	//demoHandles := append(handles, middleware.CasbinMiddleware(vars.MDB["demo"], prefix.(string), ""))
	//router1 := r.Group("demo", demoHandles...)
	//{
	//	router1.Get("/", func(ctx *fiber.Ctx) error { return nil })
	//}

	// demo1数据库中存在casbin_rule
	//prefix := vars.Config.Get("database.mysql.sources.demo1.prefix")
	//demo1Handles := append(handles, middleware.CasbinMiddleware(vars.MDB["demo1"], prefix.(string), ""))
	//router2 := r.Group("demo1", demo1Handles...)
	//{
	//	router2.Get("/", func(ctx *fiber.Ctx) error { return nil })
	//}
}
