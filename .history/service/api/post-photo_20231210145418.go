package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"bytes"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

	photo_id, err := rt.db.UploadPhoto(photo.ToDatabase())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error uploading photo: %v", err)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error reading body: %v", err)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))

	path, err := getUserFolder(id_user_photo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error getting user folder: %v", err)
		return
	}
	out, err := os.Create(filepath.Join(path, strconv.FormatInt(photo_id, 10)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error creating file: %v", err)
		return
	}
	_, err = io.Copy(out, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("error copying file: %v", err)
		return
	}

	out.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Errorf("Error uploading photo: %v", err)
		return
	}
	photo.Id_photo = photo_id
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)

}