package database

import (
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/model"
	"gorm.io/gorm"
)

var models = []interface{}{
	model.User{},
	model.Article{},
	model.Comment{},
	model.CommentContent{},
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}
