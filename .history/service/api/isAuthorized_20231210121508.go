package api

import (
	"net/http"
)

func isAuthorized(header http.Header) string {
	authToken, err := header.Get("Authorization")
	if err != nil {
		return ""
	}
}

