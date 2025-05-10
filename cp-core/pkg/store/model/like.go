package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Like struct {
	ID           uint64                `gorm:"primaryKey;auto_increment"`
	UserID       string                `gorm:"type:varchar(128);not null;comment:用户ID"`
	ResourceType string                `gorm:"type:varchar(128);not null;comment:资源类型"`
	ResourceID   string                `gorm:"type:varchar(128);not null;comment:资源ID"`
	CreateTime   time.Time             `gorm:"autoCreateTime"`
	UpdateTime   time.Time             `gorm:"autoUpdateTime"`
	DeleteTime   soft_delete.DeletedAt `gorm:"uniqueIndex:uniq_comment_id"`
}
