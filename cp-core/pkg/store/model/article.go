package model

import (
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/article"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Article struct {
	ID           uint64                `gorm:"primaryKey;auto_increment"`
	SID          string                `gorm:"uniqueIndex:uniq_article_id;column:sid;type:varchar(255);not null;comment:文章ID"`
	AuthorID     string                `gorm:"type:varchar(128);not null;comment:文章作者"`
	Title        string                `gorm:"type:varchar(128);not null;comment:文章标题"`
	Status       article.ArticleStatus `gorm:"type:varchar(128);not null;comment:文章状态"`
	Content      string                `gorm:"type:text;not null;comment:文章内容"`
	ViewCount    uint64                `gorm:"type:int;not null;default:0;comment:文章阅读量"`
	LikeCount    uint64                `gorm:"type:int;not null;default:0;comment:文章点赞量"`
	CommentCount uint64                `gorm:"type:int;not null;default:0;comment:文章评论量"`
	CreateTime   time.Time             `gorm:"autoCreateTime"`
	UpdateTime   time.Time             `gorm:"autoUpdateTime"`
	DeleteTime   soft_delete.DeletedAt `gorm:"uniqueIndex:uniq_article_id"`
}
