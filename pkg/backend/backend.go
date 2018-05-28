package backend

import (
	"github.com/vastness-io/vastup/pkg/bootstrap"
)

// Backend abstracts away the persistent options for vastup.
type Backend interface {
	Read() error
	Write(ctx *bootstrap.BuildContext) error
}
