package api


import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//   Function that set a new user's nickname
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	
	id_user_photo := ps.ByName("id")
	id_like_user := r.Header.Get("Authorization") //  take the id of the user to like from the header
	photo_id := ps.ByName("photo_id")
	fmt.Println(id_like_user+" "+id_user_photo)

	if id_user_photo == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user not logged in")
		return
	}

	if photo_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("for like a photo you must specify the photo_id")
		return
	}
	err := rt.db.LikePhoto(id_user_photo,id_like_user,photo_id)
	
	
	
	
	
	
	
	//  aggiungere l'unfollow

	w.WriteHeader(http.StatusNoContent)



	

}
	fmt.Println(banner+" "+banned)

	if banner == "" {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user not logged in")
		return
	}

	if banner == banned {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Errorf("user cannot ban himself")
		return
	}
	//  define banner (with b1) and banned (with b2)
	var b1 User
	var b2 User
	b1.ID = banner
	b2.ID = banned
	
	err := rt.db.BanUser(b1.ToDatabase(),b2.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error banning user: %v", err)
		return
	}
	//  aggiungere l'unfollow

	w.WriteHeader(http.StatusNoContent)



	

