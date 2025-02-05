package service

import (
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/model"
	"github.com/samber/lo"
	"time"
)

func convertArticleDo2GetArticleResponse(art *model.Article) *article.Article {
	return &article.Article{
		ID:           art.SID,
		AuthorID:     art.AuthorID,
		Title:        art.Title,
		Content:      art.Content,
		Status:       art.Status,
		ViewCount:    int64(art.ViewCount),
		LikeCount:    int64(art.LikeCount),
		CommentCount: int64(art.CommentCount),
		CreateTime:   art.CreateTime.Format(time.RFC3339),
		UpdateTime:   art.UpdateTime.Format(time.RFC3339),
	}
}

func convertArticleDo2ListArticleItems(articles []*model.Article) []*article.Article {
	return lo.Map(articles, func(art *model.Article, _ int) *article.Article {
		return convertArticleDo2GetArticleResponse(art)
	})
}
