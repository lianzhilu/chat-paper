package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// Comment records reply meta
type Comment struct {
	ID           uint64                `gorm:"primaryKey;auto_increment"`
	SID          string                `gorm:"uniqueIndex:uniq_comment_id;column:sid;type:varchar(255);not null;comment:评论ID"`
	AuthorID     string                `gorm:"type:varchar(128);not null;comment:评论发出者"`
	ArticleID    string                `gorm:"type:varchar(128);not null;comment:评论文章id"`
	ParentID     string                `gorm:"type:varchar(128);not null;default:'0';comment:父评论id"`
	LikeCount    int64                 `gorm:"type:int;not null;default:0;comment:评论点赞量"`
	CommentCount int64                 `gorm:"type:int;not null;default:0;comment:评论回复量"`
	CreateTime   time.Time             `gorm:"autoCreateTime"`
	UpdateTime   time.Time             `gorm:"autoUpdateTime"`
	DeleteTime   soft_delete.DeletedAt `gorm:"uniqueIndex:uniq_comment_id"`
}
