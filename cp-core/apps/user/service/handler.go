package service

import (
	"context"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user"
)

type UserServiceImpl struct{}

func (impl *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	return &user.RegisterResponse{
		ID: "response",
	}, nil
}
