package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	
	"encoding/json"
	"net/http"
	
	"fmt"
	"github.com/julienschmidt/httprouter"
)

//   Function that set a new user's nickname
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	path := ps.ByName("id")
	fmt.Println(path)
	var  user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("1")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !user.isValidID() {
		fmt.Println("2")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.SetNewUsername(path, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return

	

}

