package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/julienschmidt/httprouter"
)

//   Function that set a new user's nickname
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	auth := r.Header.Get("Authorization")
	id_user_photo := ps.ByName("id")
	
	if auth != id_user_photo {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Errorf("user not logged in")
		return
	}
	var photo Photo

	photo.

	


	

}
	



	

