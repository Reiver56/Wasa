package api

import (
	"fmt"
	"strings"
)

//  fuction that extracts the bearer token from the request header
func extractBearerToken(authorization string) string {
	fmt.Println(authorization)
	var authorization1 = strings.Split(authorization, " ")
	fmt.Println(authorization1)
	if len(authorization1) == 2 {
		fmt.Println(authorization)
		return strings.Trim(authorization1[1], " ")
	}
	return ""

}