package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getBans(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id_user := ps.ByName("id")
	userID := ctx.UserID
	if id_user != userID {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Errorf("error getting ban list: %v", "you can't see the bans of another user")
		return
	}

	// get user profile bans from the database
	dbBans, err := rt.db.GetBans(id_user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error getting bans: %v", err)
		return
	}
	bans := make([]User, len(dbBans))
	for i, ban := range dbBans {
		var user User
		err := user.FromDatabase(ban)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.Errorf("error getting bans: %v", err)
			return
		}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// write followers in the response
	if err := json.NewEncoder(w).Encode(bans); err != nil {
		ctx.Logger.Errorf("error writing response: %v", err)
	}

}
