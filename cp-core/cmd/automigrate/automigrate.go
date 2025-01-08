package automigrate

import (
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/database"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "automigrate",
		Run: func(cmd *cobra.Command, args []string) {
			db := database.GetDB()
			err := database.AutoMigrate(db)
			if err != nil {
				panic(err)
			}
		},
	}
	return cmd
}
