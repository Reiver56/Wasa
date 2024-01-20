package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var user User // Create a User object to hold the data from the request

	// Decode JSON from the request body into the User object
	err := json.NewDecoder(r.Body).Decode(&user)

	// Check for JSON decoding errors
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !user.isValidID() { // Validate User ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if user exist and return them, if not exist create a new user

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
		user.ID = user.Nickname
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
	// Struct that contain User object and authorization token
	type AuthUser struct {
		User  User   `json:"user"`
		Token string `json:"token"`
	}
	// in this case the token is the user id
	authUser := AuthUser{user, user.ID}
	
	// encode the AuthUser object in JSON and send it to the client
	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(authUser); err != nil {
		ctx.Logger.WithError(err).Error("can't encode user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
