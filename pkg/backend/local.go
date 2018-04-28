package backend

import (
	"github.com/PI-Victor/vastup/pkg/up"
)

// LocalBackend is the default backend used for storing context and component
// configuration for vastup.
type LocalBackend struct{}

// Read will return a context
func (l *LocalBackend) Read(ctx up.BuildContext) error {

}
