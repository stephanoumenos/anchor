package main

import (
	"anchor/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var cmdDown = &cobra.Command{
		Use:   "down [anchor_name]",
		Short: "Sets the current directory as the default directory",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				// Implementation for setting named anchor as default directory
			} else {
				// Implementation for setting current directory as default directory
				currentDir, err := os.Getwd()
				if err != nil {
					fmt.Println("Error getting current directory:", err)
					return
				}
				err = config.AnchorToPath(currentDir)
				if err != nil {
					fmt.Println("Error setting anchor:", err)
					return
				}
			}
		},
	}

	var cmdUp = &cobra.Command{
		Use:   "up",
		Short: "Unsets the default directory",
		Run: func(cmd *cobra.Command, args []string) {
			err := config.Unanchor()
			if err != nil {
				fmt.Println("Error unsetting anchor:", err)
				return
			}
		},
	}

	var cmdSave = &cobra.Command{
		Use:   "save [anchor_name]",
		Short: "Saves the current directory as anchor_name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for saving the current directory
		},
	}

	var cmdRemove = &cobra.Command{
		Use:   "remove [anchor_name]",
		Short: "Deletes the saved anchor_name directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for removing the saved anchor_name directory
		},
	}

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "List current saved directories",
		Run: func(cmd *cobra.Command, args []string) {
			// Implementation for listing saved directories
		},
	}

	var cmdGet = &cobra.Command{
		Use:   "get",
		Short: "Get the path of the current anchor",
		Run: func(cmd *cobra.Command, args []string) {
			err := config.PrintAnchor()
			if err != nil {
				fmt.Println("Error getting current anchor:", err)
				return
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "anchor"}
	rootCmd.AddCommand(cmdDown, cmdUp, cmdSave, cmdRemove, cmdList, cmdGet)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
