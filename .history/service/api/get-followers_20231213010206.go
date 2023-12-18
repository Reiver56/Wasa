package api
import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"net/http"

	"github.com/julienschmidt/httprouter"
)
func (rt *_router) getMyFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get user profile followers from the url
	id_user := ps.ByName("id")

	// get user profile followers from the database
	followers, err := rt.db.GetFollow(id_user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error getting followers: %v", err)
		return
	}
}
