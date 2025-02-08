package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/comment/commentservice"
	"time"
)

var (
	commentClient commentservice.Client
)

func InitCommentClient() {
	r, err := etcd.NewEtcdResolver(etcdAddrs)
	c, err := commentservice.NewClient(
		"comment",
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}
