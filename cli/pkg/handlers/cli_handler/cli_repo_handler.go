package cli_handler


import (
	"fmt"
	"errors"
	
	"github.com/spf13/cobra"
	
	"goocto/cli/pkg/handlers/gh_handler"
	"goocto/cli/pkg/models"
)

// NewCreateRepoCmd - command to create a repository.
func NewCreateRepoCmd(repoHandler *gh_handler.RepoGHHandler) *cobra.Command {
	var createRepoCmd = &cobra.Command{
		Use:   "new [repository name] [description]",
		Short: "Creates a new repository for the logged in user",
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

			err := repoHandler.CreateNewRepo(params)
			if err != nil {
				return err
			}

			fmt.Println("New repository created successfully")
			return nil
		},
	}
	return createRepoCmd
}


// NewEditInfoRepoCmd - command to edit a repository.
func NewEditInfoRepoCmd(repoHandler *gh_handler.RepoGHHandler) *cobra.Command {
    var editInfoRepoCmd = &cobra.Command{
        Use:   "edit [owner] [repository]",
        Short: "Edit information from a repository on GitHub",
        Args:  cobra.MinimumNArgs(2),
        RunE: func(cmd *cobra.Command, args []string) error {

            repoOwner := args[0]
            repoName := args[1]

            if cmd.Flags().Changed("topics") {

                topics, _ := cmd.Flags().GetStringSlice("topics")
                
                err := repoHandler.EditRepoTopics(repoOwner, repoName, topics)
                if err != nil {
                    return err
                }
                fmt.Println("Topics updated successfully")
                return nil
            }

            description, _ := cmd.Flags().GetString("description")
            name, _ := cmd.Flags().GetString("name")
            homepage, _ := cmd.Flags().GetString("homepage")

            if description == "" && name == "" && homepage == "" {
                return errors.New("at least one edit parameter must be provided")
            }

            params := models.RepoEditionParams{
                Name:        name,
                Description: description,
                Homepage:    homepage,
            }

            err := repoHandler.EditInfoRepo(repoOwner, repoName, params)
            if err != nil {
                return err
            }

            fmt.Println("Repository updated successfully")
            return nil
        },
    }

    editInfoRepoCmd.Flags().String("name", "", "New name")
    editInfoRepoCmd.Flags().String("description", "", "New description")
    editInfoRepoCmd.Flags().String("homepage", "", "New homepage")
    editInfoRepoCmd.Flags().StringSlice("topics", []string{}, "New topics. ex: \"topic1,topic2\"")

    editInfoRepoCmd.Flags().SortFlags = false

    editInfoRepoCmd.Example = `
  # Edit repository name and description:
  goocto edit <owner> <repository> --name <new name> --description <new description>

  # Edit repository homepage:
  goocto edit <owner> <repository> --homepage <new homepage>

  # Edit repository topics:
  goocto edit <owner> <repository> --topics "<topic1>,<topic2>,<topic3>"
`

    return editInfoRepoCmd
}


// NewDeleteRepoCmd - command to delete a repository.
func NewDeleteRepoCmd(repoHandler *gh_handler.RepoGHHandler) *cobra.Command {
	var deleteRepoCmd = &cobra.Command{
        Use:   "del [owner] [repository]",
        Short: "Deletes a repository on GitHub",
        Args:  cobra.MinimumNArgs(2),
        RunE: func(cmd *cobra.Command, args []string) error {

            repoOwner := args[0]
            repoName := args[1]

			err := repoHandler.DeleteRepo(repoOwner, repoName)
			if err != nil {
				return err
			}

			fmt.Println("Repository deleted successfully")
			return nil
		},
	}
	return deleteRepoCmd
}
