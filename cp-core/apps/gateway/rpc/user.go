package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user/userservice"
	"time"
)

var (
	userClient         userservice.Client
	articleServiceAddr = fmt.Sprintf("%s:%s", serviceConfig.UserServiceConfig.UserServiceHost, serviceConfig.UserServiceConfig.UserServicePort)
)

func InitUserClient() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	c, err := userservice.NewClient(
		"user",
		client.WithMuxConnection(1),           // mux
		client.WithRPCTimeout(30*time.Second), // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithHostPorts(articleServiceAddr),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func Register(ctx context.Context, req *user.RegisterRequest) (r *user.RegisterResponse, err error) {
	return userClient.Register(ctx, req)
}

func Login(ctx context.Context, req *user.LoginRequest) (r *user.LoginResponse, err error) {
	return userClient.Login(ctx, req)
}
