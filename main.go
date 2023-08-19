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
				path, err := config.GetSavedAnchorPath(args[0])
				if err != nil {
					fmt.Println("Error getting saved anchor:", err)
					return
				}

				if path == "" {
					fmt.Println("No saved anchor named '" + args[0] + "'")
					return
				}

				err = config.AnchorToPath(path)
				if err != nil {
					fmt.Println("Error setting anchor:", err)
					return
				}

				fmt.Println("⚓️ Anchored to", path)
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

				fmt.Println("⚓️ Anchored to", currentDir)
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

			fmt.Println("⛵️ Anchor lifted")
		},
	}

	var cmdSave = &cobra.Command{
		Use:   "save [anchor_name]",
		Short: "Saves the current directory as anchor_name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			currentDir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory:", err)
				return
			}

			err = config.SaveAnchor(args[0], currentDir)
			if err != nil {
				fmt.Println("Error saving anchor:", err)
				return
			}

			fmt.Println("⚓️ Anchor '" + args[0] + "' stashed at " + currentDir + ". Drop it with 'anchor down " + args[0] + "'.")
		},
	}

	var cmdRemove = &cobra.Command{
		Use:   "remove [anchor_name]",
		Short: "Deletes the saved anchor_name directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := config.RemoveAnchor(args[0])
			if err != nil {
				fmt.Println("Error removing anchor:", err)
				return
			}

			fmt.Println("⚓️ Anchor '" + args[0] + "' removed.")
		},
	}

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "List current saved directories",
		Run: func(cmd *cobra.Command, args []string) {
			savedAnchors, err := config.ListSavedAnchors()
			if err != nil {
				fmt.Println("Error listing saved anchors:", err)
				return
			}

			for anchorName, anchorPath := range savedAnchors {
				fmt.Println(anchorName, ":", anchorPath)
			}
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
