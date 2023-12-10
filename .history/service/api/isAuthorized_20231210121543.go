package api

import (
	"net/http"
)

func isAuthorized(header http.Header) string {
	authToken, err := header.Get
	if err != nil {
		return ""
	}
	return authToken
}

