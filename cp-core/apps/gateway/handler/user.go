package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lianzhilu/chat-paper/cp-core/apps/gateway/rpc"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user"
	"net/http"
)

func HealthCheck(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, user.RegisterRequest{
		ID: "health",
	})
}

func Register(ctx context.Context, c *app.RequestContext) {
	fmt.Println("register")
	req := &user.RegisterRequest{
		ID:       "xxxx",
		Password: "123456",
		Name:     "aaa",
	}
	resp, _ := rpc.Register(ctx, req)
	c.JSON(http.StatusOK, resp)
}
