package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")

	profiles, err := rt.db.GetUserProfile( User{ID: ps.ByName("id"),nickname: }, User{Nickname: r.URL.Query().Get("nickname")})
	


}