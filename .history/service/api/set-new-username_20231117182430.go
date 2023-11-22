package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

//   Function that set a new user's nickname
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	path := ps.ByName("id")
	fmt.Println(path)

	

	


	var  user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println()
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !user.isValidID() {
		fmt.Println(path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.SetNewUsername(path, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return

	

}

func validateRequestingUser(identifier string, Token string) int {

	
	if isNotLogged(Token) {
		return http.StatusForbidden
	}

	

	if identifier != Token {
		return http.StatusUnauthorized
	}
	return 0
}

func isNotLogged(auth string) bool {

	return auth == ""
}

func extractBearer(authorization string) string {
	var tokens = strings.Split(authorization, " ")
	if len(tokens) == 2 {
		return strings.Trim(tokens[1], " ")
	}
	return ""
}
