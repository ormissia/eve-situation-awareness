package model

import (
	"fmt"

	"eas-go-service/global"
)

type Role struct {
	EASBase
	Rolename      string `json:"rolename" gorm:"column:rolename"`
	ParentRoleId  uint   `json:"parent_role_id" gorm:"column:parent_role_id"`
	ChildrenRoles []Role `json:"children_roles" gorm:"-"`
	// BaseMenus
	DefaultRouter string `json:"default_router" gorm:"column:default_router"`
}

func (Role) TableName() string {
	return "role"
}

func (r *Role) Select(ids []uint, rolename string) (result []Role, err error) {
	db := global.EASMySql.Model(r)

	if len(ids) > 0 {
		db = db.Where("id in ?", ids)
	}
	if rolename != "" {
		db = db.Where("rolename like ?", fmt.Sprintf("%%%s%%", rolename))
	}

	err = db.Find(&result).Error
	return
}

func (r *Role) Creat() (err error) {
	db := global.EASMySql.Model(r)

	err = db.Create(r).Error
	return
}
