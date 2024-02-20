// pkg/util/paginate.go
package util

import (
    "net/http"
    "strings"
)

func GetNextPageURLFromLinkHeader(header http.Header) string {
    linkHeader := header.Get("Link")
    if linkHeader == "" {
        return ""
    }

    links := strings.Split(linkHeader, ",")
    for _, link := range links {
        parts := strings.Split(strings.TrimSpace(link), ";")
        if len(parts) < 2 {
            continue
        }
        url := strings.Trim(parts[0], "<>")
        rel := strings.Trim(parts[1], ` "`)
        if rel == "rel=\"next\"" {
            return url
        }
    }

    return ""
}
