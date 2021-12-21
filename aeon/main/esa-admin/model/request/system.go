package request

type InitDB struct {
	Path     string `json:"path" binding:"required"`     // 服务器地址
	Username string `json:"username" binding:"required"` // 数据库用户名
	Password string `json:"password" binding:"required"` // 数据库密码
	DBName   string `json:"dbName" binding:"required"`   // 数据库名
}

type InitDBFunc interface {
	Init() (err error)
}
