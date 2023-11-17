package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"Wasa-photo-1905917/service/database"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){

	w.Header().Set("Content-type", "application/json")
	var user User // Create a User object to hold the data from the request
	
	// Decode JSON from the request body into the User object
	err := json.NewDecoder(r.Body).Decode(&user)
	
	// Check for JSON decoding errors
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !user.isValidID(){ // Validate User ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Check if user exist and return them, if not exist create a new user 
	
	exist, err := rt.db.ExistUser(user.ID,user.Password)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't check if user exists")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//if not exists
	if !exist {
		user, err = rt.db.CreateUser(user.ID, user.)
		if err != nil{
			ctx.Logger.WithError(err).Error("can't create user")
			w.WriteHeader(http.StatusInternalServerError)
			return
	} else {
		databaseUser, err := rt.db.GetUser(user.ID, user.Password)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't get user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = user.FromDatabase(databaseUser)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't convert database user to service user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	}

}
}