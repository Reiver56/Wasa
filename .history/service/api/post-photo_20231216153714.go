package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"bytes"
	"encoding/json"
	"errors"
	"image/jpeg"
	"image/png"
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
		ctx.Logger.WithError(err).Error("error reading body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))

	// check if the body content is valid png or jpeg img
	err = checkFormatPhoto(r.Body, io.NopCloser(bytes.NewBuffer(data)), ctx)
	if err != nil {
		ctx.Logger.WithError(err).Error("invalid image format")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))

	// create a new photo in database with generated id for photo
	id_photo, err := rt.db.UploadPhoto(photo.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("error uploading photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id_photo := strconv.FormatInt(id_photo, 10)





}

func checkFormatPhoto(body io.ReadCloser, newReader io.ReadCloser, ctx reqcontext.RequestContext) error {
	// check if the body content is valid png or jpeg img
	_, errJpg := jpeg.Decode(body)
	if errJpg != nil {
		body = newReader
		_, errPng := png.Decode(body)
		if errPng != nil {
			return errors.New("invalid image format")
		}
		return nil
	}
	return nil
}
