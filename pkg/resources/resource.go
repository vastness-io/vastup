package resources

import (
	"github.com/PI-Victor/vastup/pkg/up"
)

// Resource abstracts away a platform resource
type Resource interface {
	Start(ctx up.BuildContext) error
	Build(ctx up.BuildContext) errorq
}
