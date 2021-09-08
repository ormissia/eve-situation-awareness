package model

// UserRole 仅用作数据库初始化
type UserRole struct {
	UserId uint `gorm:"user_id"`
	RoleId uint `gorm:"role_id"`
}

func (UserRole) TableName() string {
	return "user_role"
}
