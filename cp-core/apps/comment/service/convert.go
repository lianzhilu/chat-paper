package service

import (
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/comment"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/repository"
	"time"
)

func convertCompleteCommentDo2GetCommentResponse(ccm *repository.CompletedCommentModel) *comment.CompletedComment {
	return &comment.CompletedComment{
		ID:           ccm.CommentMeta.SID,
		AuthorID:     ccm.CommentMeta.AuthorID,
		ArticleID:    ccm.CommentMeta.ArticleID,
		ParentID:     ccm.CommentMeta.ParentID,
		Content:      ccm.CommentContent,
		LikeCount:    ccm.CommentMeta.LikeCount,
		CommentCount: ccm.CommentMeta.CommentCount,
		CreateTime:   ccm.CommentMeta.CreateTime.Format(time.RFC3339),
		UpdateTime:   ccm.CommentMeta.UpdateTime.Format(time.RFC3339),
	}
}
