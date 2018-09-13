package bootstrap

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// RetryFunc is a docker component stage that we want to retry.
type RetryFunc func(*BuildContext, execFunc) bool

type DockerClient struct {
	*client.Client
	ctx context.Context
}

// NewDockerClient returns a wrapper docker client
func NewDockerClient() (*DockerClient, error) {
	ctx := context.Background()

	client, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	if _, err := client.Ping(ctx); err != nil {
		return nil, err
	}

	return &DockerClient{
		client,
		context.Background(),
	}, nil
}

func (d *DockerClient) createPrivateNetwork(buildCtx *BuildContext) error {
	networkOpts := types.NetworkCreate{
		Internal: true,
		Labels: map[string]string{
			"platform":     "vastness",
			"network-type": "private",
		},
	}
	_, err := d.NetworkCreate(d.ctx, buildCtx.NetworkName, networkOpts)

	if err != nil {
		return err
	}

	return nil
}

func (d *DockerClient) pullImage(buildCtx *BuildContext) error {
	// NOTE: this will 100% panic when these details are not defined in the yaml config.
	// NOTE: we can't really have a password without a user, this will assume the same
	// and override the pass if you specified it without a user.
	// NOTE: environment variables preceed and override application config defined values.
	var encodedAuth string
	user := os.Getenv("VASTNESS_REGISTRY_USER")
	password := os.Getenv("VASTNESS_REGISTRY_PASSWORD")

	if user != "" {
		auth := fmt.Sprintf("%s:%s", user, password)
		encodedAuth = base64.StdEncoding.EncodeToString([]byte(auth))
	}

	if buildCtx.Image.RegistryAuth == "" && user != "" {
		buildCtx.Image.RegistryAuth = encodedAuth
	}

	logrus.Infof("%s", buildCtx.Image.RegistryAuth)

	pullImageOpts := types.ImagePullOptions{
		RegistryAuth: buildCtx.Image.RegistryAuth,
	}
	_, err := d.ImagePull(d.ctx, buildCtx.Image.Name, pullImageOpts)

	if err != nil {
		return err
	}

	return nil
}

func (d *DockerClient) runImage(buildCtx *BuildContext) error {
	return nil
}

func (d *DockerClient) prunePrivateNetwork(buildCtx *BuildContext) error {
	return nil
}

func (d *DockerClient) createVolumes(buildCtx *BuildContext) error {
	return nil
}

func (d *DockerClient) pruneVolumes(buildCtx *BuildContext) error {
	return nil
}

func (d *DockerClient) mountVolumes(buildCtx *BuildContext) error {
	return nil
}
