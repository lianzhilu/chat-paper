package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article/articleservice"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/config"
	"time"
)

var (
	articleClient articleservice.Client
	serviceConfig = config.GetRuntimeConfig()
	serviceAddr   = fmt.Sprintf("%s:%s", serviceConfig.ArticleServiceConfig.ArticleServiceHost, serviceConfig.ArticleServiceConfig)
)

func InitArticleClient() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	c, err := articleservice.NewClient(
		"article",
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithHostPorts(serviceAddr),
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
