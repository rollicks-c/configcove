package profiles

import (
	"github.com/rollicks-c/configcove/internal/fsio"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

func (m Manager[T]) initProfiles() profileCollection[T] {

	// init empty profiles
	profiles := profileCollection[T]{
		Current:  "default",
		Profiles: map[string]Profile[T]{"default": *new(Profile[T])},
	}
	m.saveProfiles(profiles, true)

	return profiles
}

func (m Manager[T]) loadProfiles() profileCollection[T] {

	col := profileCollection[T]{
		Profiles: map[string]Profile[T]{},
	}

	// load file
	raw, err := os.ReadFile(m.getSettingsFilePath(false))
	if err != nil {
		return m.initProfiles()
	}

	// parse
	if err := yaml.Unmarshal(raw, &col); err != nil {
		panic(err)
	}

	return col
}
func (m Manager[T]) loadFile(filePath string) (T, error) {

	var profile T

	// load file
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return profile, err
	}

	// parse
	if err := yaml.Unmarshal(raw, &profile); err != nil {
		return profile, err
	}

	return profile, nil
}

func (m Manager[T]) saveProfiles(settings profileCollection[T], isFallback bool) {

	// ensure directories
	fsio.EnsureConfigDir(m.appName)
	filePath := m.getSettingsFilePath(isFallback)

	// save
	data, err := yaml.Marshal(settings)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(filePath, data, os.ModePerm); err != nil {
		panic(err)
	}

}

func (m Manager[T]) getSettingsFilePath(isFallback bool) string {
	if isFallback {
		return m.getConfigFilePath(settingsDefaultFile)
	}
	return m.getConfigFilePath(settingsFile)
}

func (m Manager[T]) getConfigFilePath(file string) string {
	configFile := filepath.Join(fsio.GetConfigDir(m.appName), file)
	return strings.ToLower(configFile)
}

func sanitizeExpression(exp string) string {
	exp = strings.ReplaceAll(exp, " ", "-")
	exp = strings.ToLower(exp)
	return exp
}
