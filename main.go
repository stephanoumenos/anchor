package main

import (
	"fmt"
	"os"

	"anchor/commands"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "anchor"}
	rootCmd.AddCommand(
		commands.CreateDownCommand(),
		commands.CreateUpCommand(),
		commands.CreateGoCommand(),
		commands.CreateSaveCommand(),
		commands.CreateRemoveCommand(),
		commands.CreateListCommand(),
		commands.CreateGetCommand(),
		commands.CreateCompletionCommand(),
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
