package handler

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lianzhilu/chat-paper/cp-core/apps/gateway/rpc"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/user"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/errmsg"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/response"
	"net/http"
	"strings"
)

var notSupportedItems = []string{"@", "#", "=", "*", "^"}

func validateName(name string) error {
	if len(name) < 2 || len(name) > 32 {
		return errors.New("invalid name length")
	}
	for _, item := range notSupportedItems {
		if strings.Contains(name, item) {
			return errors.New("name could not include @, #, =, *, ^")
		}
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 || len(password) > 32 {
		return errors.New("invalid password length")
	}
	return nil
}

func validateForRegister(name, password string) error {
	var err error
	err = validateName(name)
	if err != nil {
		return err
	}
	err = validatePassword(password)
	if err != nil {
		return err
	}
	return nil
}

func Register(ctx context.Context, c *app.RequestContext) {
	var body user.RegisterRequest
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}

	if err := validateName(body.Name); err != nil {
		c.JSON(http.StatusBadRequest, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	if err := validatePassword(body.Password); err != nil {
		c.JSON(http.StatusBadRequest, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}

	req := &user.RegisterRequest{
		Name:     body.Name,
		Password: body.Password,
	}
	resp, err := rpc.Register(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ConstructSuccessResponse(resp))
}

func Login(ctx context.Context, c *app.RequestContext) {
	var body user.LoginRequest
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}

	req := &user.LoginRequest{
		ID:       body.ID,
		Password: body.Password,
	}
	resp, err := rpc.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ConstructSuccessResponse(resp))
}
