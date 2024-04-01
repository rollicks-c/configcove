package configcove

import (
	"github.com/rollicks-c/configcove/profiles"
	"github.com/rollicks-c/configcove/store"
)

func Profiles[T any](appName string) *profiles.Manager[T] {
	m := profiles.NewManager[T](appName)
	return m
}

func Store(appName string) *store.Manager {
	return store.New(appName)
}
