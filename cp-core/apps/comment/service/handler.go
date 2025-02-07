package service

import (
	"context"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/comment"
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
	id, err := generator.GenerateSID("cmt")
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
