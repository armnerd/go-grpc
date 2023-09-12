package server

type IServer interface {
	Start() error
	Stop() error
}

type IServerGrpc interface {
	IServer
}

type IServerHttp interface {
	IServer
}
