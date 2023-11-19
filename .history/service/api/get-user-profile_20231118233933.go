package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")
	identificator := r.URL.Query().Get("id")
	identifier := ps.ByName("id")
	if identifier == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	profiles, err := rt.db.GetUserProfile(User{ID: identifier}.ToDatabase(), User{ID: identifier}.ToDatabase())
	


}