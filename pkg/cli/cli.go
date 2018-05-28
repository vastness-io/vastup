package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/vastness-io/vastup/pkg/bootstrap"
)

var (
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

			if err := viper.Unmarshal(config); err != nil {
				logrus.Panicf("Error while loading config: %#v", err)
			}

			if config.CoordinatorContext != nil {
				logrus.Info("Validating context for Coordinator")
				if err := bootstrap.ValidateBuildContext(config.CoordinatorContext); err != nil {
					logrus.Errorf("Failed to validate context for Coordinator: %#v", err)
				}
			}

			if config.VCSWebHookContext != nil {
				logrus.Info("Validating context for VCS-Webhook")
				if err := bootstrap.ValidateBuildContext(config.VCSWebHookContext); err != nil {
					logrus.Errorf("Failed to validate context for VCS-Webhook: %#v", err)
				}
			}

			if config.LinguistContext != nil {
				logrus.Info("Validating context for Linguist")
				if err := bootstrap.ValidateBuildContext(config.LinguistContext); err != nil {
					logrus.Errorf("Failed to validate context for Linguist: %#v", err)
				}
			}

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
