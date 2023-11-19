package api

import "strings"

//fuction that extracts the bearer token from the request header
func extractBearerToken(authorization string) string {
	car authorization1 = strings.Split(authorization, " ")
	if len(authorization1) == 2 {
		return authorization1[1]
	}

}