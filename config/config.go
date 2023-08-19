package config

import (
	"encoding/json"
	"os"
)

type config struct {
	CurrentAnchor *string `json:"current_anchor"`
}

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

	// Save to config file
	return saveConfig(
		&config{
			CurrentAnchor: &path,
		},
	)
}

func Unanchor() error {
	err := os.Unsetenv("ANCHOR")
	if err != nil {
		return err
	}

	return saveConfig(
		&config{
			CurrentAnchor: nil,
		},
	)
}
