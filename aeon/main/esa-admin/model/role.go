package model

import (
	"fmt"

	"aeon/global"
)

type Role struct {
	ESABase
	Rolename      string `json:"rolename" gorm:"column:rolename"`
	ParentRoleId  uint   `json:"parent_role_id" gorm:"column:parent_role_id"`
	ChildrenRoles []Role `json:"children_roles" gorm:"-"`
	// BaseMenus
	DefaultRouter string `json:"default_router" gorm:"column:default_router"`
}

func (Role) TableName() string {
	return "role"
}

func (r *Role) Select(ids []uint, rolename string, parentRoleId uint, pageNo, pageSize int) (total int64, result []Role, err error) {
	db := global.ESAMySqlESA.Model(r)

	if len(ids) > 0 {
		db = db.Where("id in ?", ids)
	}
	if rolename != "" {
		db = db.Where("rolename like ?", fmt.Sprintf("%%%s%%", rolename))
	}
	db = db.Where("parent_role_id = ?", parentRoleId)

	db.Count(&total)

	db = db.Order("id")

	if pageNo != 0 && pageSize != 0 {
		db = db.Limit(pageSize).Offset((pageNo - 1) * pageSize)
	}
	err = db.Find(&result).Error
	return
}

func (r *Role) Creat() (err error) {
	db := global.ESAMySqlESA.Model(r)

	err = db.Create(r).Error
	return
}

func (r *Role) Update() (err error) {
	db := global.ESAMySqlESA.Model(r)

	db = db.Where("id = ?", r.ID)
	err = db.Updates(map[string]interface{}{
		"rolename":       r.Rolename,
		"parent_role_id": r.ParentRoleId,
		"update_time":    r.UpdateTime,
	}).Error
	return
}
