package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lianzhilu/chat-paper/cp-core/apps/comment/service"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/comment/commentservice"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/database"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/repository"
	"net"
)

var (
	serviceConfig      = config.GetRuntimeConfig()
	commentServiceAddr = fmt.Sprintf("%s:%s", serviceConfig.CommentServiceConfig.CommentServiceHost, serviceConfig.CommentServiceConfig.CommentServicePort)
	etcdAddrs          = []string{fmt.Sprintf("%s:%s", serviceConfig.EtcdConfig.EtcdHost, serviceConfig.EtcdConfig.EtcdPort)}
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", commentServiceAddr)
	r, err := etcd.NewEtcdRegistry(etcdAddrs)
	if err != nil {
		panic(err)
	}
	commentRepo := repository.NewMySQLCommentRepository(database.GetDB())
	commentService := service.NewCommentService(commentRepo)
	svc := commentservice.NewServer(
		commentService,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
