package main

import (
	"os"
	"fmt"
	
	"github.com/spf13/cobra"
	
	"goocto/cli/pkg/config"
	"goocto/cli/pkg/adapters"
	"goocto/cli/pkg/handlers/gh_handler"
	"goocto/cli/pkg/handlers/cli_handler"
)

func main() {
	// Load environment
	config.LoadEnv()

	APIKEY := os.Getenv("APIKEY")
	if APIKEY == "" {
	    fmt.Println("The APIKEY environment variable is not defined.")
	    os.Exit(1)
	}

    // Start adapters
    githubAdapter := github_api.NewGitHubAdapter(APIKEY)
    	
	// Start handlers
	repoHandler := gh_handler.NewRepoGHHandler(githubAdapter)


	// Start Cobra
	rootCmd := &cobra.Command{
		Use:   "goocto",
		Short: "GoOcto - A GitHub CLI",
	}

	// Command "new" to create a new repository
	createRepoCmd := cli_handler.NewCreateRepoCmd(repoHandler)
	rootCmd.AddCommand(createRepoCmd)

	// Command "edit" to edit repository details
	editInfoRepoCmd := cli_handler.NewEditInfoRepoCmd(repoHandler)
	rootCmd.AddCommand(editInfoRepoCmd)

	// Command "del" to delete repository
	deleteRepoCmd := cli_handler.NewDeleteRepoCmd(repoHandler)
	rootCmd.AddCommand(deleteRepoCmd)
		
	// Run main command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
