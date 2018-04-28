package cli

import (
	"github.com/spf13/cobra"
)

var (
	coordinatorContextPath string
	vcsWebHookContext      string

	// Up will start the platform components in the order of dependencies.
	Up = &cobra.Command{
		Use:   "",
		Short: "",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	// SetContext will set the component directory context that will be used
	// to build the component container images.
	SetContext = &cobra.Command{
		Use:   "set-context",
		Short: "",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	// Setbackend will set the applications backend for storing the bootstrap
	// configuration.
	SetBackend = &cobra.Command{
		Use:   "set-backend",
		Short: "",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)
