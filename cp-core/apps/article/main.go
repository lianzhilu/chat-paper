package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lianzhilu/chat-paper/cp-core/apps/article/service"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article/articleservice"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/database"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/repository"
	"net"
)

var (
	serviceConfig      = config.GetRuntimeConfig()
	articleServiceAddr = fmt.Sprintf("%s:%s", serviceConfig.ArticleServiceConfig.ArticleServiceHost, serviceConfig.ArticleServiceConfig.ArticleServicePort)
	etcdAddrs          = []string{fmt.Sprintf("%s:%s", serviceConfig.EtcdConfig.EtcdHost, serviceConfig.EtcdConfig.EtcdPort)}
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", articleServiceAddr)
	r, err := etcd.NewEtcdRegistry(etcdAddrs)
	if err != nil {
		panic(err)
	}
	articleRepo := repository.NewMySQLArticleRepository(database.GetDB())
	articleService := service.NewArticleServiceImpl(articleRepo)
	svc := articleservice.NewServer(
		articleService,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "article"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
