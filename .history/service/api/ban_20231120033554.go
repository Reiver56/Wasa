package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// Function that set a new user's nickname
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	
	banner := ps.ByName("id")
	banned := r.Header.Get("Authorization")

	if banner == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user not logged in")
		return
	}
	if banner == banned {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user cannot ban himself")
		return
	}
	//
	var b1 User
	var b2 User
	b1.ID = banner
	b2.ID = banned

	err := rt.db.BanUser(b1.ToDatabase(),b2.ToDatabase())



	

}