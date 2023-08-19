package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	CurrentAnchor *string      `json:"current_anchor"`
	SavedAnchors  savedAnchors `json:"saved_anchors"`
}

type savedAnchors map[string]string

// Just to avoid calling filepath.Dir to get the directory
type configPath struct {
	path string
	dir  string
}

func getConfigPath() (*configPath, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	return &configPath{
		path: home + "/.config/anchor/config.json",
		dir:  home + "/.config/anchor",
	}, nil
}

func saveConfig(config *config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(configPath.dir, 0755)
	if err != nil {
		return err
	}

	marshalledConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath.path, marshalledConfig, 0644)
	return err
}

func AnchorToPath(path string) error {
	err := os.Setenv("ANCHOR", path)
	if err != nil {
		return err
	}

	config, err := readConfig()
	if err != nil {
		return err
	}

	config.CurrentAnchor = &path

	err = saveConfig(config)
	return err
}

func Unanchor() error {
	err := os.Unsetenv("ANCHOR")
	if err != nil {
		return err
	}

	config, err := readConfig()
	if err != nil {
		return err
	}

	config.CurrentAnchor = nil

	err = saveConfig(config)
	return err
}

func readConfig() (*config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(configPath.path)
	if err != nil {
		return nil, err
	}

	var config config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func PrintAnchor() error {
	anchor := os.Getenv("ANCHOR")
	if anchor != "" {
		fmt.Println(anchor) // Print the anchor path for the shell to use
		return nil
	}

	config, err := readConfig()
	if err != nil {
		return err
	}

	if config.CurrentAnchor != nil {
		fmt.Println(*config.CurrentAnchor) // Print the current anchor path for the shell to use
		return nil
	}

	// If no anchor is set, do nothing.
	return nil
}

func SaveAnchor(anchorName, currentDir string) error {
	config, err := readConfig()
	if err != nil {
		return err
	}

	// Save the new anchor
	if config.SavedAnchors == nil {
		config.SavedAnchors = make(map[string]string)
	}
	config.SavedAnchors[anchorName] = currentDir

	// Save the updated config
	err = saveConfig(config)
	return err
}

func RemoveAnchor(anchorName string) error {
	config, err := readConfig()
	if err != nil {
		return err
	}

	// Remove the anchor
	delete(config.SavedAnchors, anchorName)

	err = saveConfig(config)
	return err
}

func ListSavedAnchors() (savedAnchors, error) {
	config, err := readConfig()
	if err != nil {
		return nil, err
	}

	return config.SavedAnchors, nil
}
