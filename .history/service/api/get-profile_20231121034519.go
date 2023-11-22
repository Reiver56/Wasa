package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"Wasa-photo-1905917/service/database"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id_user := ps.ByName("id")                //  take a user's id from the url of profile that we want to see
	req_user := r.Header.Get("Authorization") //  take the id of the user that is logged in from the header

	var followers []database.User
	var following []database.User
	var photos []database.Photo
	var nickname string
	
	//  control if user is logged in
	if req_user == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user not logged in")
		return
	}
	//  control if id_user is empty
	if id_user == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user not logged in")
		return
	}

	//  control if user is banned
	if _, banned := rt.db.IsBanned(id_user, req_user); banned == true {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Errorf("user is banned")
		return
	}

	//  control if user exists
	if _, err := rt.db.ExistUser(id_user); err != nil {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.Errorf("user not found")
		return
	}

	//  control if user requested is user banned
	if _, banned := rt.db.IsBanned(req_user, id_user); banned == true {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Errorf("user is banned")
		return
	}
	
	

	followers, _ = rt.db.GetFollow(id_user)
	following, _ = rt.db.GetFollowing(id_user)
	photos, _ = rt.db.GetListPhoto(id_user)
	nickname, _ = rt.db.GetNickname(id_user)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(P)
	
	

	//  control if user that requesting the profile is banned from the user that owns the profile

}
