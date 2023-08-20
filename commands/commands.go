package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"anchor/config"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

func CreateDownCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "down [anchor_name]",
		Short:             "Sets the current directory as the default directory",
		Args:              cobra.MaximumNArgs(1),
		Run:               downCommandRun,
		ValidArgsFunction: validArgsFunction,
	}
}

func CreateUpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Unsets the default directory",
		Run:   upCommandRun,
	}
}

func CreateGoCommand() *cobra.Command {
	cmdGo := &cobra.Command{
		Use:               "go [anchor_name]",
		Short:             "Navigates to the specified anchor or default anchor if none is given",
		Args:              cobra.MaximumNArgs(1),
		Run:               goCommandRun,
		ValidArgsFunction: validArgsFunction,
	}
	cmdGo.Flags().BoolP("fuzzy", "f", false, "Enable fuzzy finding mode")
	return cmdGo
}

func CreateSaveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "save [anchor_name]",
		Short: "Saves the current directory as anchor_name",
		Args:  cobra.ExactArgs(1),
		Run:   saveCommandRun,
	}
}

func CreateRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:               "remove [anchor_name]",
		Short:             "Deletes the saved anchor_name directory",
		Args:              cobra.ExactArgs(1),
		Run:               removeCommandRun,
		ValidArgsFunction: validArgsFunction,
	}
}

func CreateListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List current saved directories",
		Run:   listCommandRun,
	}
}

func CreateGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get the path of the current anchor",
		Run:   getCommandRun,
	}
}

func CreateCompletionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:
Bash:

$ source <(yourprogram completion bash)


Zsh:

$ source <(yourprogram completion zsh)

Fish:

$ yourprogram completion fish | source
`,
		Args: cobra.ExactValidArgs(1),
		Run:  completionCommandRun,
	}
}

// Utility functions
func validArgsFunction(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	savedAnchorNames, err := config.GetSavedAnchorNames()
	if err != nil {
		fmt.Println("Error getting saved anchors:", err)
		return nil, cobra.ShellCompDirectiveError
	}
	return savedAnchorNames, cobra.ShellCompDirectiveNoFileComp
}

func downCommandRun(cmd *cobra.Command, args []string) {
	// Implement the function as in the original code.
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
}

func upCommandRun(cmd *cobra.Command, args []string) {
	err := config.Unanchor()
	if err != nil {
		fmt.Println("Error unsetting anchor:", err)
		return
	}

	fmt.Println("⛵️ Anchor lifted")
}

func goCommandRun(cmd *cobra.Command, args []string) {
	fuzzyFlag, err := cmd.Flags().GetBool("fuzzy")
	if err != nil {
		fmt.Println("Error getting fuzzy flag:", err)
		return
	}

	if fuzzyFlag {
		savedAnchors, err := config.ListSavedAnchors()
		if err != nil {
			fmt.Println("Error getting saved anchors:", err)
			return
		}

		anchorNames := make([]string, 0, len(savedAnchors))
		for anchorName := range savedAnchors {
			anchorNames = append(anchorNames, anchorName)
		}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}

		idx, err := fuzzyfinder.Find(
			anchorNames,
			func(i int) string {
				abbreviatedPath := strings.Replace(savedAnchors[anchorNames[i]], homeDir, "~", 1)
				return anchorNames[i] + " ⚓️ " + abbreviatedPath
			},
		)
		if err != nil {
			fmt.Println("Error selecting anchor:", err)
			return
		}

		path := savedAnchors[anchorNames[idx]]
		fmt.Println(path)
	} else {
		if len(args) == 1 {
			// Get the path of the named anchor
			path, err := config.GetSavedAnchorPath(args[0])
			if err != nil {
				fmt.Println("Error getting saved anchor:", err)
				return
			}

			if path == "" {
				return
			}

			fmt.Println(path)
		} else {
			// Get the path of the default anchor
			path, err := config.GetDefaultAnchor()
			if err != nil {
				fmt.Println("Error getting default anchor:", err)
				return
			}

			if path == "" {
				return
			}

			fmt.Println(path)
		}
	}
}

func saveCommandRun(cmd *cobra.Command, args []string) {
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

	fmt.Println("📍 Anchor '" + args[0] + "' stashed at " + currentDir + ". Drop it with 'anchor down " + args[0] + "'.")
}

func removeCommandRun(cmd *cobra.Command, args []string) {
	err := config.RemoveAnchor(args[0])
	if err != nil {
		fmt.Println("Error removing anchor:", err)
		return
	}

	fmt.Println("⚓️ Anchor '" + args[0] + "' removed.")
}

func listCommandRun(cmd *cobra.Command, args []string) {
	savedAnchors, err := config.ListSavedAnchors()
	if err != nil {
		fmt.Println("Error listing saved anchors:", err)
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	anchorNames := make([]string, 0, len(savedAnchors))
	for anchorName := range savedAnchors {
		anchorNames = append(anchorNames, anchorName)
	}

	idx, err := fuzzyfinder.Find(
		anchorNames,
		func(i int) string {
			// Abbreviate home path as ~/
			abbreviatedPath := strings.Replace(savedAnchors[anchorNames[i]], homeDir, "~", 1)
			return anchorNames[i] + ": " + abbreviatedPath
		},
		fuzzyfinder.WithPromptString("⚓️ > "),
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			path := savedAnchors[anchorNames[i]]

			files, err := ioutil.ReadDir(path)
			if err != nil {
				return "Error reading directory"
			}

			var directories, filesList []os.FileInfo
			for _, file := range files {
				if file.IsDir() {
					directories = append(directories, file)
				} else {
					filesList = append(filesList, file)
				}
			}

			// Sort alphabetically
			sort.Slice(directories, func(i, j int) bool {
				return directories[i].Name() < directories[j].Name()
			})
			sort.Slice(filesList, func(i, j int) bool {
				return filesList[i].Name() < filesList[j].Name()
			})

			// Build the preview string with directories first, then files
			var preview strings.Builder
			for _, dir := range directories {
				preview.WriteString("📁  " + dir.Name() + "/\n")
			}
			for _, file := range filesList {
				preview.WriteString("📄  " + file.Name() + "\n")
			}

			return preview.String()
		}),
	)
	if err != nil {
		fmt.Println("Error listing saved anchors:", err)
		return
	}

	abbreviatedPath := strings.Replace(savedAnchors[anchorNames[idx]], homeDir, "~", 1)
	fmt.Println("⚓", anchorNames[idx]+":", abbreviatedPath)
}

func getCommandRun(cmd *cobra.Command, args []string) {
	err := config.PrintAnchor()
	if err != nil {
		fmt.Println("Error getting current anchor:", err)
		return
	}
}

func completionCommandRun(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "bash":
		cmd.Root().GenBashCompletion(os.Stdout)
	case "zsh":
		cmd.Root().GenZshCompletion(os.Stdout)
	case "fish":
		cmd.Root().GenFishCompletion(os.Stdout, true)
	case "powershell":
		cmd.Root().GenPowerShellCompletion(os.Stdout)
	}
}
