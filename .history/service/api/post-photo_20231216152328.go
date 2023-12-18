package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Function that set a new user's nickname
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	auth := ctx.UserID
	id_user_photo := ps.ByName("id")

	if auth != id_user_photo {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// define the photo
	var photo Photo

	photo.User_ID = id_user_photo
	photo.Date = time.Now().UTC()
	// -----------------

	// create copy of body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger
		w.WriteHeader(http.StatusInternalServerError)
		return
	}




}
