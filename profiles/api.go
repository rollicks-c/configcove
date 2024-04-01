package profiles

import (
	"fmt"
	"sort"
)

type Option[T any] func(manager *Manager[T])

type Manager[T any] struct {
	appName        string
	defaultProfile T
}

func WithDefault[T any](profile T) Option[T] {
	return func(manager *Manager[T]) {
		manager.defaultProfile = profile
	}
}

func NewManager[T any](appName string, options ...Option[T]) *Manager[T] {
	m := &Manager[T]{
		appName: sanitizeExpression(appName),
	}
	for _, opt := range options {
		opt(m)
	}
	return m
}

func (m Manager[T]) LoadCurrent() T {
	profiles := m.loadProfiles()
	return m.Load(profiles.Current)
}

func (m Manager[T]) Load(name string) T {

	// load
	profiles := m.loadProfiles()
	profile, ok := profiles.getProfile(name)
	if !ok {
		panic(fmt.Sprintf("profile [%s] not found, ", name))
	}

	// enrich
	profile = m.inject(profile)

	return profile
}

func (m Manager[T]) Update(profile T) {
	profileList := m.loadProfiles()
	current := profileList.Current
	profileList.Profiles[current] = profile
	m.saveProfiles(profileList, false)
}

func (m Manager[T]) List() []string {
	profiles := m.loadProfiles()
	nameList := make([]string, 0, len(profiles.Profiles))
	for name := range profiles.Profiles {
		nameList = append(nameList, name)
	}
	sort.Slice(nameList, func(i, j int) bool {
		return nameList[i] < nameList[j]
	})
	return nameList
}

func (m Manager[T]) Switch(name string) error {

	// validate
	profileList := m.loadProfiles()
	if _, ok := profileList.Profiles[name]; !ok {
		return fmt.Errorf("unknown profile [%s]", name)
	}

	// activate
	profileList.Current = name
	m.saveProfiles(profileList, false)

	return nil
}
