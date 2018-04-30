package cli

import (
	"github.com/spf13/cobra"
)

var (
	coordinatorContext string
	vcsWebHookContext  string

	// Up will start the platform components in the order of dependencies.
	Up = &cobra.Command{
		Use:   "run",
		Short: "Bootstrap a new vastness platform instance",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// SetContext will set the component directory context that will be used
	// to build the component container images.
	SetContext = &cobra.Command{
		Use:   "context",
		Short: "Set the build context for a specific component",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	// SetBackend will set the applications backend for storing the bootstrap
	// configuration.
	SetBackend = &cobra.Command{
		Use:   "backend",
		Short: "Set the backend for the bootstrap generated configuration",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)
