package bootstrap

import (
	"fmt"
	"math"
	"time"

	"github.com/vastness-io/vastup/pkg/util"
)

// Config holds the configuration of the application and loads each components specific
// context. If a component context is not available, then the component container image
// will be pulled from the registry.
type Config struct {
	LogLevel string
	Context  []*BuildContext
}

// BuildContext contains the context for each component that is created from
// the component repository.
type BuildContext struct {
	ComponentName  string          `json:"component_name"`
	RepoPath       string          `json:"repo_path,omitempty"`
	BinPath        string          `json:"bin_path,omitempty"`
	LogMountPath   string          `json:"log_mount_path,omitempty"`
	ContainerImage *ContainerImage `json:"container_image"`
}

// ContainerImage holds metadata about the image used by the resource.
type ContainerImage struct {
	Name     string `json:"name"`
	Registry string `json:"registry"`
	BaseSHA  string `json:"base_sha"`
	Tag      string `json:"tag"`
}

func ValidateBuildContext(components []*BuildContext) error {
	for _, componentCtx := range components {
		if err := util.ValidateContext(componentCtx.RepoPath, componentCtx.ComponentName); err != nil {
			return fmt.Errorf("failed to validate %s context: %s", componentCtx.ComponentName, err)
		}
	}
	return nil
}

type execFunc func(*BuildContext) error

// RetryOnFailure implements an exponential backoff algorithm that will
// retry a function till hitting errorMaxCap.
func RetryOnFailure(f execFunc, ctx *BuildContext) bool {
	const errorMaxCap = 10

	for i := 0; i <= errorMaxCap; i++ {
		var (
			t    = time.Duration(math.Pow(2, float64(i)))
			wait = time.Millisecond * t
		)

		if err := f(ctx); err != nil {
			time.Sleep(wait)
			continue
		}
		return true
	}
	return false
}
