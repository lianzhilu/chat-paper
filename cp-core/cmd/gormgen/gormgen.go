package gormgen

import (
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/pkg/store/database"
	"github.com/spf13/cobra"
)

var OutPath = ""

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "gormgen",
		Run: func(cmd *cobra.Command, args []string) {
			if OutPath == "" {
				fmt.Println("please enter an usable output path")
				return
			}
			fmt.Println(OutPath)
			database.GenerateGorm(OutPath)
		},
	}
	cmd.Flags().StringVarP(&OutPath, "out", "o", "", "define gorm-gen output path")
	return cmd
}
