package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")
	identificator := r.URL.Query().Get("search")

	userID := ctx.UserID

	if userID == "" {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Errorf("user not logged in")
		return
	}

	var user1 User
	var user2 User
	user1.ID = identifier
	user2.ID = identificator

	profiles, err := rt.db.GetUserProfile(user1.ToDatabase(), user2.ToDatabase())
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
	_ = json.NewEncoder(w).Encode(profiles)

}
