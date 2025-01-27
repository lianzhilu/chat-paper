package repository

import (
	"context"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/constants"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/dal"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/model"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"time"
)

type CreateArticleParams struct {
	ID       string
	AuthorID string
	Title    string
	Status   string
	Content  string
}

type UpdateArticleParams struct {
	ID      string
	Title   *string
	Status  *string
	Content *string
}

type ListArticlesParams struct {
	PageNumber int
	PageSize   int
	SortBy     string
	SortOrder  string
	Statuses   []article.ArticleStatus
	AuthorIDs  []string
}

type ArticleRepository interface {
	CreateArticle(ctx context.Context, param *CreateArticleParams) error
	GetArticle(ctx context.Context, id string) (*model.Article, error)
	UpdateArticle(ctx context.Context, param *UpdateArticleParams) error
	DeleteArticle(ctx context.Context, id string) error
}

type MySQLArticleRepository struct {
	*gorm.DB
}

func NewMySQLArticleRepository(db *gorm.DB) *MySQLArticleRepository {
	dal.SetDefault(db)
	return &MySQLArticleRepository{db}
}

func (r *MySQLArticleRepository) CreateArticle(ctx context.Context, param *CreateArticleParams) error {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := dal.Article.WithContext(tCtx)
	art := &model.Article{
		SID:      param.ID,
		AuthorID: param.AuthorID,
		Title:    param.Title,
		Status:   param.Status,
		Content:  param.Content,
	}
	if err := query.Create(art); err != nil {
		return err
	}
	return nil
}

func (r *MySQLArticleRepository) GetArticle(ctx context.Context, id string) (*model.Article, error) {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	articleOrm := dal.Article
	query := articleOrm.WithContext(tCtx)
	art, err := query.Where(articleOrm.SID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return art, nil
}

func (r *MySQLArticleRepository) UpdateArticle(ctx context.Context, param *UpdateArticleParams) error {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	articleOrm := dal.Article
	query := articleOrm.WithContext(tCtx)

	var updateExprs []field.AssignExpr
	if param.Title != nil {
		updateExprs = append(updateExprs, articleOrm.Title.Value(*param.Title))
	}
	if param.Status != nil {
		updateExprs = append(updateExprs, articleOrm.Status.Value(*param.Status))
	}
	if param.Content != nil {
		updateExprs = append(updateExprs, articleOrm.Content.Value(*param.Content))
	}
	_, err := query.Where(articleOrm.SID.Eq(param.ID)).UpdateSimple(updateExprs...)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLArticleRepository) DeleteArticle(ctx context.Context, id string) error {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	articleOrm := dal.Article
	query := articleOrm.WithContext(tCtx)

	_, err := query.Where(articleOrm.SID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLArticleRepository) ListArticles(ctx context.Context, param *ListArticlesParams) ([]*model.Article, error) {
	tCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	articleOrm := dal.Article
	query := articleOrm.WithContext(tCtx)

	if len(param.Statuses) > 0 {
		query = query.Where(articleOrm.Status.In(param.Statuses...))
	}

	if len(param.AuthorIDs) > 0 {
		query = query.Where(articleOrm.AuthorID.In(param.AuthorIDs...))
	}

	offset := (param.PageNumber - 1) * param.PageSize
	if offset > 0 {
		query = query.Offset(offset)
	}
	if param.PageSize > 0 {
		query = query.Limit(param.PageSize)
	}

	sortByField, ok := articleOrm.GetFieldByName(param.SortBy)
	if !ok {
		sortByField = articleOrm.CreateTime
	}
	if param.SortOrder == constants.SortOrderAsc {
		query.Order(sortByField)
	} else {
		query.Order(sortByField.Desc())
	}

	articles, err := query.Find()
	if err != nil {
		return nil, err
	}
	return articles, nil
}
