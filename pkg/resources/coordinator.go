package resources

import (
	"github.com/vastness-io/vastup/pkg/up"
)

const (
	coordinatorImage = "quay.io/vastness-io/coordinator:latest"
)

// Coordinator has the functionality to build/slice and operate the vastness
// resource.
type Coordinator struct {
	image    string
	registry string
}

// Start will start the container image for coordinator resource
func (c *Coordinator) Start() error { return nil }

// Build will search for the binary of the compiled coordinator resource within
// the given context, then it will search the history of the base docker image
// of the component and get the SHA of the container commit where the binary is
// added in the base image. It will build a new container image based on that
// SHA in the history with the new binary copied on top.
func (c *Coordinator) Build(ctx up.BuildContext) error { return nil }

// pullimage will pull the component's related image from a registry
func (c *Coordinator) pullImage(ctx up.BuildContext) error { return nil }
