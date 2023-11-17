package api

import 
(
	"net/http"
	"os"
	"path/filepath"
	"encoding/json"
	""

)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){

	w.Header().Set("Content-type", "application/json")
	

}