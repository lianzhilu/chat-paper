package service

import (
	"context"
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/constants"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/generator"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/repository"
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
