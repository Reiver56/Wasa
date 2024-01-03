package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getBans(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get user profile followers from the url
	id_user := ps.ByName("id")
	userID := ctx.UserID()
	

	// get user profile followers from the database
	followers, err := rt.db.GetBans(id_user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error getting followers: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// write followers in the response
	if err := json.NewEncoder(w).Encode(followers); err != nil {
		ctx.Logger.Errorf("error writing response: %v", err)
	}

}
