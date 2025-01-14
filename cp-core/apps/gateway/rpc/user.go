package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user/userservice"
	"time"
)

var userClient userservice.Client

func InitUserClient() {
	c, err := userservice.NewClient(
		"user",
		client.WithMuxConnection(1),           // mux
		client.WithRPCTimeout(30*time.Second), // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithHostPorts("localhost:8888"),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func Register(ctx context.Context, req *user.RegisterRequest) (r *user.RegisterResponse, err error) {
	return userClient.Register(ctx, req)
}
