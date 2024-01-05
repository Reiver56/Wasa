package api

import (
	"net/http"
)

func isAuthorized(header http.Header) string {
	authToken:= header.Get("Authorization")
	
	return authToken
}
