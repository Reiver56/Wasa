package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")
	var user User // Create a User object to hold the data from the request
	err := json.NewDecoder(r.Body).Decode(&user.Nickname)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !user.isValid() { // Validate User ID
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user, "   ", err)


}