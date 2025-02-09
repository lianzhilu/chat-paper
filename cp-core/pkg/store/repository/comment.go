package repository

import (
	"context"
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

type UpdateCommentParams struct {
	ID      string
	Content string
}

type CompletedCommentModel struct {
	CommentMeta    model.Comment
	CommentContent string
}

type CommentRepository interface {
	CreateComment(ctx context.Context, param CreateCommentParams) error
	GetComment(ctx context.Context, id string) (*CompletedCommentModel, error)
	UpdateComment(ctx context.Context, param UpdateCommentParams) error
	DeleteComment(ctx context.Context, id string) error
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
		commentContent := &model.CommentContent{
			CommentID: param.ID,
			Content:   param.Content,
		}

		err = tx.CommentContent.WithContext(tCtx).Create(commentContent)
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

func (r *MySQLCommentRepository) GetComment(ctx context.Context, id string) (*CompletedCommentModel, error) {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var completedComment *CompletedCommentModel

	commentOrm := dal.Comment
	commentContentOrm := dal.CommentContent
	query := dal.Use(r.DB)
	err := query.Transaction(func(tx *dal.Query) error {
		commentMeta, err := tx.Comment.WithContext(tCtx).Where(commentOrm.SID.Eq(id)).First()
		if err != nil {
			return err
		}
		commentContent, err := tx.CommentContent.WithContext(tCtx).Where(commentContentOrm.CommentID.Eq(id)).First()
		if err != nil {
			return err
		}
		completedComment = &CompletedCommentModel{
			CommentMeta:    *commentMeta,
			CommentContent: commentContent.Content,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return completedComment, nil
}

func (r *MySQLCommentRepository) UpdateComment(ctx context.Context, param UpdateCommentParams) error {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	commentContentOrm := dal.CommentContent
	query := commentContentOrm.WithContext(tCtx)
	_, err := query.Where(commentContentOrm.CommentID.Eq(param.ID)).Update(commentContentOrm.Content, param.Content)
	if err != nil {
		return err
	}
	return nil

}

func (r *MySQLCommentRepository) DeleteComment(ctx context.Context, id string) error {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	commentOrm := dal.Comment
	commentContentOrm := dal.CommentContent
	query := dal.Use(r.DB)
	err := query.Transaction(func(tx *dal.Query) error {
		_, err := tx.Comment.WithContext(tCtx).Where(commentOrm.SID.Eq(id)).Delete()
		if err != nil {
			return err
		}
		_, err = tx.CommentContent.WithContext(tCtx).Where(commentContentOrm.CommentID.Eq(id)).Delete()
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
