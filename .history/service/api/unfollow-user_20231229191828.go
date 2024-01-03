package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	identifier := ps.ByName("id")
	auth := ctx.UserID



	var user_req User
	user_req.ID = identifier
	var user User
	user.ID = ps.ByName("follow_id")

	if user.ID != auth {
		w.WriteHeader(http.StatusForbidden)
		return
	}

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

	if rt.db.IsBanned(user_req.ID, user.ID) {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Errorf("user is banned")
		return
	}

	err := rt.db.Unfollow(user_req.ID, user.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error unfollowing user: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)

}
