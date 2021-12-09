package core

type Source interface {
	Listening(outPut func(msg []byte))
}
