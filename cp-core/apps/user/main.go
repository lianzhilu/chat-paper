package main

import (
	"fmt"
	"github.com/cloudwego/kitex/server"
	"github.com/lianzhilu/chat-paper/cp-core/apps/user/service"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user/userservice"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/database"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/repository"
	"net"
)

var (
	serviceConfig = config.GetRuntimeConfig()
	serviceAddr   = fmt.Sprintf("%s:%s", serviceConfig.UserServiceConfig.UserServiceHost, serviceConfig.UserServiceConfig.UserServicePort)
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		panic(err)
	}
	userRepo := repository.NewUserRepository(database.GetDB())
	userService := service.NewUserServiceImpl(userRepo)
	svc := userservice.NewServer(
		userService,
		server.WithServiceAddr(addr),
	)
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
