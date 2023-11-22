package api

import (
	"Wasa-photo-1905917/service/api/reqcontext"
	"fmt"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)


func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {