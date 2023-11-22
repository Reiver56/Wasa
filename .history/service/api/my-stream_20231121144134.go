package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) myStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	w.Header().Set("Content-Type", "application/json")
	id
}