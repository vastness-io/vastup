package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/vastness-io/vastup/pkg/bootstrap"
)

var (
	devBootstrap       bool
	coordinatorContext string
	vcsWebHookContext  string
	linguistContext    string
	parserContext      string
	backendType        string

	// Run will start the platform components.
	Run = &cobra.Command{
		Use:   "run",
		Short: "Bootstrap a new vastness platform instance",
		Run: func(cmd *cobra.Command, args []string) {
			config := &bootstrap.Config{}

			logrus.Debugf("Loading configuration from: %s", viper.ConfigFileUsed())
			if err := viper.Unmarshal(config); err != nil {
				logrus.Panicf("Error while loading config: %#v", err)
			}

			logrus.Debug("Validating components context")
			if err := bootstrap.ValidateBuildContext(config.Context); err != nil {
				logrus.Errorf("Failed to validate component context: %s", err)
			}
			// NOTE: at this point we will handle errors lower in the stack because we have
			// components that could fail the bootstrap while others are successful, thus
			// making the bootstrap operation valid.
			bootstrap.Up(config)
		},
	}

	// SetContext will set the component directory context that will be used
	// to build the component container images.
	SetContext = &cobra.Command{
		Use:   "set-context",
		Short: "Set the build context for a specific component",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// SetBackend will set the components backend method.
	// Containers - VMS, etc.
	SetBackend = &cobra.Command{
		Use:   "backend",
		Short: "Set components bootstrap backend method",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func init() {
	Run.PersistentFlags().BoolVar(
		&devBootstrap,
		"dev-bootstrap",
		false,
		"Bootstrap new components with updated local binaries",
	)

	SetContext.PersistentFlags().StringVar(
		&coordinatorContext,
		"coordinator",
		"",
		"Specify the Coordinator component context",
	)

	SetContext.PersistentFlags().StringVar(
		&vcsWebHookContext,
		"vcs-webhook",
		"",
		"Specify the VCS-Webhook component context",
	)

	SetContext.PersistentFlags().StringVar(
		&linguistContext,
		"linguist",
		"",
		"Specify the Linguist component context",
	)

	SetContext.PersistentFlags().StringVar(
		&parserContext,
		"parser",
		"",
		"Specify the Parser component context",
	)

	SetBackend.PersistentFlags().StringVar(
		&backendType,
		"type",
		"local",
		"Set the persisten backend type",
	)
}
