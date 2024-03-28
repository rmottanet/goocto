package repository

import (
    "goocto/cli/pkg/adapters"	
	"goocto/cli/pkg/models"
)


// RepoCoreHandler contains the handlers for operations related to repositories in the GitHub API.
type RepoCoreHandler struct {
    githubAdapter *github_api.GitHubAdapter
}


type RepoCore interface {
    CreateNewRepo(params models.RepoCreationParams) error
    EditInfoRepo(repoOwner, repoName string, params models.RepoEditionParams) error
    EditRepoTopics(repoOwner, repoName string, newTopics []string) error
    DeleteRepo(repoOwner, repoName string) error
}

// NewRepoCoreHandler creates a new instance of RepoCoreHandler.
func NewRepoCoreHandler(adapter *github_api.GitHubAdapter) *RepoCoreHandler {
    return &RepoCoreHandler{
        githubAdapter: adapter,
    }
}
