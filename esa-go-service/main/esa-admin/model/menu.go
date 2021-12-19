package model

type SysBaseMenu struct {
	ESABase
	ParentId  string        `json:"parentId" gorm:"comment:父菜单ID"`     // 父菜单ID
	Name      string        `json:"name" gorm:"comment:路由name"`        // 路由name
	Path      string        `json:"path" gorm:"comment:路由path"`        // 路由path
	Hidden    bool          `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component string        `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Children  []SysBaseMenu `json:"children" gorm:"-"`
}
