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
		router.Get("/attachment/index", backend.Attachment.Index)
		// 忘记密码
		router.Post("/auth/forget-pass", backend.Auth.ForgetPass)
		// 省市区
		router.Get("/region/province", backend.Region.Province)
		router.Get("/region/city", backend.Region.City)
		router.Get("/region/area", backend.Region.Area)
		router.Get("/region/street", backend.Region.Street)
		router.Get("/region/village", backend.Region.Village)

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

		// ============================== 商品 ==============================
		router.Get("/product/index", backend.Product.Index)
		router.Post("/product/delete", backend.Product.Delete)
		router.Post("/product/create", backend.Product.Create)
		router.Get("/product/view", backend.Product.View)
		router.Post("/product/status", backend.Product.Status)
		router.Get("/product/attrs", backend.Product.Index)
		router.Get("/product/attrs-list", backend.Product.Index)
		router.Post("/product/attrs-save", backend.Product.Index)
		router.Post("/product/sku-no-create", backend.Product.Index)

		// ============================== 商品分类 ==============================
		router.Get("/cate/index", backend.Cate.Index)
		router.Post("/cate/delete", backend.Cate.Delete)
		router.Post("/cate/create", backend.Cate.Create)
		router.Post("/cate/update", backend.Cate.Update)
		router.Get("/cate/view", backend.Cate.View)
		router.Post("/cate/status", backend.Cate.Status)
		router.Get("/cate/list", backend.Cate.List)

		// ============================== 商品品牌 ==============================
		router.Get("/brand/index", backend.Brand.Index)
		router.Post("/brand/delete", backend.Brand.Delete)
		router.Post("/brand/create", backend.Brand.Create)
		router.Post("/brand/update", backend.Brand.Update)
		router.Get("/brand/view", backend.Brand.View)
		router.Post("/brand/status", backend.Brand.Status)
		router.Get("/brand/list", backend.Brand.List)

		// ============================== 商品属性 ==============================
		router.Get("/attr-cates/index", backend.AttrCate.Index)
		router.Get("/attr-cates/list", backend.AttrCate.List)
		router.Get("/attr-cates/view", backend.AttrCate.View)
		router.Post("/attr-cates/update", backend.AttrCate.Update)
		router.Post("/attr-cates/delete", backend.AttrCate.Delete)
		router.Post("/attr-cates/create", backend.AttrCate.Create)
		// 属性列表 和 参数列表
		router.Post("/attr-cates/attr-list", backend.AttrCate.AttrParamsList)
		router.Post("/attr-cates/attr-create", backend.AttrCate.AttrParamsCreate)
		router.Post("/attr-cates/attr-update", backend.AttrCate.AttrParamsUpdate)
		router.Post("/attr-cates/attr-delete", backend.AttrCate.AttrParamsDelete)
	}
}
