package rbac

import (
	"context"
	"errors"
	"github.com/spf13/cast"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/spf13/cast"

	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/app/pkg/pagination"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/app/service/sredis"
	"github.com/MQEnergy/mqshop/internal/request/admin"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/MQEnergy/mqshop/pkg/helper"
	"github.com/jinzhu/copier"
)

type AdminService struct {
	service.Service
}

var Admin = &AdminService{}

// Index 获取列表
func (s *AdminService) Index(reqParams admin.IndexReq) (*pagination.PaginateResp, error) {
	var (
		u           = dao.CAdmin
		parsePage   = pagination.New().ParsePage(reqParams.Page, reqParams.Limit)
		cnAdminList = make([]*admin.CAdmin, 0)
	)
	q := u.Order(u.ID.Desc())
	if reqParams.Search != "" {
		searchForm := make(map[string]interface{})
		if err := json.Unmarshal([]byte(reqParams.Search), &searchForm); err == nil {
			keyword := cast.ToString(searchForm["keyword"])
			status := cast.ToInt8(searchForm["status"])
			if keyword != "" {
				q.Where(u.Where(u.Account.Like("%" + keyword + "%")).
					Or(u.RealName.Like("%" + keyword + "%")).
					Or(u.Phone.Like("%" + keyword + "%")))
			}
			if searchForm["status"] != "" {
				q.Where(u.Status.Eq(status))
			}
		}
	}
	adminList, count, err := q.Debug().Omit(u.Password, u.Salt).FindByPage(parsePage.GetOffset(), parsePage.GetLimit())
	if err != nil {
		return nil, err
	}
	for i, item := range adminList {
		adminItem := &admin.CAdmin{
			CAdmin:   *item,
			RoleList: make([]*model.CRole, 0),
			IsSuper:  0,
		}
		cnAdminList = append(cnAdminList, adminItem)
		if item.RoleIds != "" {
			roleIds := strings.Split(item.RoleIds, ",")
			if len(roleIds) > 0 {
				roleIdList := helper.AnyToInts[string](roleIds)
				roles, err := dao.CRole.Where(dao.CRole.ID.In(roleIdList...)).Find()
				if err != nil {
					continue
				}
				cnAdminList[i].RoleList = roles
			}
		}
	}
	parsePage.Total = count
	parsePage.GetLastPage()
	parsePage.List = cnAdminList
	return &parsePage, nil
}

// Create ...
func (s *AdminService) Create(reqParams admin.CreateReq) error {
	var (
		adminInfo model.CAdmin
		uuid      = helper.GenerateBaseSnowId(0, nil)
		timeNow   = time.Now().Unix()
		salt      = helper.GenerateUuid(32)
		pass      = helper.GeneratePasswordHash(reqParams.Password, salt)
	)
	_, err := dao.CAdmin.Where(dao.CAdmin.Account.Eq(reqParams.Account)).First()
	if err == nil {
		return errors.New("账号已存在")
	}
	if err := copier.Copy(&adminInfo, reqParams); err != nil {
		return err
	}
	adminInfo.UUID = uuid
	adminInfo.RoleIds = reqParams.RoleIds
	adminInfo.Salt = salt
	adminInfo.Password = pass
	adminInfo.RegisterTime = timeNow
	return dao.CAdmin.Save(&adminInfo)
}

// Update ...
func (s *AdminService) Update(reqParams admin.UpdateReq) error {
	// 检测账号是否有重复的
	_, err := dao.CAdmin.Where(dao.CAdmin.ID.Neq(reqParams.Id), dao.CAdmin.Account.Eq(reqParams.Account)).First()
	if err == nil {
		return errors.New("当前账号已经存在")
	}
	adminInfo, err := dao.CAdmin.Where(dao.CAdmin.ID.Eq(reqParams.Id)).First()
	if err != nil {
		return err
	}
	if err := copier.Copy(&adminInfo, reqParams); err != nil {
		return err
	}
	return dao.CAdmin.Save(adminInfo)
}

// View ...
func (s *AdminService) View(uuid string) (*admin.CAdmin, error) {
	var (
		isSuper   = 0
		u         = dao.CAdmin
		adminInfo admin.CAdmin
		err       error
	)
	cAdminInfo, err := u.Select(u.ID, u.UUID, u.Account, u.Phone, u.Avatar, u.RealName, u.RoleIds).Where(u.UUID.Eq(uuid)).First()
	if err != nil {
		return nil, err
	}
	adminInfo.CAdmin = *cAdminInfo
	roleList := strings.Split(adminInfo.CAdmin.RoleIds, ",")
	if len(roleList) > 0 {
		if helper.InAnySlice[string](roleList, vars.Config.GetString("server.superRoleId")) {
			isSuper = 1
		}
	}
	adminInfo.IsSuper = isSuper
	return &adminInfo, nil
}

// Info ...
func (s *AdminService) Info(reqParams admin.InfoReq) (*model.CAdmin, error) {
	return dao.CAdmin.Omit(dao.CAdmin.Password, dao.CAdmin.Salt).Where(dao.CAdmin.ID.Eq(reqParams.Id)).First()
}

// Delete ...
func (s *AdminService) Delete(reqParams admin.DeleteReq) error {
	ids := strings.Split(reqParams.Ids, ",")
	adminIds := make([]int64, 0)
	for _, id := range ids {
		adminIds = append(adminIds, cast.ToInt64(id))
	}
	if _, err := dao.CAdmin.Where(dao.CAdmin.ID.In(adminIds...)).Delete(); err != nil {
		return errors.New("删除失败 err: " + err.Error())
	}
	return nil
}

// RoleDistribution 角色分配
func (s *AdminService) RoleDistribution(reqParams admin.RoleDistributionReq) error {
	u := dao.CAdmin
	adminInfo, err := u.Where(u.ID.Eq(reqParams.Id)).First()
	if err != nil {
		return err
	}
	// 删除redis的角色菜单关系
	redisKey := sredis.BuildAuthRedisKey(adminInfo.UUID) + ":*"
	keys, _ := vars.Redis.Keys(context.Background(), redisKey).Result()
	for _, key := range keys {
		vars.Redis.Del(context.Background(), key)
	}
	adminInfo.RoleIds = reqParams.RoleIds
	return u.Save(adminInfo)
}

// ChangePass 修改密码
func (s *AdminService) ChangePass(reqParams admin.ChangePasswordReq) error {
	adminInfo, err := dao.CAdmin.Where(dao.CAdmin.UUID.Eq(reqParams.Uuid)).First()
	if err != nil {
		return err
	}
	adminInfo.Salt = helper.GenerateUuid(32)
	adminInfo.Password = helper.GeneratePasswordHash(reqParams.NewPass, adminInfo.Salt)
	return dao.CAdmin.Save(adminInfo)
}
