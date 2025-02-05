package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/base"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/constants"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/generator"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/repository"
	"gorm.io/gorm"
)

type ArticleServiceImpl struct {
	articleRepo repository.ArticleRepository
}

func NewArticleServiceImpl(articleRepo repository.ArticleRepository) *ArticleServiceImpl {
	return &ArticleServiceImpl{
		articleRepo: articleRepo,
	}
}

func (impl *ArticleServiceImpl) CreateArticle(ctx context.Context, req *article.CreateArticleRequest) (resp *article.CreateArticleResponse, err error) {
	id, err := generator.GenerateSID(constants.ArticlePrefix)
	if err != nil {
		return nil, err
	}
	param := &repository.CreateArticleParams{
		ID:       id,
		AuthorID: req.AuthorID,
		Title:    req.Title,
		Status:   req.Status,
		Content:  req.Content,
	}
	err = impl.articleRepo.CreateArticle(ctx, param)
	if err != nil {
		return nil, err
	}
	return &article.CreateArticleResponse{
		ID:     id,
		Status: req.Status,
	}, nil
}

func (impl *ArticleServiceImpl) GetArticle(ctx context.Context, req *article.GetArticleRequest) (resp *article.Article, err error) {
	art, err := impl.articleRepo.GetArticle(ctx, req.ID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, fmt.Errorf("%s not found", req.ID)
		}
		return nil, err
	}
	resp = &article.Article{
		ID:           art.SID,
		AuthorID:     art.AuthorID,
		Title:        art.Title,
		Content:      art.Content,
		Status:       art.Status,
		ViewCount:    int64(art.ViewCount),
		LikeCount:    int64(art.LikeCount),
		CommentCount: int64(art.CommentCount),
		CreateTime:   art.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:   art.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return resp, nil
}

func (impl *ArticleServiceImpl) UpdateArticle(ctx context.Context, req *article.UpdateArticleRequest) (resp *base.EmptyBody, err error) {
	param := &repository.UpdateArticleParams{
		ID:      req.ID,
		Title:   req.Title,
		Status:  req.Status,
		Content: req.Content,
	}
	err = impl.articleRepo.UpdateArticle(ctx, param)
	if err != nil {
		return nil, err
	}
	return &base.EmptyBody{}, nil
}

func (impl *ArticleServiceImpl) ListArticles(ctx context.Context, req *article.ListArticlesRequest) (resp *article.ListArticlesResponse, err error) {
	param := &repository.ListArticlesParams{
		PageNumber: int(req.GetPageNumber()),
		PageSize:   int(req.GetPageSize()),
		SortBy:     req.GetSortBy(),
		SortOrder:  req.GetSortOrder(),
		Statuses:   req.Statuses,
		AuthorIDs:  req.AuthorIDs,
	}
	arts, totalCount, err := impl.articleRepo.ListArticles(ctx, param)
	articleResps := convertArticleDo2ListArticleItems(arts)
	if err != nil {
		return nil, err
	}
	return &article.ListArticlesResponse{
		TotalCount: totalCount,
		PageNumber: req.GetPageNumber(),
		PageSize:   req.GetPageSize(),
		Articles:   articleResps,
	}, nil
}

func (impl *ArticleServiceImpl) DeleteArticle(ctx context.Context, req *article.DeleteArticleRequest) (resp *base.EmptyBody, err error) {
	err = impl.articleRepo.DeleteArticle(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &base.EmptyBody{}, nil
}
