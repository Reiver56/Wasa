package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Function that set a new user's nickname
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user User

	path := ps.ByName("id") // get id from url

	userID := ctx.UserID

	if userID != path {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !user.isValidID() {

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.SetNewUsername(path, user.Nickname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
