package service

import (
	"context"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/base"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/comment"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/constants"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/generator"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/repository"
)

type CommentServiceImpl struct {
	commentRepo repository.CommentRepository
}

func NewCommentService(commentRepo repository.CommentRepository) *CommentServiceImpl {
	return &CommentServiceImpl{commentRepo: commentRepo}
}

func (impl *CommentServiceImpl) CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (resp *comment.CreateCommentResponse, err error) {
	id, err := generator.GenerateSID(constants.CommentPrefix)
	if err != nil {
		return nil, err
	}
	param := repository.CreateCommentParams{
		ID:        id,
		AuthorID:  req.AuthorID,
		ArticleID: req.ArticleID,
		ParentID:  req.ParentID,
		Content:   req.Content,
	}
	err = impl.commentRepo.CreateComment(ctx, param)
	if err != nil {
		return nil, err
	}
	return &comment.CreateCommentResponse{
		ID: id,
	}, nil
}

func (impl *CommentServiceImpl) GetComment(ctx context.Context, req *comment.GetCommentRequest) (resp *comment.CompletedComment, err error) {
	ccm, err := impl.commentRepo.GetComment(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	resp = convertCompleteCommentDo2GetCommentResponse(ccm)
	return resp, nil
}

func (impl *CommentServiceImpl) UpdateComment(ctx context.Context, req *comment.UpdateCommonRequest) (resp *base.EmptyBody, err error) {
	param := repository.UpdateCommentParams{
		ID:      req.ID,
		Content: req.Content,
	}
	err = impl.commentRepo.UpdateComment(ctx, param)
	if err != nil {
		return nil, err
	}
	return &base.EmptyBody{}, nil
}

func (impl *CommentServiceImpl) DeleteComment(ctx context.Context, req *comment.DeleteCommonRequest) (resp *base.EmptyBody, err error) {
	err = impl.commentRepo.DeleteComment(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &base.EmptyBody{}, nil
}
