package repository

import (
	"fmt"
	"encoding/json"
	
	"goocto/cli/pkg/models"
)


// CreateNewRepo cria um repositório na conta do usuário logado ou em uma organização especificada.
func (hrepo *RepoCoreHandler) CreateNewRepo(params models.RepoCreationParams, owner string) error {

    var subroute string
    if owner == "" {
        subroute = "/user/repos"
    } else {
        subroute = fmt.Sprintf("/orgs/%s/repos", owner)
    }

	// to create as option
    if !params.Private {
        params.Private = true
    }
    if !params.IsTemplate {
        params.IsTemplate = false
    }

    requestDataJSON, err := json.Marshal(params)
    if err != nil {
        return fmt.Errorf("error when converting data to JSON: %v", err)
    }

    _, err = hrepo.githubAdapter.Post(subroute, requestDataJSON)
    if err != nil {
        return err
    }

    return nil
}
