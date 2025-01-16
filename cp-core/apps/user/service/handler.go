package service

import (
	"context"
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user"
)

type UserServiceImpl struct{}

func (impl *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	testResp := fmt.Sprintf("%s:%s", req.Name, req.Password)
	return &user.RegisterResponse{
		ID: testResp,
	}, nil
}
