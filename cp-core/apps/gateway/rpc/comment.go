package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/base"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/comment"
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

func CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (*comment.CreateCommentResponse, error) {
	return commentClient.CreateComment(ctx, req)
}

func GetComment(ctx context.Context, req *comment.GetCommentRequest) (*comment.CompletedComment, error) {
	return commentClient.GetComment(ctx, req)
}

func UpdateComment(ctx context.Context, req *comment.UpdateCommonRequest) (*base.EmptyBody, error) {
	return commentClient.UpdateComment(ctx, req)
}

func DeleteComment(ctx context.Context, req *comment.DeleteCommonRequest) (*base.EmptyBody, error) {
	return commentClient.DeleteComment(ctx, req)
}
