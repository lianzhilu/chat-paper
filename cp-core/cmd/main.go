package main

import (
	"fmt"
	"github.com/lianzhilu/chat-paper/cp-core/cmd/automigrate"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "cp-core",
		Short: "cp-core",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("please enter an option")
		},
	}
	cmd.AddCommand(automigrate.NewCommand())
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
