package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/vastness-io/vastup/pkg/cli"
)

func main() {
	cmd := &cobra.Command{
		Use:   "vastup",
		Short: "vastup - Bootstrap vastness components on various environments",
		Example: `Bootstrap vastness components locally for testing:
$ vastctl run
...
Endpoint available at https://127.0.0.1:8081
`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(cli.Run)
	cmd.AddCommand(cli.SetContext)
	cmd.AddCommand(cli.SetBackend)

	if err := cmd.Execute(); err != nil {
		logrus.Errorf("Failed to load command: %s", err)
	}
}

func init() {
	// TODO: make this configurable
	logrus.SetLevel(logrus.DebugLevel)

	viper.SetConfigName(".vastup")
	viper.SetConfigType("json")

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panicf("Failed to read config file: %s", err)
	}
}
