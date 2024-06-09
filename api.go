package configcove

import (
	"github.com/rollicks-c/configcove/internal/fsio"
	"github.com/rollicks-c/configcove/profiles"
	"github.com/rollicks-c/configcove/store"
	"os"
)

func Profiles[T any](appName string) *profiles.Manager[T] {
	m := profiles.NewManager[T](appName)
	return m
}

func Store(appName string) *store.Manager {
	return store.New(appName)
}

func HomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return homeDir
}

func ConfigDir(appName string) string {
	return fsio.GetConfigDir(appName)
}
