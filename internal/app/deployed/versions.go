package deployed

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(versions)
}

var versions = &cobra.Command{
	Use:     "versions",
	Aliases: []string{""},
	Short:   "",
	Long:    "",
	Run:     versionsRun,
}

func versionsRun(cmd *cobra.Command, args []string) {
}

func CreateVersionsFile() error {
	return nil
}
