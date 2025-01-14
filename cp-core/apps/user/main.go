package main

import (
	"github.com/lianzhilu/chat-paper/cp-core/apps/user/service"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user/userservice"
)

func main() {
	svc := userservice.NewServer(new(service.UserServiceImpl))
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
