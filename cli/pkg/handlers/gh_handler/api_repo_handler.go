package gh_handler

import (
	"fmt"
	"errors"
	"encoding/json"

    "goocto/cli/pkg/adapters"	
	"goocto/cli/pkg/models"
)


// RepoGHHandler contains the handlers for operations related to repositories in the GitHub API.
type RepoGHHandler struct {
    githubAdapter *github_api.GitHubAdapter
}


// NewRepoGHHandler creates a new instance of RepoGHHandler.
func NewRepoGHHandler(adapter *github_api.GitHubAdapter) *RepoGHHandler {
    return &RepoGHHandler{
        githubAdapter: adapter,
    }
}


// CreateNewRepo creates a repository on the logged-in user's GitHub account.
func (hrepo *RepoGHHandler) CreateNewRepo(params models.RepoCreationParams) error {
    if params.Private == false {
        params.Private = true
    }
    if params.IsTemplate == false {
        params.IsTemplate = false
    }

    requestDataJSON, err := json.Marshal(params)
    if err != nil {
        return fmt.Errorf("error when converting data to JSON: %v", err)
    }

    subroute := "/user/repos"
    _, err = hrepo.githubAdapter.Post(subroute, requestDataJSON)
    if err != nil {
        return fmt.Errorf("error creating repository on GitHub: %v", err)
    }

    return nil
}


// EditInfoRepo edits the fields of a repository on the logged-in user's GitHub account.
func (hrepo *RepoGHHandler) EditInfoRepo(repoOwner string, repoName string, params models.RepoEditionParams) error {
    if params.Name == "" && params.Description == "" && params.Homepage == "" {
        return errors.New("no edit parameters provided")
    }

    updateFields := make(map[string]interface{})

    if params.Name != "" {
        updateFields["name"] = params.Name
    }
    if params.Description != "" {
        updateFields["description"] = params.Description
    }
    if params.Homepage != "" {
        updateFields["homepage"] = params.Homepage
    }

    updateData, err := json.Marshal(updateFields)
    if err != nil {
        return fmt.Errorf("error converting update data to JSON: %v", err)
    }

    subroute := fmt.Sprintf("/repos/%s/%s", repoOwner, repoName)
    _, err = hrepo.githubAdapter.Patch(subroute, updateData)
    if err != nil {
        return fmt.Errorf("error when editing repository on GitHub: %v", err)
    }

    return nil
}


// EditRepoTopics edits the list of topics of a repository.
func (hrepo *RepoGHHandler) EditRepoTopics(repoOwner string, repoName string, newTopics []string) error {
    subroute := fmt.Sprintf("/repos/%s/%s/topics", repoOwner, repoName)

    topicsData := map[string][]string{"names": newTopics}
    requestData, err := json.Marshal(topicsData)
    if err != nil {
        return fmt.Errorf("error converting topic data to JSON: %v", err)
    }

    _, err = hrepo.githubAdapter.Put(subroute, requestData)
    if err != nil {
        return fmt.Errorf("error when updating repository topics on GitHub: %v", err)
    }

    return nil
}


// DeleteRepo deletes a repository from GitHub.
func (hrepo *RepoGHHandler) DeleteRepo(repoOwner, repoName string) error {
	subroute := fmt.Sprintf("/repos/%s/%s", repoOwner, repoName)

	err := hrepo.githubAdapter.Delete(subroute)
	if err != nil {
		return fmt.Errorf("error when deleting repository on GitHub: %v", err)
	}

	return nil
}

