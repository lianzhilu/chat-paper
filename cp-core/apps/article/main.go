package main

import (
	"fmt"
	"github.com/cloudwego/kitex/server"
	"github.com/lianzhilu/chat-paper/cp-core/apps/article/service"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article/articleservice"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/database"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/repository"
	"net"
)

var (
	serviceConfig = config.GetRuntimeConfig()
	serviceAddr   = fmt.Sprintf("%s:%s", serviceConfig.ArticleServiceConfig.ArticleServiceHost, serviceConfig.ArticleServiceConfig.ArticleServicePort)
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		panic(err)
	}
	articleRepo := repository.NewMySQLArticleRepository(database.GetDB())
	articleService := service.NewArticleServiceImpl(articleRepo)
	svc := articleservice.NewServer(
		articleService,
		server.WithServiceAddr(addr),
	)
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
