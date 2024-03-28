package util

import "net/http"

var lastResponse *http.Response

func SetLastHTTPResponse(resp *http.Response) {
    lastResponse = resp
}

func LastHTTPResponse() *http.Response {
    return lastResponse
}
