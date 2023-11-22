package api

import (
	"net/http"
)

//   Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//   Endpoints:
	//   -----------
	//  -----doLogin--------------------------
	rt.router.POST("/login",rt.wrap(rt.doLogin))
	
	//  -----setMyUserName-------------------
	rt.router.PUT("/users/:id",rt.wrap(rt.setMyUserName))
	//  ------getUserProfile--------------------------------
	rt.router.GET("/users",rt.wrap(rt.getUserProfile))
	return rt.router 
}