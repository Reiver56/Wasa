package api
import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	w.Header().Set("Content-Type", "application/json")

	auth := r.Header.Get("Authorization")

	if auth != ps.ByName("id") {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("errr")
		return
	}

	var user_req User
	user_req.ID = auth
	var user User
	user.ID = ps.ByName("follower_id")

	if user.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("error getting user id")
		return
	}

	if user_req.ID == user.ID {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("error following user")
		return
	}
	if user_req.ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("error getting user id")
		return
	}

	if rt.db.IsBanned(user_req.ID, user.ID){
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Errorf("user is banned")
		return
	}

	err := rt.db.Follow(user_req.ID, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error following user: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)



}
