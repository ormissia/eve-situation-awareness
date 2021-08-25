package model

type UserRole struct {
	UserId uint `gorm:"user_id"`
	RoleId uint `gorm:"role_id"`
}

func (UserRole) TableName() string {
	return "user_role"
}
