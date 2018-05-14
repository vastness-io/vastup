package cri

import (
	"context"

	_ "github.com/docker/docker/api/types"
	_ "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	_ "golang.org/x/net/context"

	"github.com/vastness-io/vastup/pkg/up"
)

// DockerCRI holds functionality to interact with the docker API.
type DockerCRI struct {
	Client *client.Client
}

// NewDockerClient returns a new docker client
func NewDockerClient() (*DockerCRI, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &DockerCRI{
		client,
	}, nil
}

// BuildImage will take
func (d *DockerCRI) BuildImage(ctx context.Context, buildCtx up.BuildContext) error {
	return nil
}

func (d *DockerCRI) PullImage(ctx context.Context, buildCtx up.BuildContext) error {
	return nil
}

func (d *DockerCRI) RebuildImage(ctx context.Context, buildCtx up.BuildContext) error {
	return nil
}

func (d *DockerCRI) RunImage(ctx context.Context, buildCtx up.BuildContext) error {
	return nil
}
