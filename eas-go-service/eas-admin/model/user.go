package model

import (
	uuid "github.com/satori/go.uuid"

	"eas-go-service/global"
)

// User 用户相关
type User struct {
	EASBase
	UUID      uuid.UUID `json:"uuid" gorm:"column:uuid"`                                                                                    // 用户UUID
	Username  string    `json:"userName" gorm:"column:username"`                                                                            // 用户登录名
	Password  string    `json:"-" gorm:"column:password"`                                                                                   // 用户登录密码
	Nickname  string    `json:"nickName" gorm:"column:nickname;default:系统用户"`                                                               // 用户昵称
	HeaderImg string    `json:"headerImg" gorm:"column:header_img;default:https://imageserver.eveonline.com/Character/2115581995_1024.jpg"` // 用户头像
	RoleId    string    `json:"role_id" gorm:"role_id"`                                                                                     // 用户角色ID
	Roles     []Role    `json:"roles" gorm:"many2many:user_role"`
}

func (User) TableName() string {
	return "user"
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
	err = db.Preload("Roles").Find(&user).Error

	return
}
