package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"fmt"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

// Function that set a new user's nickname
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	
	id_user_photo := ps.ByName("id")
	//control for existance of the parameter(id) in the url
	if id_user_photo == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user not logged in")
		return
	}

	id_like_user := r.Header.Get("Authorization") //take the id of the user to like from the header
	//control for existance of the parameter(id) in the header 
	if id_like_user == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user not logged in")
		return
	}
	//control for existance of the parameter(photo_id)
	if ps.ByName("photo_id") == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("for like a photo you must specify the photo_id")
		return
	}
	// this  control for photo_id, that must be an integer
	photo_id, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("photo_id must be an integer")
		return
	}

	fmt.Println(id_like_user+" "+id_user_photo)
	

	
	var user_photo User
	user_photo.ID = id_user_photo
	var user_like User
	user_like.ID = id_like_user
	err = rt.db.ULikePhoto(photo_id, user_like.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error banning user: %v", err)
		return
	}
	
	
	
	
	
	
	//aggiungere l'unfollow

	w.WriteHeader(http.StatusNoContent)



	

}
	



	

