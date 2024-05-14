package rbac

import (
	"context"
	"encoding/json"
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/app/service/sredis"
	"github.com/MQEnergy/mqshop/internal/request/permission"
	"github.com/MQEnergy/mqshop/internal/request/resource"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/helper"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

type PermissionService struct {
	service.Service
}

var Permission = &PermissionService{}

// View 查看角色权限
func (s *PermissionService) View(reqParams permission.ViewReq) ([]*resource.Item, []int, error) {
	var (
		menuIds    = make([]string, 0)
		menuIntIds = make([]int, 0)
		u          = dao.CRole
		p          = dao.CRoleAuth
		m          = dao.CResource
	)
	_, err := u.Where(u.ID.Eq(reqParams.RoleId)).First()
	if err != nil {
		return nil, menuIntIds, err
	}
	// 查询角色分类的菜单ID列表
	if authInfo, err := p.Where(p.RoleID.Eq(reqParams.RoleId)).First(); err == nil {
		if authInfo.MenuIds != "" {
			menuIds = append(menuIds, strings.Split(authInfo.MenuIds, ",")...)
		}
	}
	// 查找所有分类和操作
	menuList, err := m.Where(m.Status.Eq(1)).Find()
	if err != nil {
		return nil, menuIntIds, err
	}
	// 找出非操作分类列表
	for _, menu := range menuList {
		if menu.MenuType == 1 {
			if lo.IndexOf(menuIds, strconv.Itoa(int(menu.ID))) >= 0 {
				menuIntIds = append(menuIntIds, cast.ToInt(menu.ID))
			}
		}
	}
	// 查询到所有的菜单ID列表
	return BuildMenuItems(SortCategories(menuList, 0, 0), menuIds), menuIntIds, nil
}

// Create 创建角色权限
func (s *PermissionService) Create(reqParams permission.CreateReq) error {
	var (
		pType       = "p"
		v0          = strconv.Itoa(reqParams.RoleId)
		casbinRoles = []*model.CCasbinRule{{
			Ptype: helper.ToPointer("p"),
			V0:    helper.ToPointer(v0),
			V1:    helper.ToPointer("/backend/auth/*"),
			V2:    helper.ToPointer("GET,POST"),
		}}
		rolePermissions = make([]*model.CRoleAuth, 0)
	)
	if reqParams.RoleId == vars.Config.GetInt("server.superRoleId") {
		return errors.New("超级管理员无需添加权限")
	}
	if _, err := dao.CRole.Where(dao.CRole.ID.Eq(reqParams.RoleId)).First(); err != nil {
		return err
	}
	if len(reqParams.MenuIds) == 0 {
		return errors.New("请至少选择一个菜单")
	}
	menuIds := strings.Split(reqParams.MenuIds, ",")
	menuIdList := gconv.Int64s(lo.Uniq[string](menuIds))
	menuList, err := dao.CResource.Where(dao.CResource.ID.In(menuIdList...)).Find()
	if err != nil {
		return err
	}
	rolePermissions = append(rolePermissions, &model.CRoleAuth{
		RoleID:  reqParams.RoleId,
		MenuIds: helper.IntsToString[int64](menuIdList, ","),
	})
	for _, item := range menuList {
		if item.BURL == "/" || item.BURL == "" {
			continue
		}
		casbinRoles = append(casbinRoles, &model.CCasbinRule{
			Ptype: helper.ToPointer(pType),
			V0:    helper.ToPointer(v0),
			V1:    &item.BURL,
			V2:    helper.ToPointer("GET,POST,DELETE"),
		})
	}
	return dao.Use(vars.DB).Transaction(func(tx *dao.Query) error {
		var (
			r = tx.CCasbinRule
			a = tx.CRoleAuth
		)
		if _, err := r.Where(r.V0.Eq(strconv.Itoa(reqParams.RoleId))).Delete(); err != nil {
			return err
		}
		if _, err := a.Where(a.RoleID.Eq(reqParams.RoleId)).Delete(); err != nil {
			return err
		}
		if len(casbinRoles) > 0 {
			if err := r.Save(casbinRoles...); err != nil {
				return err
			}
		}
		// 新增role_auth
		if err := a.Save(rolePermissions...); err != nil {
			return err
		}
		return nil
	})
}

// Alias
// @Description: 获取角色的前端的别名列表
// @receiver s
// @param uuid
// @return []string
// @return error
// @author cx
func (s *PermissionService) Alias(uuid string) ([]string, error) {
	var (
		menuIds     = make([]int64, 0)
		alias       = make([]string, 0)
		tokenExpire = vars.Config.GetDuration("jwt.expire")
		redisKey    = sredis.BuildAuthRedisKey(uuid) + ":"
		u           = dao.CRoleAuth
	)
	adminInfo, err := dao.CAdmin.Where(dao.CAdmin.UUID.Eq(uuid)).First()
	if err != nil {
		return nil, err
	}
	roleIds := strings.Split(adminInfo.RoleIds, ",")
	if len(roleIds) == 0 {
		return nil, errors.New("当前用户未分配角色")
	}
	// 超级管理员
	if helper.InAnySlice[string](roleIds, vars.Config.GetString("server.superRoleId")) {
		return []string{}, nil
	}
	if len(roleIds) == 1 {
		redisKey += roleIds[0]
	} else {
		sort.Strings(roleIds)
		redisKey += strings.Join(roleIds, "_")
	}
	menuRedisList, err := vars.Redis.Get(context.Background(), redisKey).Result()
	if err == nil && menuRedisList != "" {
		if err := json.Unmarshal([]byte(menuRedisList), &alias); err != nil {
			return nil, errors.New("用户菜单权限获取失败 err: " + err.Error())
		}
		return alias, nil
	}
	roleIdList := helper.String2Int(roleIds)
	rolePermissions, err := u.Where(u.RoleID.In(roleIdList...)).Find()
	if err != nil {
		return nil, err
	}
	for _, role := range rolePermissions {
		roleMenuIds := strings.Split(role.MenuIds, ",")
		roleMIds := gconv.Int64s(roleMenuIds)
		menuIds = append(menuIds, roleMIds...)
	}
	menuIds = lo.Uniq(menuIds)
	menuList, err := dao.CResource.Where(dao.CResource.ID.In(menuIds...)).Find()
	if err != nil {
		return nil, err
	}
	for _, menu := range menuList {
		alias = append(alias, menu.Alias_)
	}
	// 存入redis
	marshal, _ := json.Marshal(alias)
	if err := vars.Redis.Set(context.Background(), redisKey, string(marshal), tokenExpire*time.Second).Err(); err != nil {
		return []string{}, errors.New("用户信息持久化失败 err: " + err.Error())
	}
	return alias, nil
}
