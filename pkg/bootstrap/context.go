package bootstrap

import (
	"fmt"

	"github.com/vastness-io/vastup/pkg/cri"
	"github.com/vastness-io/vastup/pkg/util"
)

// BuildContext contains the context for each component that is created from
// the component repository.
type BuildContext struct {
	ComponentName  string              `json:"component_name"`
	RepoPath       string              `json:"repo_path,omitempty"`
	BinPath        string              `json:"bin_path,omitempty"`
	LogMountPath   string              `json:"log_mount_path,omitempty"`
	ContainerImage *cri.ContainerImage `json:"container_image"`
}

// Config holds the configuration of the application and loads each components specific
// context. If a component context is not available, then the component container image
// will be pulled from the registry.
type Config struct {
	LogLevel string
	Context  []*BuildContext
}

func ValidateBuildContext(components []*BuildContext) error {
	for _, componentCtx := range components {
		if err := util.ValidateContext(componentCtx.RepoPath, componentCtx.ComponentName); err != nil {
			return fmt.Errorf("failed to validate %s context: %s", componentCtx.ComponentName, err)
		}
	}
	return nil
}
