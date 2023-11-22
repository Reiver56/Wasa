package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unCommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	w.Header().Set("Content-Type", "application/json")
	
	auth := r.Header.Get("Authorization")
	
	var user_comment User
	user_comment.ID = auth
	var user_photo User
	user_photo.ID = ps.ByName("id")
	id_photo := ps.ByName("photo_id")
	id_photo_int, err := strconv.ParseInt(id_photo, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error parsing id_photo: %v", err)
		return
	}
	

}