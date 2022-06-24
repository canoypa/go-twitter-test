package core

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetConfigPath() string {
	path := os.Getenv("HOME")

	if path == "" && runtime.GOOS == "windows" {
		path = os.Getenv("APPDATA")
	} else {
		path = filepath.Join(path, ".config")
	}

	path = filepath.Join(path, "twcli")

	return path
}
