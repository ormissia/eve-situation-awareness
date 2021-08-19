package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"eas-go-service/global"
)

type User struct {
	EASBase
	UUID      uuid.UUID `json:"uuid" gorm:"column:uuid"`                                                                  // 用户UUID
	Username  string    `json:"userName" gorm:""`                                                                         // 用户登录名
	Password  string    `json:"-" gorm:""`                                                                                // 用户登录密码
	NickName  string    `json:"nickName" gorm:"default:系统用户"`                                                             // 用户昵称
	HeaderImg string    `json:"headerImg" gorm:"default:https://imageserver.eveonline.com/Character/2115581995_1024.jpg"` // 用户头像
	// Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId"`
	AuthorityId string `json:"authorityId" gorm:"default:888"` // 用户角色ID
	// Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}

func (User) TableName() string {
	return "user"
}

// AfterFind 每次查询后清空密码
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.Password = ""
	return
}

func (u *User) Login() (user User, err error) {
	db := global.EASMySql.Model(u)

	db = db.Where("username = ?", u.Username).Where("password = ?", u.Password)
	err = db.Find(&user).Error

	return
}

func (u *User) SelectUserByUUID(uuid string) (user User, err error) {
	db := global.EASMySql.Model(u)

	db = db.Where("uuid = ?", uuid)
	err = db.First(&user).Error

	return
}
