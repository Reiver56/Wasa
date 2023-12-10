package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Function that set a new user's nickname
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get user profile and photo_id from the url
	id_user_photo := ps.ByName("id")
	photo_id, err := strconv.ParseInt(ps.ByName("photo_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("photo_id must be an integer")
		return
	}
	id_like_user := ctx.UserID // take the id of the user to like from the header
	// control for existance of the parameter(id) in the header

	// control for existance of the parameter(photo_id)
	if ps.ByName("photo_id") == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("for like a photo you must specify the photo_id")
		return
	}



	var like Like
	like.User = id_like_user
	like.ID_photo = photo_id

	err = rt.db.UnlikePhoto(like.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error banning user: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
