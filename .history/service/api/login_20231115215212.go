package api

import 
(
	"fmt"

)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	type Api struct {
		Name string `json:"name"`
		}
		
}