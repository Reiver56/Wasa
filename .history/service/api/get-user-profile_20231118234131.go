package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")
	identificator := r.URL.Query().Get("id")
	
	if identificator == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	identifier := ps.ByName("id")
	
	
	profiles, err := rt.db.GetUserProfile(User{ID: identifier}.ToDatabase(), User{ID: identificator}.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error getting user profile: %v", err)
		_ = json.NewEncoder(w).Encode([]Profile)
		return
	}

	


}