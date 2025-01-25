package database

import (
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/model"
	"gorm.io/gen"
)

func GenerateGorm(outPath string) {
	g := gen.NewGenerator(gen.Config{
		OutPath: outPath,
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.ApplyBasic(
		model.User{},
		model.Article{},
	)
	g.Execute()
}
