import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	w.Header().Set("Content-Type", "application/json")
	
	auth := r.Header.Get("Authorization")
	
	var user User
	user = auth
	var user User
	user_photo.ID = ps.ByName("id")
	
	

}