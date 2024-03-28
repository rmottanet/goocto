package inputs


import (
	"os"
	"fmt"
	"errors"
	
	"github.com/spf13/cobra"
	
	"goocto/cli/pkg/core/repository"
	"goocto/cli/pkg/models"
)


// NewEditInfoRepoCmd - command to edit a repository.
func NewEditInfoRepoCmd(repoHandler *repository.RepoCoreHandler) *cobra.Command {
    var editInfoRepoCmd = &cobra.Command{
        Use:   "edit [repository]",
        Aliases: []string{"-e"},
        Short: "Edit information from a repository on GitHub.",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {

            repoOwner := cmd.Flag("owner").Value.String()
            repoName := args[0]

            if repoOwner == "" {
                // default is loged user
                repoOwner = os.Getenv("GITHUB_USER")
            }

            if cmd.Flags().Changed("topics") {
                topics, _ := cmd.Flags().GetStringSlice("topics")
                
                err := repoHandler.EditRepoTopics(repoOwner, repoName, topics)
                if err != nil {
                    fmt.Println(err)
                } else {
                    fmt.Println("Topics updated successfully.")
                }
                return nil
            }

            description, _ := cmd.Flags().GetString("description")
            name, _ := cmd.Flags().GetString("name")
            homepage, _ := cmd.Flags().GetString("homepage")

            if description == "" && name == "" && homepage == "" {
                return errors.New("At least one edit parameter must be provided.")
            }

            params := models.RepoEditionParams{
                Name:        name,
                Description: description,
                Homepage:    homepage,
            }

            err := repoHandler.EditInfoRepo(repoOwner, repoName, params)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Repository updated successfully.")
            }
            return nil
        },
    }

    editInfoRepoCmd.Flags().String("name", "", "New name")
    editInfoRepoCmd.Flags().String("description", "", "New description")
    editInfoRepoCmd.Flags().String("homepage", "", "New homepage")
    editInfoRepoCmd.Flags().StringSlice("topics", []string{}, "New topics. ex: \"topic1,topic2\"")
	editInfoRepoCmd.Flags().StringP("owner", "o", "", "Owner of the repository")
	
    editInfoRepoCmd.Flags().SortFlags = false

    editInfoRepoCmd.Example = `
  # Edit repository name and description:
  goocto edit <repository> --name <new name> --description <new description>

  # Edit repository homepage:
  goocto edit <repository> --homepage <new homepage>

  # Edit repository topics:
  goocto edit <repository> --topics "<topic1>,<topic2>,<topic3>"
  
  # Edit repository from an organization:
  goocto edit <repository> --topics "<topic1>,<topic2>,<topic3> --owner <org name>"
`
	editInfoRepoCmd.SilenceUsage = true
    return editInfoRepoCmd
}
