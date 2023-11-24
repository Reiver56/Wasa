package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"Wasa-photo-1905917/service/database"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id_user := ps.ByName("id") //  take a user's id from the url of profile that we want to see
	req_user := r.Header.Get("Authorization") //  take the id of the user that is logged in from the header

	var followers []database.Profile
	var following []database.Profile
	var photos []database.Photo

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
	//  control if user that requesting the profile is banned from 
}