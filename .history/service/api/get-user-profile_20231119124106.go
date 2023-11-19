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
	
	if identificator == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	identifier := ps.ByName("id")

	var user1 User
	var user2 User
	user1.ID = identifier
	user2.ID = identificator

	fmt.Println()
	
	profiles, err := rt.db.GetUserProfile(user1.ToDatabase(),user2.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error getting user profile: %v", err)
		_ = json.NewEncoder(w).Encode([]Profile{})
		return
	}
	w.WriteHeader(http.StatusOK)
if len(profiles) == 0 {
		_ = json.NewEncoder(w).Encode([]Profile{})
		return
	}
	


}