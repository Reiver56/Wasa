package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Endpoints:
	// -----------
	//-----doLogin--------------------------
	rt.router.POST("/login",rt.wrap(rt.doLogin))
	//-----setNewUserName-------------------
	rt.router.PUT("/users/:id",rt.wrap(rt.set))
	//------------- 

	return rt.router 
}