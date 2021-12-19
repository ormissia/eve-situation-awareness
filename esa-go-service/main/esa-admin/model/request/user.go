package request

type Login struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
	// TODO 验证码相关添加 binding:"required"
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type Register struct {
	Username     string   `json:"userName" binding:"required"`
	Password     string   `json:"passWord" binding:"required"`
	NickName     string   `json:"nickName" binding:"required" gorm:"default:'NewUser'"`
	HeaderImg    string   `json:"headerImg" gorm:"default:'head.jpg'"`
	AuthorityId  string   `json:"authorityId" gorm:"default:1"`
	AuthorityIds []string `json:"authorityIds"`
}
