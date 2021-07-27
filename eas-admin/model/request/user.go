package request

type Login struct {
	Username  string `json:"username"`  //用户名
	Password  string `json:"password"`  //密码
	Captcha   string `json:"captcha"`   //验证码
	CaptchaId string `json:"captchaId"` //验证码ID
}

type Register struct {
	Username     string   `json:"userName"`
	Password     string   `json:"passWord"`
	NickName     string   `json:"nickName" gorm:"default:'NewUser'"`
	HeaderImg    string   `json:"headerImg" gorm:"default:'head.jpg'"`
	AuthorityId  string   `json:"authorityId" gorm:"default:1"`
	AuthorityIds []string `json:"authorityIds"`
}
