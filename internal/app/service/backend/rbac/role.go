package rbac

import (
	"errors"
	"github.com/goccy/go-json"
	"github.com/spf13/cast"
	"strings"
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
	var (
		u         = dao.CRole
		parsePage = pagination.New().ParsePage(reqParams.Page, reqParams.Limit)
	)
	q := dao.CRole.Order(u.ID.Desc())
	if reqParams.Search != "" {
		searchForm := make(map[string]interface{})
		if err := json.Unmarshal([]byte(reqParams.Search), &searchForm); err == nil {
			keyword := cast.ToString(searchForm["keyword"])
			status := cast.ToInt8(searchForm["status"])
			if keyword != "" {
				q.Where(u.Where(u.Name.Like("%" + keyword + "%")))
			}
			if searchForm["status"] != "" {
				q.Where(u.Status.Eq(status))
			}
		}
	}
	roleList, count, err := q.FindByPage(parsePage.GetOffset(), parsePage.GetLimit())
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
	ids := strings.Split(reqParams.Ids, ",")
	roleIds := make([]int, 0)
	for _, id := range ids {
		roleIds = append(roleIds, cast.ToInt(id))
	}
	if _, err := dao.CRole.Where(dao.CRole.ID.In(roleIds...)).Delete(); err != nil {
		return errors.New("删除失败 err: " + err.Error())
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
