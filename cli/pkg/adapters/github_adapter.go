package github_api

import (
    "net/http"
    
    "goocto/cli/pkg/util"
)


// URLBase is the base URL of the GitHub API.
const URLBase = "https://api.github.com"


// GitHubAPI defines the interface for interacting with the GitHub API.
type GitHubAPI interface {
    Get(subroute string) ([]byte, error)
    Post(subroute string, data []byte) ([]byte, error)
    Put(subroute string, data []byte) ([]byte, error)
    Patch(subroute string, data []byte) ([]byte, error)
    Delete(subroute string) error
}


// GitHubAdapter implements the GitHubAPI interface.
type GitHubAdapter struct {
    apiKey string
    client *http.Client
    lastResponse *http.Response
}


// NewGitHubAdapter creates a new instance of GitHubAdapter with the specified API key.
func NewGitHubAdapter(apiKey string) *GitHubAdapter {
    return &GitHubAdapter{
        apiKey: apiKey,
        client: &http.Client{},
    }
}


// Get makes a GET request to the specified subroute in the GitHub API.
func (g *GitHubAdapter) Get(subroute string) ([]byte, error) {
    url := URLBase + subroute
    headers := map[string]string{
        "X-GitHub-Api-Version": "2022-11-28",
        "Accept":               "application/vnd.github+json",
        "Authorization":        "Bearer " + g.apiKey,
    }

    responseData, err := util.FetchData(url, util.RequestOptions{Headers: headers})
    if err != nil {
        return nil, err
    }

    g.lastResponse = util.LastHTTPResponse()

    return responseData, nil
}


// Post makes a POST request to the specified subroute in the GitHub API.
func (g *GitHubAdapter) Post(subroute string, data []byte) ([]byte, error) {
    url := URLBase + subroute
    headers := map[string]string{
        "X-GitHub-Api-Version": "2022-11-28",
        "Accept":               "application/vnd.github+json",
        "Authorization":        "Bearer " + g.apiKey,
    }

    responseData, err := util.SubmitData(url, data, util.RequestOptions{Headers: headers})
    if err != nil {
        return nil, err
    }

    g.lastResponse = util.LastHTTPResponse()

    return responseData, nil
}


// Put makes a PUT request to the specified subroute in the GitHub API.
func (g *GitHubAdapter) Put(subroute string, data []byte) ([]byte, error) {
    url := URLBase + subroute
    headers := map[string]string{
        "X-GitHub-Api-Version": "2022-11-28",
        "Accept":               "application/vnd.github+json",
        "Authorization":        "Bearer " + g.apiKey,
    }

    responseData, err := util.UpdateData(url, data, util.RequestOptions{Headers: headers})
    if err != nil {
        return nil, err
    }

    g.lastResponse = util.LastHTTPResponse()

    return responseData, nil
}


// Patch makes a PATCH request to the specified subroute in the GitHub API.
func (g *GitHubAdapter) Patch(subroute string, data []byte) ([]byte, error) {
    url := URLBase + subroute
    headers := map[string]string{
        "X-GitHub-Api-Version": "2022-11-28",
        "Accept":               "application/vnd.github+json",
        "Authorization":        "Bearer " + g.apiKey,
    }

    responseData, err := util.PatchData(url, data, util.RequestOptions{Headers: headers})
    if err != nil {
        return nil, err
    }

    g.lastResponse = util.LastHTTPResponse()

    return responseData, nil
}


// Delete makes a DELETE request to the specified subroute in the GitHub API.
func (g *GitHubAdapter) Delete(subroute string) error {
    url := URLBase + subroute
    headers := map[string]string{
        "X-GitHub-Api-Version": "2022-11-28",
        "Accept":               "application/vnd.github+json",
        "Authorization":        "Bearer " + g.apiKey,
    }

    err := util.DeleteData(url, util.RequestOptions{Headers: headers})
    return err
}


// LastResponse returns the last HTTP response received by the adapter.
func (g *GitHubAdapter) LastResponse() *http.Response {
    return g.lastResponse
}
