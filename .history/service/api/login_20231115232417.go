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
	var user User //   Create a User object to hold the data from the request
	//   Decode JSON from the request body into the User object
	err := json.NewDecoder(r.Body).Decode(&user)
	//   Check for JSON decoding errors
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !user.isValidID(){ //   Validate User ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//  check if user exist and return them, if not exist create a new user 
	exist, err := rt.db.

	
	

}