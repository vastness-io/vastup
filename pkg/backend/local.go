package backend

import (
	"github.com/vastness-io/vastup/pkg/up"
)

const (
	repoBinPath = "bin/linux/amd64"
)

func NewLocalBackend(workDir string) *LocalBackend {
	return &LocalBackend{}
}

// LocalBackend is the default backend used for storing context and component
// configuration for vastup.
type LocalBackend struct {
	config *up.Config
}

// Read will return the current context read from the backend.
func (l *LocalBackend) Read() error {
	return nil
}

// Write will store an update to the build context storage.
func (l *LocalBackend) Write(ctx up.BuildContext) error {
	return nil
}
