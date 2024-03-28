package repository

import (
	"fmt"
)


// DeleteRepo delete a GitHub repository.
func (hrepo *RepoCoreHandler) DeleteRepo(repoOwner, repoName string) error {

    subroute := fmt.Sprintf("/repos/%s/%s", repoOwner, repoName)

    err := hrepo.githubAdapter.Delete(subroute)
    if err != nil {
        return err
    }

    return nil
}
