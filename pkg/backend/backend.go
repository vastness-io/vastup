package backend

import (
	"github.com/vastness-io/vastup/pkg/up"
)

// Backend abstracts away the persistent options for vastup.
type Backend interface {
	Read() error
	Write(ctx up.BuildContext) error
}
