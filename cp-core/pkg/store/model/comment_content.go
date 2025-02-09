package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// CommentContent record reply content
type CommentContent struct {
	ID         uint64                `gorm:"primaryKey;auto_increment"`
	CommentID  string                `gorm:"uniqueIndex:uniq_comment_id;column:sid;type:varchar(255);not null;comment:评论ID"`
	Content    string                `gorm:"type:varchar(255);not null;comment:评论内容"`
	CreateTime time.Time             `gorm:"autoCreateTime"`
	UpdateTime time.Time             `gorm:"autoUpdateTime"`
	DeleteTime soft_delete.DeletedAt `gorm:"uniqueIndex:uniq_comment_id"`
}
