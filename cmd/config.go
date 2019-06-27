package cmd

import (
	"github.com/spf13/cobra"

	"github.com/yeonfeel/gangaji/config"
)

// configCmd represents the configuration
var configCmd = &cobra.Command{
	Use:   "config FLAG",
	Short: "specific controller configuration",
	Long:  `specific controller configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Load()

		// flagutil.SetIfChanged(cmd, "seq", &conf.Bartender.ClusterSeq)
		// flagutil.SetIfChanged(cmd, "cluster", &conf.Bartender.Cluster)
		// flagutil.SetIfChanged(cmd, "apps", &conf.Bartender.Apps)

		if err := config.Write(); err != nil {
			log.Error(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(configCmd)
	// configCmd.Flags().StringP("seq", "", config.Default.Bartender.Name, "Specify a name of bartender")
	// configCmd.Flags().StringP("cluster", "", config.Default.Bartender.Cluster, "Specify a cluster id")
	// configCmd.Flags().StringP("apps", "", json.ToStringWithoutError(config.Default.Bartender.Apps), "Specify a release of bartender")
}
