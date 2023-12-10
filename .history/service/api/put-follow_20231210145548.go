package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	auth := ctx.UserID
	// verify if user is logged in
	if auth != ps.ByName("id") {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("error: user is not logged in")
		return
	}

	// define two users: user_req is the user that is logged in and user is the user that is going to be followed
	var user_req User
	user_req.ID = ps.ByName("id")
	var user User
	user.ID = ps.ByName("follow_id")

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

	err := rt.db.Follow(user_req.ID, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error following user: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)

}
