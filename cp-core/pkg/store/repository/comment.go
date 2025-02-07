package repository

import (
	"context"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/comment"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/dal"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/model"
	"gorm.io/gorm"
	"time"
)

type CreateCommentParams struct {
	ID        string
	AuthorID  string
	ArticleID string
	ParentID  *string
	Content   string
}

type CommentRepository interface {
	CreateComment(ctx context.Context, param CreateCommentParams) error
	GetComment(ctx context.Context, id string) (*comment.CompletedComment, error)
}

type MySQLCommentRepository struct {
	*gorm.DB
}

func NewMySQLCommentRepository(db *gorm.DB) *MySQLCommentRepository {
	dal.SetDefault(db)
	return &MySQLCommentRepository{db}
}

func (r *MySQLCommentRepository) CreateComment(ctx context.Context, param CreateCommentParams) error {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := dal.Use(r.DB)
	err := query.Transaction(func(tx *dal.Query) error {
		// 1. create comment meta
		commentMeta := &model.Comment{
			SID:       param.ID,
			AuthorID:  param.AuthorID,
			ArticleID: param.ArticleID,
			ParentID:  "0",
		}
		if param.ParentID != nil {
			commentMeta.ParentID = *param.ParentID
		}
		err := tx.Comment.WithContext(tCtx).Create(commentMeta)
		if err != nil {
			return err
		}

		// 2. create comment content
		commentContent := &model.CommentContents{
			CommentID: param.ID,
			Content:   param.Content,
		}

		err = tx.CommentContents.WithContext(tCtx).Create(commentContent)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLCommentRepository) GetComment(ctx context.Context, id string) (*comment.CompletedComment, error) {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var completedComment *comment.CompletedComment

	commentOrm := dal.Comment
	commentContentOrm := dal.CommentContents
	query := dal.Use(r.DB)
	err := query.Transaction(func(tx *dal.Query) error {
		commentMeta, err := tx.Comment.WithContext(tCtx).Where(commentOrm.SID.Eq(id)).First()
		if err != nil {
			return err
		}
		commentContent, err := tx.CommentContents.WithContext(tCtx).Where(commentContentOrm.CommentID.Eq(id)).First()
		if err != nil {
			return err
		}
		completedComment = &comment.CompletedComment{
			ID:        commentMeta.SID,
			AuthorID:  commentMeta.AuthorID,
			ArticleID: commentMeta.ArticleID,
			ParentID:  commentMeta.ParentID,
			Content:   commentContent.Content,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return completedComment, nil
}
