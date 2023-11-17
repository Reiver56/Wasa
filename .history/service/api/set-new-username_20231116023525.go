package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"strconv"
	"net/http"
	"github.com/julienschmidt/httprouter"
)
func(rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	
	id_Path, err := strconv.Unquote(ps.ByName("id"))
	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	//get new username
	var user _User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	err = rt.db.SetNewUsername(User_ID{ID: id_Path}.ToDatabase(),user.ToDatabase())
	if err!=nil{
		ctx.Logger.WithError(err).
		return
	}	
	

}