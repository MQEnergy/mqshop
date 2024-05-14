package rbac

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/MQEnergy/mqshop/internal/app/dao"
	"github.com/MQEnergy/mqshop/internal/app/model"
	"github.com/MQEnergy/mqshop/internal/request/resource"
	"github.com/MQEnergy/mqshop/internal/vars"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"gorm.io/gen"
)

type ResourceService struct{}

var Resource = &ResourceService{}

// ResourceNode Define a struct to hold the hierarchical data
type ResourceNode struct {
	ID        int64           `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name      string          `gorm:"column:name;type:varchar(64);not null;comment:菜单名称" json:"name"`                              // 菜单名称
	Alias_    string          `gorm:"column:alias;type:varchar(64);not null;comment:菜单别名" json:"alias"`                            // 菜单别名
	Desc      string          `gorm:"column:desc;type:varchar(64);not null;comment:菜单描述" json:"desc"`                              // 菜单描述
	FURL      string          `gorm:"column:f_url;type:varchar(64);not null;comment:菜单前端URL" json:"f_url"`                         // 菜单前端URL
	BURL      string          `gorm:"column:b_url;type:varchar(64);not null;comment:后端URL" json:"b_url"`                           // 后端URL
	Icon      string          `gorm:"column:icon;type:varchar(16);not null;comment:菜单icon" json:"icon"`                            // 菜单icon
	ParentID  int64           `gorm:"column:parent_id;type:bigint unsigned;not null;comment:父级ID" json:"parent_id"`                // 父级ID
	Path      string          `gorm:"column:path;type:varchar(32);not null;comment:ID路径 1-2-3..." json:"path"`                     // ID路径 1-2-3...
	MenuType  int8            `gorm:"column:menu_type;type:tinyint;not null;default:1;comment:菜单类型 1：模块 2：操作" json:"menu_type"`    // 菜单类型 1：模块 2：操作
	Status    int8            `gorm:"column:status;type:tinyint unsigned;not null;default:1;comment:状态：1正常(默认) 2停用" json:"status"` // 状态：1正常(默认) 2停用
	SortOrder int16           `gorm:"column:sort_order;type:smallint unsigned;not null;default:50;comment:排序" json:"sort_order"`   // 排序
	CreatedAt int64           `gorm:"column:created_at;type:bigint;not null" json:"created_at"`
	Children  []*ResourceNode `json:"children"`
}

// BuildHierarchy Function to build the hierarchical structure
func BuildHierarchy(resources []*model.CResource) []*ResourceNode {
	menuMap := make(map[int64]*ResourceNode)
	var roots []*ResourceNode

	for _, source := range resources {
		node := &ResourceNode{
			ID:        source.ID,
			Name:      source.Name,
			Alias_:    source.Alias_,
			Desc:      source.Desc,
			FURL:      source.FURL,
			BURL:      source.BURL,
			Icon:      source.Icon,
			ParentID:  source.ParentID,
			Path:      source.Path,
			MenuType:  source.MenuType,
			Status:    source.Status,
			SortOrder: source.SortOrder,
			CreatedAt: source.CreatedAt,
			Children:  []*ResourceNode{},
		}
		menuMap[source.ID] = node
	}
	// Build the hierarchical structure by linking parent and child nodes
	for _, source := range resources {
		node := menuMap[source.ID]
		if source.ParentID == 0 {
			roots = append(roots, node)
		} else {
			parentNode := menuMap[source.ParentID]
			if parentNode != nil {
				parentNode.Children = append(parentNode.Children, node)
			}
		}
	}
	return roots
}

// SortCategories sorts the categories
func SortCategories(categories []*model.CResource, parentID int64, level int) []*resource.Item {
	var sorted []*resource.Item
	for _, category := range categories {
		if category.ParentID == parentID {
			menuItem := &resource.Item{
				ID:       category.ID,
				ParentID: category.ParentID,
				Label:    fmt.Sprintf("%s%s", strings.Repeat("- ", level), category.Name),
				Name:     category.Name,
				Alias_:   category.Alias_,
				Level:    level,
				Path:     category.Path,
				MenuType: category.MenuType,
			}
			sorted = append(sorted, menuItem)
			sorted = append(sorted, SortCategories(categories, category.ID, level+1)...)
		}
	}
	return sorted
}

// BuildMenuItems Function to build the menu items
func BuildMenuItems(items []*resource.Item, menuIds []string) []*resource.Item {
	var menu1List []*resource.Item
	var menu2List []*resource.Item
	for _, item := range items {
		if item.MenuType == 2 {
			menu2List = append(menu2List, item)
		} else {
			menu1List = append(menu1List, item)
		}
	}
	for _, item := range menu1List {
		for _, menuItem := range menu2List {
			if item.ID == menuItem.ParentID {
				item.Actions = append(item.Actions, menuItem)
				if len(menuIds) > 0 {
					if lo.IndexOf(menuIds, strconv.Itoa(int(menuItem.ID))) >= 0 {
						item.CheckData = append(item.CheckData, int(menuItem.ID))
					}
				}
			}
		}
		if len(item.CheckData) > 0 {
			if len(item.CheckData) == len(item.Actions) {
				item.CheckedAll = true
			} else {
				item.Indeterminate = true
			}
		}
	}
	return menu1List
}

// Index 获取列表
func (s *ResourceService) Index(reqParams resource.IndexReq) ([]*ResourceNode, error) {
	resourceList, err := dao.CResource.Order(dao.CResource.ID.Asc(), dao.CResource.SortOrder.Asc()).Find()
	if err != nil {
		return nil, err
	}
	return BuildHierarchy(resourceList), nil
}

// MenuList 获取菜单列表
func (s *ResourceService) MenuList(reqParams resource.ListReq) ([]*resource.Item, error) {
	var (
		whereCondition = []gen.Condition{
			dao.CResource.MenuType.Eq(1),
		}
		exceptIds = make([]int64, 0)
	)
	if reqParams.Id != 0 {
		//	查询排除当前ID下的所有分类
		parentMenuList, err := dao.CResource.Where(dao.CResource.ParentID.Eq(reqParams.Id), dao.CResource.MenuType.Eq(1)).Find()
		if err != nil {
			return nil, err
		}
		exceptIds = append(exceptIds, reqParams.Id)
		for _, menuItem := range parentMenuList {
			exceptIds = append(exceptIds, menuItem.ID)
		}
		whereCondition = append(whereCondition, dao.CResource.ID.NotIn(exceptIds...))
	}
	resourceList, err := dao.CResource.Where(whereCondition...).Find()
	if err != nil {
		return nil, err
	}
	return SortCategories(resourceList, 0, 0), nil
}

// CreateOrUpdate
// @Description: 新增和编辑
// @receiver s
// @param reqParams
// @return error
// @author cx
func (s *ResourceService) CreateOrUpdate(reqParams resource.CreateOrUpdateReq) error {
	var (
		menuInfo           model.CResource
		resourceList       = make([]*model.CResource, 0)
		updateResourceList = make([]*model.CResource, 0)
		err                error
	)
	if reqParams.Id == 0 {
		_, err = dao.CResource.Where(dao.CResource.Alias_.Eq(reqParams.Alias_)).First()
		if err == nil {
			return errors.New("菜单别名已存在")
		}
	} else {
		_menuInfo, err := dao.CResource.Where(dao.CResource.ID.Eq(reqParams.Id)).First()
		if err != nil {
			return errors.New("菜单不存在")
		}
		// 找到下游所有的菜单path修改
		resourceList, err = dao.CResource.Where(dao.CResource.Path.Like(_menuInfo.Path + "-%")).Find()
		if err != nil {
			return err
		}
		// 替换所有查找到的path
		for _, source := range resourceList {
			newPath := strings.Replace(source.Path, _menuInfo.Path+"-", fmt.Sprintf("%s-%d-", reqParams.ParentPath, _menuInfo.ID), -1)
			source.Path = newPath
			updateResourceList = append(updateResourceList, source)
		}
		menuInfo = *_menuInfo
	}
	if err := copier.Copy(&menuInfo, reqParams); err != nil {
		return err
	}
	return dao.Use(vars.DB).Transaction(func(tx *dao.Query) error {
		// 保持当前分类
		if err := tx.CResource.Save(&menuInfo); err != nil {
			return errors.New("保存失败 err:" + err.Error())
		}
		menuPath := strconv.Itoa(int(menuInfo.ID))
		if reqParams.ParentPath != "" {
			menuPath = reqParams.ParentPath + "-" + menuPath
		}
		// 修改当前分类的path
		if _, err := tx.CResource.Where(tx.CResource.ID.Eq(menuInfo.ID)).Update(tx.CResource.Path, menuPath); err != nil {
			return errors.New("保存失败 err:" + err.Error())
		}
		if reqParams.Id != 0 && len(updateResourceList) > 0 {
			// 修改当前分类下的所有子分类的path
			if err := tx.CResource.Save(updateResourceList...); err != nil {
				return errors.New("保存失败 err:" + err.Error())
			}
		}
		return nil
	})
}

// Delete
// @Description: 删除菜单
// @receiver s
// @param reqParams
// @return error
// @author cx
func (s *ResourceService) Delete(reqParams resource.DeleteReq) error {
	menuInfo, err := dao.CResource.Where(dao.CResource.ID.Eq(reqParams.Id)).First()
	if err != nil {
		return err
	}
	return dao.Use(vars.DB).Transaction(func(tx *dao.Query) error {
		_, err := tx.CResource.Where(tx.CResource.Path.Like(menuInfo.Path + "-%")).Delete()
		if err != nil {
			return err
		}
		if _, err := tx.CResource.Delete(menuInfo); err != nil {
			return err
		}
		return nil
	})
}

// View
// @Description: 详情
// @receiver s
// @param reqParams
// @return *model.CResource
// @return error
// @author cx
func (s *ResourceService) View(reqParams resource.ViewReq) (*model.CResource, error) {
	return dao.CResource.Where(dao.CResource.ID.Eq(reqParams.Id)).First()
}
