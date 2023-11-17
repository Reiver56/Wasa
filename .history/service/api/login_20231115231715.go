package api

import 
(
	"net/http"
	"os"
	"path/filepath"
	"encoding/json"
	"Wasa-photo-1905917/service/api/reqcontext"

)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){

	w.Header().Set("Content-type", "application/json")
	var user structures.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !isValidID(user.ID){
		w.WriteHeader(http.StatusForbidden)
		return
	}

	
	

}