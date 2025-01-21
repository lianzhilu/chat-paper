package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/generator"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/repository"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserServiceImpl(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (impl *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	_, err = impl.userRepo.GetUserByName(ctx, req.Name)
	if err == nil {
		return nil, fmt.Errorf("username %s already exists", req.Name)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	id := generator.GenerateUserID()
	param := repository.CreateUserParams{
		ID:       id,
		Name:     req.Name,
		Password: req.Password,
	}
	err = impl.userRepo.CreateUser(ctx, &param)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{
		ID:   id,
		Name: req.Name,
	}, nil
}

func (impl *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	loginUser, err := impl.userRepo.GetUserByName(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user %s not found", req.ID)
		} else {
			return nil, err
		}
	}
	if loginUser.Password != req.Password {
		return nil, fmt.Errorf("user %s password error, please retry", req.ID)
	}
	return &user.LoginResponse{
		ID:   loginUser.SID,
		Name: loginUser.Name,
	}, nil
}
