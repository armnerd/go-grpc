package main

import (
	"flag"
	"go-grpc/internal/base"
	"go-grpc/internal/server"
	"os"
	"os/signal"
	"syscall"
)

var (
	env string
)

func init() {
	flag.StringVar(&env, "env", "dev", "config path, eg: -env dev")
}

// App 应用
type App struct {
	Grpc server.IServerGrpc
	Http server.IServerHttp
}

func (app *App) start() {
	go func() {
		app.Grpc.Start()
	}()
	go func() {
		app.Http.Start()
	}()
	waitQuit()
}

func (app *App) stop() {
	app.Grpc.Stop()
	app.Http.Stop()
}

func main() {
	flag.Parse()
	base.Env = env
	app, cleanUp, err := NewApp()
	if err != nil {
		panic(err)
	}
	defer cleanUp()
	defer app.stop()
	app.start()
}

func waitQuit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
