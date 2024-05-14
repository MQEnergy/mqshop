package rbac

import (
	"context"
	"errors"
	"strings"
	"time"

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
	if reqParams.Keyword != "" {
		q.Where(u.Account.Like("%" + reqParams.Keyword + "%")).
			Or(u.RealName.Like("%" + reqParams.Keyword + "%")).
			Or(u.Phone.Like("%" + reqParams.Keyword + "%"))
	}
	adminList, count, err := q.FindByPage(parsePage.GetOffset(), parsePage.GetLimit())
	if err != nil {
		return nil, err
	}
	for i, item := range adminList {
		cnAdminList = append(cnAdminList, &admin.CAdmin{
			Info:     item,
			RoleList: make([]*model.CRole, 0),
			IsSuper:  0,
		})
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
	parsePage.List = adminList
	return &parsePage, nil
}

// Create ...
func (s *AdminService) Create(reqParams admin.CreateReq) error {
	var (
		uuid      = helper.GenerateBaseSnowId(0, nil)
		timeNow   = time.Now().Unix()
		adminInfo model.CAdmin
	)
	_, err := dao.CAdmin.Where(dao.CAdmin.Account.Eq(reqParams.Account)).First()
	if err == nil {
		return errors.New("账号已存在")
	}
	if err := copier.Copy(&adminInfo, reqParams); err != nil {
		return err
	}
	adminInfo.UUID = uuid
	adminInfo.RegisterTime = timeNow
	adminInfo.RoleIds = reqParams.RoleIds
	return dao.CAdmin.Save(&adminInfo)
}

// Update ...
func (s *AdminService) Update(reqParams admin.UpdateReq) error {
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
	adminInfo.Info, err = u.Select(u.ID, u.UUID, u.Account, u.Phone, u.Avatar, u.RealName, u.RoleIds).Where(u.UUID.Eq(uuid)).First()
	if err != nil {
		return nil, err
	}
	roleList := strings.Split(adminInfo.Info.RoleIds, ",")
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
	if _, err := dao.CAdmin.Where(dao.CAdmin.ID.Eq(reqParams.Id)).Delete(); err != nil {
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
