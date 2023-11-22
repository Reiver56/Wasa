package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

//   Function that set a new user's nickname
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	
	banner := ps.ByName("id")
	banned := ps.ByName("banned_id")

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
	//  define banner (with b1) and banned (with b2)
	var b1 User
	var b2 User
	b1.ID = banner
	b2.ID = banned
	
	err := rt.db.(b1.ToDatabase(),b2.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error banning user: %v", err)
		return
	}
	//  aggiungere l'unfollow

	w.WriteHeader(http.StatusNoContent)



	

}