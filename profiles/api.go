package profiles

import (
	"fmt"
	"sort"
)

type Profile[T any] struct {
	Name string `yaml:"name"`
	Data T      `yaml:"data"`
}

type Option[T any] func(manager *Manager[T])

type Manager[T any] struct {
	appName        string
	defaultProfile T
}

func WithDefault[T any](profileData T) Option[T] {
	return func(manager *Manager[T]) {
		manager.defaultProfile = profileData
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

func (m Manager[T]) LoadFile(filePath string) (T, error) {
	return m.loadFile(filePath)
}

func (m Manager[T]) LoadCurrent() Profile[T] {
	profiles := m.loadProfiles()
	return m.Load(profiles.Current, false)
}

func (m Manager[T]) LoadCurrentRaw() Profile[T] {
	profiles := m.loadProfiles()
	return m.Load(profiles.Current, true)
}

func (m Manager[T]) Load(name string, loadRaw bool) Profile[T] {

	// load
	profiles := m.loadProfiles()
	profile, ok := profiles.getProfile(name)
	if !ok {
		panic(fmt.Sprintf("profile [%s] not found, ", name))
	}

	// enrich
	profile.Name = name
	if !loadRaw {
		profile.Data = m.inject(profile.Data)
	}

	return profile
}

func (m Manager[T]) Update(profile Profile[T]) {
	profileList := m.loadProfiles()
	profileList.Profiles[profile.Name] = profile
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
