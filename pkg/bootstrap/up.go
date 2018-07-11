package bootstrap

import (
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	defaultCoordinatorImage = "quay.io/vastness/coordinator:latest"
	defaultLinguistImage    = "quay.io/vastness/linguist:latest"
	defaultVCSWebHookImage  = "quay.io/vastness/vcs-webhook:latest"
	defaultParserImage      = "quay.io/vastness/parser:latest"
)

// Up will start the component bootstrap process
func Up(config *Config) {
	var wg sync.WaitGroup
	wg.Add(len(config.Context))

	logrus.Debugf("Connecting to docker...")

	logrus.Debugf("Starting bootstrapper with config %#v:", config)

	for _, component := range config.Context {
		go func(c *BuildContext) {
			defer wg.Done()
			logrus.Infof("Bootstrapping %s component.", c.ComponentName)
			
		}(component)
		wg.Wait()
	}
}
