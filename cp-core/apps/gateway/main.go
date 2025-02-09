package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/lianzhilu/chat-paper/cp-core/apps/gateway/handler"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
)

var (
	serviceConfig = config.GetRuntimeConfig()
	gatewayAddr   = fmt.Sprintf("%s:%s", serviceConfig.GatewayConfig.GateWayHost, serviceConfig.GatewayConfig.GateWayPort)
)

func registerGroup(hz *server.Hertz) {
	chatpaper := hz.Group("/v1")
	{
		user := chatpaper.Group("/user")
		{
			user.POST("/register", handler.Register)
			user.POST("/login", handler.Login)
		}
		article := chatpaper.Group("/article")
		{
			article.POST("/", handler.CreateArticle)
			article.GET("/", handler.GetArticle)
			article.PUT("/", handler.UpdateArticle)
			article.DELETE("/", handler.DeleteArticle)
			article.GET("/list", handler.ListArticles)
		}
		comment := chatpaper.Group("/comment")
		{
			comment.POST("/", handler.CreateComment)
			comment.GET("/", handler.GetComment)
			comment.PUT("/", handler.UpdateComment)
			comment.DELETE("/", handler.DeleteComment)
		}
	}
}

func initHertz() *server.Hertz {
	hz := server.Default(server.WithHostPorts(gatewayAddr))
	return hz
}

func main() {
	hz := initHertz()
	registerGroup(hz)
	hz.Spin()
}
