package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Function that set a new user's nickname
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	banner := ps.ByName("id")
	banned := ps.ByName("banned_id")

	userID := ctx.UserID

	if banner != userID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// define banner (with b1) and banned (with b2)
	var b1 User
	var b2 User
	b1.ID = banner
	b2.ID = banned

	err := rt.db.UnbanUser(b1.ToDatabase(), b2.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error unbanning user: %v", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
