package bootstrap

import (
	_ "context"

	"github.com/docker/docker/client"
)

type DockerClient struct {
	// client.
	client.
}

func NewDockerClient() *client.Client {
	// var (
	// 	ctx = context.Background()
	// )
	//
	// client, err := client.NewEnvClient()
	// if err != nil {
	//
	// }

	return nil
}

func (d *DockerClient) createPrivateNetwork() error {
	return nil
}
