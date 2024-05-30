package v1

import (
	"github.com/spf13/cobra"
)

const (
	GitHub = iota
	GitLab
)

func init() {
	rootCmd.AddCommand(config)
}

var config = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  "",
	Run:   configRun,
}

func configRun(cmd *cobra.Command, args []string) {
}
