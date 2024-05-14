package rbac

import (
	"errors"
	"time"

	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/app/pkg/pagination"
	"github.com/MQEnergy/mqshop/internal/app/service"
	"github.com/MQEnergy/mqshop/internal/request/role"
)

type RoleService struct {
	service.Service
}

var Role = &RoleService{}

// Index 获取列表
func (s *RoleService) Index(reqParams role.IndexReq) (*pagination.PaginateResp, error) {
	parsePage := pagination.New().ParsePage(reqParams.Page, reqParams.Limit)
	roleList, count, err := dao.CRole.Order(dao.CRole.ID.Desc()).FindByPage(parsePage.GetOffset(), parsePage.GetLimit())
	if err != nil {
		return nil, err
	}
	parsePage.Total = count
	parsePage.GetLastPage()
	parsePage.List = roleList
	return &parsePage, nil
}

// CreateOrUpdate 创建
func (s *RoleService) CreateOrUpdate(reqParams role.CreateOrUpdateReq) error {
	var (
		roleInfo model.CRole
		err      error
	)
	if reqParams.Id == 0 {
		_, err = dao.CRole.Where(dao.CRole.Name.In(reqParams.Name)).First()
		if err == nil {
			return errors.New("角色名称已存在")
		}
	} else {
		// update
		roleInfo.ID = reqParams.Id
	}
	roleInfo.Name = reqParams.Name
	roleInfo.Desc = reqParams.Desc
	roleInfo.Status = reqParams.Status
	roleInfo.CreatedAt = int(time.Now().Unix())
	return dao.CRole.Save(&roleInfo)
}

// Delete 状态禁止
func (s *RoleService) Delete(reqParams role.DeleteReq) error {
	_, err := dao.CRole.Where(dao.CRole.ID.In(reqParams.Id)).Delete()
	if err != nil {
		return err
	}
	return nil
}

// View
// @Description: 详情
// @receiver s
// @param reqParams
// @return *model.CRole
// @return error
// @author cx
func (s *RoleService) View(reqParams role.ViewReq) (*model.CRole, error) {
	return dao.CRole.Where(dao.CRole.ID.Eq(reqParams.Id)).First()
}

// List
// @Description: 获取全部列表
// @receiver s
// @return []model.CRole
// @author cx
func (s *RoleService) List() []*model.CRole {
	list, _ := dao.CRole.Where(dao.CRole.ID.Neq(1), dao.CRole.Status.Eq(1)).Find()
	return list
}
