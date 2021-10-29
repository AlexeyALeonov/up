package cmd

import (
	"fmt"
	"github.com/compose-spec/compose-go/types"
	"github.com/elek/sjr/pkg/common"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "print all the configured services from the dockerfile",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Use all services to allow checking for any possible service
			composeProject, err := common.LoadComposeFromFile(ComposeFile)
			updatedComposeProject, err := common.UpdateEach(composeProject, list, "", []string{"storj", "db"})
			if err != nil {
				return err
			}
			return common.WriteComposeFile(updatedComposeProject)
		},
	}
}

func init() {
	rootCmd.AddCommand(ListCmd())
}

func list(composeService *types.ServiceConfig, _ string) error {
	fmt.Println(composeService.Name, composeService.Image)
	return nil
}