package deployed

import (
	"fmt"
	"os"
	"strconv"

	"github.com/deployix/deployed/internal/constants"
	"github.com/deployix/deployed/internal/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP("force", "f", false, "Force initialization")
}

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{""},
	Short:   "",
	Long:    "",
	RunE:    initRun,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	},
}

func initRun(cmd *cobra.Command, args []string) error {
	force, _ := cmd.Flags().GetBool("force")
	// create working directory
	if err := generateWorkingDir(force); err != nil {
		return err
	}

	// get application name
	applicationNamePrompt := promptui.Prompt{
		Label: "Application Name",
	}

	applicationName, err := applicationNamePrompt.Run()
	if err != nil {
		return err
	}

	// get git provider
	gitProviderPrompt := promptui.Prompt{
		Label:     "git provider",
		Default:   "github",
		AllowEdit: true,
	}

	gitProvider, err := gitProviderPrompt.Run()
	if err != nil {
		return err
	}

	// get git domain
	gitDomainPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("%s's domain (i.e. 'github.com')", gitProvider),
		Default: "",
	}

	gitDomain, err := gitDomainPrompt.Run()
	if err != nil {
		return err
	}

	// get git repo name
	gitRepoNamePrompt := promptui.Prompt{
		Label: "repository name",
	}

	repoName, err := gitRepoNamePrompt.Run()
	if err != nil {
		return err
	}

	// get default branch
	gitDefaultBranchPrompt := promptui.Prompt{
		Label:     "default branch",
		Default:   constants.DEFAULT_GIT_BRANCH,
		AllowEdit: true,
	}

	defaultBranch, err := gitDefaultBranchPrompt.Run()
	if err != nil {
		return err
	}

	// get datetime format
	dateTimeFormatPrompt := promptui.Prompt{
		Label:     "datetime format",
		Default:   constants.DEFAULT_DATETIME_FORMAT,
		AllowEdit: true,
	}

	dateTimeFormat, err := dateTimeFormatPrompt.Run()
	if err != nil {
		return err
	}

	// get default channel
	defaultChannelPrompt := promptui.Prompt{
		Label: "default channel",
	}

	defaultChannel, err := defaultChannelPrompt.Run()
	if err != nil {
		return err
	}

	// get maximumVersion History length
	maxVersionHistoryPrompt := promptui.Prompt{
		Label: "max version history",
	}

	maxVersionHistory, err := maxVersionHistoryPrompt.Run()
	if err != nil {
		return err
	}
	maxVersionHistoryInt, err := strconv.Atoi(maxVersionHistory)
	if err != nil {
		return err
	}

	config := Config{
		ApplicationName: applicationName,
		DateTimeFormat:  dateTimeFormat,
		DefaultBranch:   defaultBranch,
		DefaultChannel:  defaultChannel,
		ChannelsConfig: ChannelsConfig{
			MaxVersionHistoryLength: maxVersionHistoryInt,
		},
		GitConfig: GitConfig{
			Provider: gitProvider,
			Domain:   gitDomain,
			RepoName: repoName,
		},
	}

	// write to config file
	if err := config.WriteToFile(); err != nil {
		return err
	}

	channels := Channels{
		Channels: map[string]Channel{
			defaultChannel: Channel{},
		},
	}

	// write to channels file
	if err := channels.WriteToFile(); err != nil {
		return err
	}

	promotions := Promotions{}
	if err := promotions.WriteToFile(); err != nil {
		return err
	}

	return nil

}

func generateWorkingDir(force bool) error {
	if _, err := os.Stat(utils.FilePaths.GetDirectoryPath()); err == nil && !force {
		// Dir exists and we are not forcing the creation
		return fmt.Errorf("dir %s already exists. Use --force to overwrite", utils.FilePaths.GetDirectoryPath())
	} else {
		err := os.RemoveAll(utils.FilePaths.GetDirectoryPath())
		if err != nil {
			return err
		}
	}

	if err := os.Mkdir(utils.FilePaths.GetDirectoryPath(), constants.DEFAULT_DIR_FILEMODE); err != nil {
		return err
	}

	return nil
}

func createInitConfig() Config {
	return Config{}
}
