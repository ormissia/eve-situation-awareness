package core

type Source interface {
	// TODO 根据情况修改这个回调函数接收[]byte还是string
	Listening(outPut func(msg string))
}
