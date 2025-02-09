package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lianzhilu/chat-paper/cp-core/apps/gateway/rpc"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/comment"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/errmsg"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/response"
	"net/http"
)

func CreateComment(ctx context.Context, c *app.RequestContext) {
	var req comment.CreateCommentRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	resp, err := rpc.CreateComment(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ConstructSuccessResponse(resp))
}

func GetComment(ctx context.Context, c *app.RequestContext) {
	var req comment.GetCommentRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	resp, err := rpc.GetComment(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ConstructSuccessResponse(resp))
}

func UpdateComment(ctx context.Context, c *app.RequestContext) {
	var req comment.UpdateCommonRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	resp, err := rpc.UpdateComment(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ConstructSuccessResponse(resp))
}

func DeleteComment(ctx context.Context, c *app.RequestContext) {
	var req comment.DeleteCommonRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	resp, err := rpc.DeleteComment(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ConstructSuccessResponse(resp))
}
