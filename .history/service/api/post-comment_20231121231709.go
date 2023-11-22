package api


import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"Wasa-photo-1905917/service/database"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	w.Header().Set("Content-Type", "application/json")
	
	auth := r.Header.Get("Authorization")
	
	var user_comment User
	user_comment.ID = auth
	var user_photo User
	user_photo.ID = ps.ByName("photo_id")
	err,value := rt.db.IsBanned(user_comment.ID, user_photo.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error getting followers: %v", err)
		return
	}
	if value == true{
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Errorf("user banned")
		return
	}
	var comment Comment
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("error decoding request body: %v", err)
		return
	}
	if len(comment.Text_comment) == 0 && len() {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("error decoding request body: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(photos)

}