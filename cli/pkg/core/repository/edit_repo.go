package repository

import (
	"fmt"
	"errors"
	"encoding/json"

	"goocto/cli/pkg/models"
)


// EditInfoRepo edits the fields of a repository on the logged-in user's GitHub account.
func (hrepo *RepoCoreHandler) EditInfoRepo(repoOwner string, repoName string, params models.RepoEditionParams) error {
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
        return err
    }

    return nil
}


// EditRepoTopics edits the list of topics of a repository.
func (hrepo *RepoCoreHandler) EditRepoTopics(repoOwner string, repoName string, newTopics []string) error {
    subroute := fmt.Sprintf("/repos/%s/%s/topics", repoOwner, repoName)

    topicsData := map[string][]string{"names": newTopics}
    requestData, err := json.Marshal(topicsData)
    if err != nil {
        return fmt.Errorf("error converting topic data to JSON: %v", err)
    }

    _, err = hrepo.githubAdapter.Put(subroute, requestData)
    if err != nil {
        return err
    }

    return nil
}
