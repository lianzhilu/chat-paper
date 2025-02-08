// Code generated by Kitex v0.12.1. DO NOT EDIT.

package commentservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	base "github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/base"
	comment "github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/comment"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateComment": kitex.NewMethodInfo(
		createCommentHandler,
		newCommentServiceCreateCommentArgs,
		newCommentServiceCreateCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetComment": kitex.NewMethodInfo(
		getCommentHandler,
		newCommentServiceGetCommentArgs,
		newCommentServiceGetCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateComment": kitex.NewMethodInfo(
		updateCommentHandler,
		newCommentServiceUpdateCommentArgs,
		newCommentServiceUpdateCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteComment": kitex.NewMethodInfo(
		deleteCommentHandler,
		newCommentServiceDeleteCommentArgs,
		newCommentServiceDeleteCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	commentServiceServiceInfo                = NewServiceInfo()
	commentServiceServiceInfoForClient       = NewServiceInfoForClient()
	commentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return commentServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return commentServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func createCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCreateCommentArgs)
	realResult := result.(*comment.CommentServiceCreateCommentResult)
	success, err := handler.(comment.CommentService).CreateComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCreateCommentArgs() interface{} {
	return comment.NewCommentServiceCreateCommentArgs()
}

func newCommentServiceCreateCommentResult() interface{} {
	return comment.NewCommentServiceCreateCommentResult()
}

func getCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceGetCommentArgs)
	realResult := result.(*comment.CommentServiceGetCommentResult)
	success, err := handler.(comment.CommentService).GetComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceGetCommentArgs() interface{} {
	return comment.NewCommentServiceGetCommentArgs()
}

func newCommentServiceGetCommentResult() interface{} {
	return comment.NewCommentServiceGetCommentResult()
}

func updateCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceUpdateCommentArgs)
	realResult := result.(*comment.CommentServiceUpdateCommentResult)
	success, err := handler.(comment.CommentService).UpdateComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceUpdateCommentArgs() interface{} {
	return comment.NewCommentServiceUpdateCommentArgs()
}

func newCommentServiceUpdateCommentResult() interface{} {
	return comment.NewCommentServiceUpdateCommentResult()
}

func deleteCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceDeleteCommentArgs)
	realResult := result.(*comment.CommentServiceDeleteCommentResult)
	success, err := handler.(comment.CommentService).DeleteComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceDeleteCommentArgs() interface{} {
	return comment.NewCommentServiceDeleteCommentArgs()
}

func newCommentServiceDeleteCommentResult() interface{} {
	return comment.NewCommentServiceDeleteCommentResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (r *comment.CreateCommentResponse, err error) {
	var _args comment.CommentServiceCreateCommentArgs
	_args.Req = req
	var _result comment.CommentServiceCreateCommentResult
	if err = p.c.Call(ctx, "CreateComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetComment(ctx context.Context, req *comment.GetCommentResponse) (r *comment.CompletedComment, err error) {
	var _args comment.CommentServiceGetCommentArgs
	_args.Req = req
	var _result comment.CommentServiceGetCommentResult
	if err = p.c.Call(ctx, "GetComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateComment(ctx context.Context, req *comment.UpdateCommonResponse) (r *base.EmptyBody, err error) {
	var _args comment.CommentServiceUpdateCommentArgs
	_args.Req = req
	var _result comment.CommentServiceUpdateCommentResult
	if err = p.c.Call(ctx, "UpdateComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteComment(ctx context.Context, req *comment.DeleteCommonResponse) (r *base.EmptyBody, err error) {
	var _args comment.CommentServiceDeleteCommentArgs
	_args.Req = req
	var _result comment.CommentServiceDeleteCommentResult
	if err = p.c.Call(ctx, "DeleteComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
