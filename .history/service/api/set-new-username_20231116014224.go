package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"net/http"
	"github.com/julienschmidt/httprouter"
)
func(rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	
	id_Path:= ps.ByName("id")
	if id, err := strconv.Atoi(id_Path);err != nil {
		return http.StatusBadRequest, "Invalid ID",nil
		} else if user, err := srv.UserService().GetById(ctx, int64(id));err!=nil{
			return http.StatusInternalServerError,"Failed to get user by id: "+err.Error(),nil
			

}