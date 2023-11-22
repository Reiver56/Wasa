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
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	//   Check for JSON decoding errors
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !isValidID(user.ID){ //   Validate User ID
		w.WriteHeader(http.StatusForbidden)
		return
	}


	
	

}