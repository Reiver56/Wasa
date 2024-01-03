package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"

	"Wasa-photo-1905917/service/api/reqcontext"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	http.ServeFile(w, r,
		filepath.Join(photoFolder, ps.ByName("id"), "photos", ps.ByName("photo_id")))
}
