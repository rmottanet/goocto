package inputs


import (
	"os"
	"fmt"
	
	"github.com/spf13/cobra"
	
	"goocto/cli/pkg/core/repository"
)


// NewDeleteRepoCmd - command to delete a repository.
func NewDeleteRepoCmd(repoHandler *repository.RepoCoreHandler) *cobra.Command {
    var deleteRepoCmd = &cobra.Command{
        Use:     "del [repository]",
        Aliases: []string{"d"},
        Short:   "Deletes a repository on GitHub.",
        Args:    cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            
            repoName := args[0]
            repoOwner := cmd.Flag("owner").Value.String()

            if repoOwner == "" {
                // default is loged user
                repoOwner = os.Getenv("GITHUB_USER")
            }

            err := repoHandler.DeleteRepo(repoOwner, repoName)
            if err != nil {
                return err
            }
            fmt.Println("Repository deleted successfully.")
            return nil
        },
    }

    deleteRepoCmd.Flags().StringP("owner", "o", "", "Owner of the repository")

    deleteRepoCmd.SilenceUsage = true
    return deleteRepoCmd
}
