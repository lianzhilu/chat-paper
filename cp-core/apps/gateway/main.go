package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lianzhilu/chat-paper/cp-core/apps/gateway/handler"
)

func registerGroup(hz *server.Hertz) {
	chatpaper := hz.Group("/chatpaper")
	{
		user := chatpaper.Group("/user")
		{
			user.POST("/register/", handler.Register)
		}
	}
}

func initHertz() *server.Hertz {
	hz := server.Default(server.WithHostPorts("localhost:9090"))
	return hz
}

func main() {
	hz := initHertz()
	registerGroup(hz)
	hz.Spin()
}
