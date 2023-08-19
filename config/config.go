package config

import "os"

func AnchorToPath(path string) error {
	err := os.Setenv("ANCHOR", path)
	if err != nil {
		return err
	}

	// Save to config file
	return nil
}
