package util

import (
	"fmt"
	"net/http"
	"io/ioutil"

	"bytes"
)


// HTTPClient interface.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}


// DefaultHTTPClient default HTTPClient using http.Client.
type DefaultHTTPClient struct{}


func (c *DefaultHTTPClient) Do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

// RequestOptions
type RequestOptions struct {
	Headers map[string]string
	Client  HTTPClient
}


// FetchData faz uma requisição HTTP para a URL fornecida e retorna os bytes da resposta.
func FetchData(url string, options ...RequestOptions) ([]byte, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("error creating HTTP request: %v", err)
    }

    var client HTTPClient
    for _, opt := range options {
        if opt.Headers != nil {
            for key, value := range opt.Headers {
                req.Header.Set(key, value)
            }
        }
        if opt.Client != nil {
            client = opt.Client
        } else {
            client = &DefaultHTTPClient{}
        }
    }

    var responseData []byte

    for {
        resp, err := client.Do(req)
        if err != nil {
            return nil, fmt.Errorf("error when making HTTP request: %v", err)
        }
        defer resp.Body.Close()

        if resp.StatusCode < 200 || resp.StatusCode >= 300 {
            return nil, fmt.Errorf("HTTP response was unsuccessful, status code: %d", resp.StatusCode)
        }

        data, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return nil, fmt.Errorf("error reading HTTP response: %v", err)
        }

        responseData = append(responseData, data...)

        nextPageURL := GetNextPageURLFromLinkHeader(resp.Header)
        if nextPageURL == "" {
            break
        }

        req, err = http.NewRequest("GET", nextPageURL, nil)
        if err != nil {
            return nil, fmt.Errorf("error creating HTTP request for next page: %v", err)
        }
    }

    return responseData, nil
}


// SubmitData 
func SubmitData(url string, data []byte, options RequestOptions) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error when creating HTTP POST request: %v", err)
	}

	if options.Headers != nil {
		for key, value := range options.Headers {
			req.Header.Set(key, value)
		}
	}

	var client HTTPClient
	if options.Client != nil {
		client = options.Client
	} else {
		client = &DefaultHTTPClient{}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error when making HTTP POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP POST response was unsuccessful, status code: %d", resp.StatusCode)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading HTTP POST response: %v", err)
	}

	return responseData, nil
}


// UpdateData 
func UpdateData(url string, data []byte, options RequestOptions) ([]byte, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP PUT request: %v", err)
	}

	if options.Headers != nil {
		for key, value := range options.Headers {
			req.Header.Set(key, value)
		}
	}

	var client HTTPClient
	if options.Client != nil {
		client = options.Client
	} else {
		client = &DefaultHTTPClient{}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error when making HTTP PUT request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP PUT response was unsuccessful, status code: %d", resp.StatusCode)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading HTTP PUT response: %v", err)
	}

	return responseData, nil
}


// PatchData 
func PatchData(url string, data []byte, options RequestOptions) ([]byte, error) {
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP PATCH request: %v", err)
	}

	if options.Headers != nil {
		for key, value := range options.Headers {
			req.Header.Set(key, value)
		}
	}

	var client HTTPClient
	if options.Client != nil {
		client = options.Client
	} else {
		client = &DefaultHTTPClient{}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error when making HTTP PATCH request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP PATCH response was unsuccessful, status code: %d", resp.StatusCode)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading HTTP PATCH response: %v", err)
	}

	return responseData, nil
}


// deletter
func DeleteData(url string, options RequestOptions) error {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("error when creating HTTP DELETE request: %v", err)
	}

	if options.Headers != nil {
		for key, value := range options.Headers {
			req.Header.Set(key, value)
		}
	}

	var client HTTPClient
	if options.Client != nil {
		client = options.Client
	} else {
		client = &DefaultHTTPClient{}
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error when making HTTP DELETE request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP DELETE response was not successful, status code: %d", resp.StatusCode)
	}

	return nil
}
