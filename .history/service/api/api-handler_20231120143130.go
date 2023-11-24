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
	
	//  ------getUserProfile--------------------------------
	rt.router.GET("/users",rt.wrap(rt.searchUser))
	
	//  -----setMyUserName-------------------
	rt.router.PUT("/users/:id",rt.wrap(rt.setMyUserName))


	//  -----likePhoto--------------------------
	rt.router.PUT("/users/:id/like",rt.wrap(rt.likePhoto))

	//  -----banUser--------------------------
	rt.router.PUT("/users/:id/ban",rt.wrap(rt.banUser))

	//  -----unbanUser--------------------------
	rt.router.DELETE("/users/:id/ban/:banned_id",rt.wrap(rt.unbanUser))
	
	return rt.router 
}