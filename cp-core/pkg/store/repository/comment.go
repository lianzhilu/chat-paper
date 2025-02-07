package repository

import (
	"context"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/dal"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/model"
	"gorm.io/gorm"
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
}

type MySQLCommentRepository struct {
	*gorm.DB
}

func NewMySQLCommentRepository(db *gorm.DB) *MySQLCommentRepository {
	dal.SetDefault(db)
	return &MySQLCommentRepository{db}
}

func (r *MySQLCommentRepository) CreateComment(ctx context.Context, param CreateCommentParams) error {
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
		err := tx.Comment.WithContext(ctx).Create(commentMeta)
		if err != nil {
			return err
		}

		// 2. create comment content
		commentContent := &model.CommentContents{
			CommentID: param.ID,
			Content:   param.Content,
		}

		err = tx.CommentContents.WithContext(ctx).Create(commentContent)
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
