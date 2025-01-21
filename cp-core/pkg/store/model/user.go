package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type User struct {
	ID          uint64                `gorm:"primaryKey;auto_increment"`
	SID         string                `gorm:"uniqueIndex:uniq_user_id;column:sid;type:varchar(255);not null;comment:用户ID"`
	Name        string                `gorm:"type:varchar(128);not null;comment:用户名"`
	PhoneNumber string                `gorm:"type:varchar(20);"`
	Email       string                `gorm:"type:varchar(40);"`
	Password    string                `gorm:"type:varchar(255);"`
	CreateTime  time.Time             `gorm:"autoCreateTime"`
	UpdateTime  time.Time             `gorm:"autoUpdateTime"`
	DeleteTime  soft_delete.DeletedAt `gorm:"uniqueIndex:uniq_user_id"`
}
