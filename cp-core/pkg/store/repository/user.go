package repository

import (
	"context"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/model"
	"gorm.io/gorm"
)

type CreateUserParams struct {
	ID       string
	Name     string
	Password string
}

type UserRepository interface {
	CreateUser(ctx context.Context, param *CreateUserParams) error
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

type MySQLUserRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db}
}

func (r *MySQLUserRepository) CreateUser(ctx context.Context, param *CreateUserParams) error {
	db := r.DB.WithContext(ctx)
	user := &model.User{
		SID:      param.ID,
		Name:     param.Name,
		Password: param.Password,
	}
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySQLUserRepository) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	db := r.DB.WithContext(ctx)
	var user model.User
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	db := r.DB.WithContext(ctx)
	var user model.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
