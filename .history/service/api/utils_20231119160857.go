package api

import "strings"

//  fuction that extracts the bearer token from the request header
func extractBearerToken(authorization string) string {
	car authorization = strings.Split(authorization, " ")

}