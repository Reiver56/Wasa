package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Endpoints:
	// -----------
	//-----ogin--------
	rt.router.POST("/login",rt.wrap(rt.doLogin))
	//-----setNewUserN
	//------------- 

	return rt.router 
}