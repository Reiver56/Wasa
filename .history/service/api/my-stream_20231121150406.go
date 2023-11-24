package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"Wasa-photo-1905917/service/database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) myStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	auth := r.Header.Get("Authorization")
	if auth != id {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Errorf("user not logged in")
		return
	}
	followers, err := rt.db.GetFollow(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error getting followers: %v", err)
		return
	}
	var photos []database.Photo
	for _, follower := range followers {
		followerphotos, err = rt.db.GetListPhoto(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.Errorf("error getting photos: %v", err)
			return
		}
		for i, photo := range followerphotos {
			if i >=

		}
	}

}