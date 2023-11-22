package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

//   Function that set a new user's nickname
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
	var banner User
	var banned User
	

	err := rt.db.BanUser(banner.ToDatabase(),banned.ToDatabase())



	

}