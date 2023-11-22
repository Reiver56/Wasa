package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"	
	"net/http"
	"github.com/julienschmidt/httprouter"
)

//   Function that set a new user's nickname
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	auth := r.Header.Get("Authorization")
	id_user_photo := ps.ByName("id")
	if auth != {
	id_photo := ps.ByName("photo_id")

	
	
}
	



	

