package store

import (
	"fmt"
	"github.com/rollicks-c/configcove/internal/fsio"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

type Manager struct {
	appName string
}

func New(appName string) *Manager {
	return &Manager{
		appName: appName,
	}
}

func (m *Manager) SaveNumber(name string, value int) {

	// gather data
	dir := fsio.EnsureConfigDir(m.appName)
	filePath := filepath.Join(dir, name)

	// save
	data := []byte(strconv.Itoa(value))
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		panic(err)
	}

	return
}

func (m *Manager) LoadNumber(name string) (int, bool) {

	// gather data
	dir := fsio.EnsureConfigDir(m.appName)
	filePath := path.Join(dir, name)

	// load and parse
	valueRaw, err := os.ReadFile(filePath)
	if err != nil {
		return 0, false
	}
	value, err := strconv.Atoi(string(valueRaw))
	if err != nil {
		return 0, false
	}

	return value, true

}

func (m *Manager) LoadSecret(name string) (string, bool) {
	filePath := m.createSecretsFilePath(name)
	valueRaw, err := os.ReadFile(filePath)
	if err != nil {
		return "", false
	}
	return string(valueRaw), true
}

func (m *Manager) SaveSecret(name, value string) {
	filePath := m.createSecretsFilePath(name)
	if err := os.WriteFile(filePath, []byte(value), 0400); err != nil {
		panic(fmt.Errorf("failed to write file: %v", err))
	}
}
