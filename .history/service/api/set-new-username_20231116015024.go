package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"strconv"
	"net/http"
	"github.com/julienschmidt/httprouter"
)
func(rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	
	id_Path, err := strconv.Atoi(ps.ByName("id"))
	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	//  get new username
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	

}