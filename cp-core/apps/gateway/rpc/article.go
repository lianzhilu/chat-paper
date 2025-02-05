package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article/articleservice"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/base"
	"time"
)

var (
	articleClient articleservice.Client
)

func InitArticleClient() {
	r, err := etcd.NewEtcdResolver(etcdAddrs)
	c, err := articleservice.NewClient(
		"article",
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	articleClient = c
}

func CreateArticle(ctx context.Context, req *article.CreateArticleRequest) (*article.CreateArticleResponse, error) {
	return articleClient.CreateArticle(ctx, req)
}

func GetArticle(ctx context.Context, req *article.GetArticleRequest) (*article.Article, error) {
	return articleClient.GetArticle(ctx, req)
}

func UpdateArticle(ctx context.Context, req *article.UpdateArticleRequest) (*base.EmptyBody, error) {
	return articleClient.UpdateArticle(ctx, req)
}

func ListArticles(ctx context.Context, req *article.ListArticlesRequest) (*article.ListArticlesResponse, error) {
	return articleClient.ListArticles(ctx, req)
}

func DeleteArticle(ctx context.Context, req *article.DeleteArticleRequest) (*base.EmptyBody, error) {
	return articleClient.DeleteArticle(ctx, req)
}
