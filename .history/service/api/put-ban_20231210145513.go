package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Function that set a new user's nickname
func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	banner := ps.ByName("id")
	control := ctx.UserID
	banned := ps.ByName("banned_id")

	if banner != control {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user not logged in")
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

	w.WriteHeader(http.StatusNoContent)

}
