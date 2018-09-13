package bootstrap

import (
	"errors"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

// NOTE: these should be extracted to the config.
const (
	defaultCoordinatorImage = "quay.io/vastness/coordinator:latest"
	defaultLinguistImage    = "quay.io/vastness/linguist:latest"
	defaultVCSWebHookImage  = "quay.io/vastness/vcs-webhook:latest"
	defaultParserImage      = "quay.io/vastness/parser:latest"
)

// Up will start the component bootstrap process
func Up(config *Config) {
	if config.Components == nil {
		logrus.Fatalf("No component context defined in configuration file! Define at least one!")
	}

	errChan := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(len(config.Components))

	logrus.Debugf("Connecting to docker...")
	d, err := NewDockerClient()

	if err != nil {
		logrus.Fatalf("Failed to connect to docker client: %s", err)
	}

	logrus.Debugf("Successfully connected to docker!")

	go boostrapCore(config, d, errChan)
	if err := <-errChan; err != nil {
		logrus.Fatal(err)
	}

	logrus.Debugf("Starting bootstrapper with config %#v:", config)

	for _, componentCtx := range config.Components {
		go func(c *BuildContext) {
			defer wg.Done()

			logrus.Infof("Bootstrapping %s component.", c.Name)
			// NOTE: be careful about mutating DockerClient.
			err := bootstrapStages(c, d)
			if err != nil {
				logrus.Warn(err)
			}
		}(componentCtx)
	}
	wg.Wait()
}

func boostrapCore(config *Config, client *DockerClient, errChan chan error) {
	logrus.Info("Commencing core docker components bootstrap...")
	var networkName string

	if config.Network == nil || config.Network.Name == "" {
		// TODO: Make this unique!
		networkName = "auto-generated"
	} else {
		networkName = config.Network.Name
	}
	networkCtx := &BuildContext{
		NetworkName: networkName,
	}

	if ok := RetryOnFailure(networkCtx, client.createPrivateNetwork); !ok {
		errChan <- errors.New("Failed to bootstrap core docker components")
	}

	errChan <- nil
}

func bootstrapStages(c *BuildContext, d *DockerClient) error {
	c.Stage = "pullImage"
	if ok := RetryOnFailure(c, d.pullImage); !ok {
		return fmt.Errorf("Failed to bootstrap component: %s at stage: %s", c.Name, c.Stage)
	}

	c.Stage = "mountVolumes"
	if ok := RetryOnFailure(c, d.createVolumes); !ok {
		return fmt.Errorf("Failed to bootstrap component: %s at stage: %s", c.Name, c.Stage)
	}

	c.Stage = "startContainer"
	if ok := RetryOnFailure(c, d.runImage); !ok {
		return fmt.Errorf("Failed to bootstrap component: %s at stage: %s", c.Name, c.Stage)
	}
	return nil
}
