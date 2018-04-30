package backend

import (
	"path"

	"github.com/PI-Victor/vastup/pkg/up"
)

// LocalBackend is the default backend used for storing context and component
// configuration for vastup.
type LocalBackend struct {
	WorkDir string
}

// Read will return the current context read from the backend.
func (l *LocalBackend) Read() error {
	return nil
}

// Write will store an update to the build context storage.
func (l *LocalBackend) Write(ctx up.BuildContext) error {
	return nil
}

// New will create a new build context and store it to disk.
func (l *LocalBackend) New() error {

	const (
		repoBinPath = "bin/linux/amd64"
	)

	newCtx := up.BuildContext{
		RepoPath: l.WorkDir,
		BinPath:  path.Join(l.WorkDir, repoBinPath),
	}

	if err := l.Write(newCtx); err != nil {
		return err
	}

	return nil
}
