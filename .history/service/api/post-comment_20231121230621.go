package api


import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"Wasa-photo-1905917/service/database"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	auth := r.Header.Get("Authorization")
	if auth != id {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Errorf("user not logged in")
		return
	}
	banned, err := rt.db.GetFollow(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error getting followers: %v", err)
		return
	}
	var photos []database.Photo
	var followerPhotos []database.Photo
	var follower database.User
	follower.ID = id
	for _, follower := range followers {
		followerPhotos, err = rt.db.GetListPhoto(follower.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.Errorf("error getting photos: %v", err)
			return
		}
		for i, photo := range followerPhotos {
			if i >= PHOTO_IN_HOME{
				break
			}
			photos = append(photos, photo)

		}
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(photos)

}