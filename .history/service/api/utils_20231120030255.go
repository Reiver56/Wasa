package api

import ("strings")

//  fuction that extracts the bearer token from the request header
func extractBearerToken(authorization string) string {
	fmt
	var authorization1 = strings.Split(authorization, " ")
	if len(authorization1) == 2 {
		return strings.Trim(authorization1[1], " ")
	}
	return ""

}