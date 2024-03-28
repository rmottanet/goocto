package main

import (
	"os"
	"fmt"
	
	"github.com/spf13/cobra"
	
	"goocto/cli/pkg/config"
	"goocto/cli/pkg/adapters"
	"goocto/cli/pkg/core/repository"
	"goocto/cli/pkg/inputs"
)

func main() {
	// Load environment
	config.LoadEnv()
	
	
	APIKEY := os.Getenv("GITHUB_TOKEN")
	if APIKEY == "" {
	    fmt.Println("The GITHUB_TOKEN environment variable is not defined.")
	    os.Exit(1)
	}

    // Start adapters
    githubAdapter := github_api.NewGitHubAdapter(APIKEY)
    	
	// Start handlers
	repoHandler := repository.NewRepoCoreHandler(githubAdapter)


	// Start Cobra
	rootCmd := &cobra.Command{
		Use:   "goocto",
		Short: "GoOcto - A tiny GitHub CLI.",
	}

	// Command "new" to create a new repository
	createRepoCmd := inputs.NewCreateRepoCmd(repoHandler)
	rootCmd.AddCommand(createRepoCmd)

	// Command "edit" to edit repository details
	editInfoRepoCmd := inputs.NewEditInfoRepoCmd(repoHandler)
	rootCmd.AddCommand(editInfoRepoCmd)

	// Command "del" to delete repository
	deleteRepoCmd := inputs.NewDeleteRepoCmd(repoHandler)
	rootCmd.AddCommand(deleteRepoCmd)
		
	// Run main command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
