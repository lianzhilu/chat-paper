package repository

import (
	"context"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/dal"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/model"
	"gorm.io/gorm"
	"time"
)

type CreateUserParams struct {
	ID       string
	Name     string
	Password string
}

type UserRepository interface {
	CreateUser(ctx context.Context, param *CreateUserParams) error
	GetUser(ctx context.Context, id string) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type MySQLUserRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *MySQLUserRepository {
	dal.SetDefault(db)
	return &MySQLUserRepository{db}
}

func (r *MySQLUserRepository) CreateUser(ctx context.Context, param *CreateUserParams) error {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := dal.User.WithContext(tCtx)

	user := &model.User{
		SID:      param.ID,
		Name:     param.Name,
		Password: param.Password,
	}
	if err := query.Create(user); err != nil {
		return err
	}
	return nil
}

func (r *MySQLUserRepository) GetUser(ctx context.Context, id string) (*model.User, error) {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	userOrm := dal.User
	query := userOrm.WithContext(tCtx)
	user, err := query.Where(userOrm.SID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MySQLUserRepository) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	userOrm := dal.User
	query := userOrm.WithContext(tCtx)
	user, err := query.Where(userOrm.Name.Eq(name)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}
