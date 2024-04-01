package fsio

import (
	"github.com/kirsle/configdir"
	"os"
	"path"
	"strings"
)

func GetConfigDir(appName string) string {
	dirPath := configdir.LocalConfig(appName)
	return strings.ToLower(dirPath)
}

func EnsureConfigDir(appName string) string {
	dirPath := GetConfigDir(appName)
	err := configdir.MakePath(dirPath)
	if err != nil {
		panic(err)
	}
	return dirPath
}

func EnsureSecretsPath(appName string) string {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	secretsDir := path.Join(homeDir, ".secrets", appName)

	if err := os.MkdirAll(secretsDir, os.FileMode(0700)); err != nil {
		panic(err)
	}

	return secretsDir
}
