package storage

type Storage interface {
	// Save TODO 需要修改存储类型
	Save(msg interface{})
}
