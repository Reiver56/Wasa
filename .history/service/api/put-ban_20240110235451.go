package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	banner := ps.ByName("id")
	control := ctx.UserID
	banned := ps.ByName("banned_id")

	if banner != control {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// define banner (with b1) and banned (with b2)
	var b1 User
	var b2 User
	b1.ID = banner
	b2.ID = banned

	err := rt.db.BanUser(b1.ToDatabase(), b2.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error banning user: %v", err)
		return
	}


	// Ban implement unfollow from banned to banner (if exists)
	err = rt.db.Unfollow(b1.ID, b2.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error unfollowing user: %v", err)
		return
	}
	// Ban implement unfollow from bannedr to banned (if exists)
	err = rt.db.Unfollow(b2.ID, b1.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error unfollowing user: %v", err)
		return
	}
	// Response 204 
	w.WriteHeader(http.StatusNoContent)

}
