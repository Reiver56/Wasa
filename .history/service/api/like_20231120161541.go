package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//   Function that set a new user's nickname
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	
	id_user_photo := ps.ByName("id")
	id_like_user := r.Header.Get("Authorization") //  take the id of the user to like from the header
	
	//   this  control for photo
	if ps.ByName("photo_id") == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("for like a photo you must specify the photo_id")
		return
	}
	photo_id, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("photo_id must be an integer")
		return
	}

	fmt.Println(id_like_user+" "+id_user_photo)

	if id_user_photo == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user not logged in")
		return
	}

	
	var user_photo User
	user_photo.ID = id_user_photo
	var user_like User
	user_like.ID = id_like_user
	err := rt.db.LikePhoto(photo_id, user_like.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error banning user: %v", err)
		return
	}
	
	
	
	
	
	
	//  aggiungere l'unfollow

	w.WriteHeader(http.StatusNoContent)



	

}
	



	

