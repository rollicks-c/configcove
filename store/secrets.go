package store

import (
	"fmt"
	"github.com/rollicks-c/configcove/internal/fsio"
	"path"
)

func (m *Manager) createSecretsFilePath(name string) string {
	dirPath := fsio.EnsureSecretsPath(m.appName)
	name = fmt.Sprintf(".%s", name)
	filePath := path.Join(dirPath, name)
	return filePath
}
