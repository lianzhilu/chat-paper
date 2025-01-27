package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lianzhilu/chat-paper/cp-core/apps/gateway/rpc"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/errmsg"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/response"
	"net/http"
)

func CreateArticle(ctx context.Context, c *app.RequestContext) {
	var req article.CreateArticleRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	resp, err := rpc.CreateArticle(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ConstructErrorResponse(errmsg.ErrInvalidParameter, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ConstructSuccessResponse(resp))
}
