package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article/articleservice"
	"time"
)

var (
	articleClient      articleservice.Client
	articleServiceAddr = fmt.Sprintf("%s:%s", serviceConfig.ArticleServiceConfig.ArticleServiceHost, serviceConfig.ArticleServiceConfig)
)

func InitArticleClient() {
	r, err := etcd.NewEtcdResolver(etcdAddrs)
	c, err := articleservice.NewClient(
		"article",
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithHostPorts(articleServiceAddr),
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
