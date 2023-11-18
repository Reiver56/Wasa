package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"net/http"
	
	

	"github.com/julienschmidt/httprouter"
)

// Function that set a new user's nickname
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	path := ps.ByName("id")
	validate := 
	
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !user.isValidID() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.SetNewUsername(id, user.Nickname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return

	

}

func validateRequestingUser(identifier string, bearerToken string) int {

	// If the requesting user has an invalid token then respond with a fobidden status
	if isNotLogged(bearerToken) {
		return http.StatusForbidden
	}

	//  If the requesting user's id is different than the one in the path then respond with a unathorized status.

	if identifier != bearerToken {
		return http.StatusUnauthorized
	}
	return 0
}
