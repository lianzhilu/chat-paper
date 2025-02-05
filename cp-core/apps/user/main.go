package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lianzhilu/chat-paper/cp-core/apps/user/service"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user/userservice"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/database"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/repository"
	"net"
)

var (
	serviceConfig   = config.GetRuntimeConfig()
	userServiceAddr = fmt.Sprintf("%s:%s", serviceConfig.UserServiceConfig.UserServiceHost, serviceConfig.UserServiceConfig.UserServicePort)
	etcdAddrs       = []string{fmt.Sprintf("%s:%s", serviceConfig.EtcdConfig.EtcdHost, serviceConfig.EtcdConfig.EtcdPort)}
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", userServiceAddr)
	r, err := etcd.NewEtcdRegistry(etcdAddrs)
	if err != nil {
		panic(err)
	}
	userRepo := repository.NewUserRepository(database.GetDB())
	userService := service.NewUserServiceImpl(userRepo)
	svc := userservice.NewServer(
		userService,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
