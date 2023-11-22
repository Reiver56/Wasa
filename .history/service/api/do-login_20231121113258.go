package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"path/filepath"
	"os"

	"github.com/julienschmidt/httprouter"
)

/*
doLogin handles the login request and creates a new user if it doesn't exist.
 It decodes the JSON from the request body into a User object and validates the user ID.
If the user exists, it retrieves the user from the database and returns it.
If the user doesn't exist, it creates a new user and returns it.
*/
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("Content-type", "application/json")
	var user User //   Create a User object to hold the data from the request

	//   Decode JSON from the request body into the User object
	err := json.NewDecoder(r.Body).Decode(&user)

	//   Check for JSON decoding errors
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !user.isValidID() { //   Validate User ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//  Check if user exist and return them, if not exist create a new user

	exist, err := rt.db.ExistUser(user.Nickname)
	if err != nil {

		ctx.Logger.WithError(err).Error("can't check if user exists")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !exist {
		

		err := rt.db.CreateUser(user.ToDatabase())
		if err != nil {
			ctx.Logger.WithError(err).Error("can't create user")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = createFolder(user.ID, ctx)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't create folder")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}


		w.WriteHeader(http.StatusCreated)

	} else {
		databaseUser, err := rt.db.GetUser(user.Nickname)
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
	//   Create a new AuthUser object to hold the user ID and nickname
	type AuthUser struct {
		ID       string `json:"id"`
		nickname string `json:"nickname"`
	}

	authUser := AuthUser{user.ID, user.ID}

	err = json.NewEncoder(w).Encode(authUser)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		ctx.Logger.WithError(err).Error("can't encode user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}



}





func createFolder(id_user string, ctx reqcontext.RequestContext) error {
	path := filepath.Join("photos", id_user)

	err := os.MkdirAll(filepath.Join(path, "photos"))
	if err != nil {
		ctx.Logger.WithError(err).Error("can't create folder")
		return err
	}
	return nil
}
