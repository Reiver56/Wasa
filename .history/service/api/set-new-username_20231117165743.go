package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// Function that set a new user's nickname
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	path := ps.ByName("id")
	validate := validateRequestingUser(path, extractBearer(r.Header.Get("Authorization")))
	if validate != 0 {
		w.WriteHeader(validate)
		return
	}

	var nickname nickname
	err := json.NewDecoder(r.Body).Decode(&nickname)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !nickname.isValidNick() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.SetNewUsername(dat, nickname.nickname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return

	

}

func validateRequestingUser(identifier string, Token string) int {

	
	if isNotLogged(bearerToken) {
		return http.StatusForbidden
	}

	

	if identifier != bearerToken {
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
