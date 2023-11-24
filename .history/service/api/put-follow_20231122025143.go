package api
import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	w.Header().Set("Content-Type", "application/json")
	
	auth := r.Header.Get("Authorization")
	
	var user_req User
	user_req.ID = auth 
	var user User
	user.ID = ps.ByName("id")

	err, value := rt.db.IsBanned(user_req.ID, user.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error checking if user is banned: %v", err)
		return
	}
	if value == true {
		w.WriteHeader(http.StatusForbidden)
		ctx.Logger.Errorf("user is banned")
		return
	}

	err = rt.db.FollowUser(user_req.ID, user.ID)
	
	

}