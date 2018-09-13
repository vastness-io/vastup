package bootstrap

import (
	"fmt"
	"math"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/vastness-io/vastup/pkg/util"
)

// Config holds the configuration of the application and loads each components specific
// context. If a component context is not available, then the component container image
// will be pulled from the registry.
type Config struct {
	LogLevel   string          `yaml:"logLevel,omitempty"`
	Components []*BuildContext `yaml:"components"`

	// Network is not a private component implementation but a general one, since we create
	// and use a single private network for all components.
	Network *PrivateNetwork `yaml:"network,omitempty"`
}

// BuildContext contains the context for each component that is created from
// the component repository.
type BuildContext struct {
	Name           string `yaml:"name"`
	RepositoryPath string `yaml:"repositoryPath,omitempty"`
	BinPath        string `yaml:"binaryPath,omitempty"`
	LogMountPath   string `yaml:"logMountPath,omitempty"`
	Image          *Image `yaml:"image"`

	NetworkName string `yaml:"-"`

	Stage string `yaml:"-"`
}

// Image holds metadata about the image used by the resource.
type Image struct {
	Name         string `yaml:"name"`
	BaseSHA      string `yaml:"-"`
	Tag          string `yaml:"tag"`
	RegistryAuth string `yaml:"registryAuth,omitempty"`
}

// PrivateNetwork describes the networking between components.
type PrivateNetwork struct {
	Name string `yaml:"name"`
}

// ValidateBuildContext searches your local path for the vastness component binaries and
// starts a docker image build process locally with the compiled binaries found in your
// path.
func ValidateBuildContext(components []*BuildContext) error {
	for _, componentCtx := range components {
		if err := util.ValidateContext(componentCtx.RepositoryPath, componentCtx.Name); err != nil {
			return fmt.Errorf("failed to validate %s context: %s", componentCtx.Name, err)
		}
	}
	return nil
}

// execFunc represents a docker bootstrap function that satisfies a docker action in the
// component bootstrap process.
// e.g.: create a private network
type execFunc func(*BuildContext) error

// RetryOnFailure implements an exponential backoff algorithm that will
// retry a function till hitting errorMaxCap.
func RetryOnFailure(ctx *BuildContext, f execFunc) bool {
	const errorMaxCap = 10

	for i := 0; i <= errorMaxCap; i++ {
		var (
			t    = time.Duration(math.Pow(2, float64(i)))
			wait = time.Millisecond * t
		)

		if err := f(ctx); err != nil {
			logrus.Debugf("Component bootstrap stage %s failed: %s!", ctx.Stage, err)

			time.Sleep(wait)
			continue
		}
		return true
	}
	return false
}
