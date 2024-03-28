package inputs


import (
	"fmt"
	
	"github.com/spf13/cobra"
	
	"goocto/cli/pkg/core/repository"
	"goocto/cli/pkg/models"
)


// NewCreateRepoCmd - command to create a repository.
func NewCreateRepoCmd(repoHandler *repository.RepoCoreHandler) *cobra.Command {
	var owner string

	var createRepoCmd = &cobra.Command{
		Use:     "new [repository name] [description]",
		Aliases: []string{"-n"},
		Short:   "Creates a new repository for the logged in user or in a specified organization.",
		RunE: func(cmd *cobra.Command, args []string) error {

			var repoName, description string

			if len(args) >= 1 {
				repoName = args[0]
			}

			if len(args) >= 2 {
				description = args[1]
			}

			params := models.RepoCreationParams{
				Name:        repoName,
				Description: description,
			}

			err := repoHandler.CreateNewRepo(params, owner)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("New repository created successfully.")
			}
			return nil
		},
	}

	createRepoCmd.Flags().StringVarP(&owner, "owner", "o", "", "Owner of the repository (optional)")

	createRepoCmd.SilenceUsage = true
	return createRepoCmd
}
