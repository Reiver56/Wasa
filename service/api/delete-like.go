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

	targetUserID := ps.ByName("like_id")
	if targetUserID != id_like_user {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// control for existance of the parameter(photo_id)
	if ps.ByName("photo_id") == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("for like a photo you must specify the photo_id")
		return
	}

	if rt.db.IsBanned(id_user_photo, id_like_user) {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Errorf("user is banned")
		return
	}

	var like Like
	like.User = targetUserID
	like.ID_photo = photo_id

	err = rt.db.UnlikePhoto(like.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error banning user: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
