package api

import (
	"net/http"
)

//   Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//   Endpoints:
	//   -----------

	//  -----doLogin--------------------------
	rt.router.POST("/login", rt.wrap(rt.doLogin))

	//  ------searchUser--------------------------------
	rt.router.GET("/users", rt.wrap(rt.searchUser))

	//  -----setMyUserName-------------------
	rt.router.PUT("/users/:id", rt.wrap(rt.setMyUserName))
	//  -----getProfile--------------------------
	rt.router.GET("/users/:id", rt.wrap(rt.getProfile))

	//  -----myStream--------------------------
	rt.router.GET("/users/:id/mystream", rt.wrap(rt.myStream))

	//  -----uploadPhoto--------------------------
	rt.router.POST("/users/:id/photos", rt.wrap(rt.uploadPhoto))
	//  -----deletePhoto--------------------------
	rt.router.DELETE("/users/:id/photos/:photo_id", rt.wrap(rt.deletePhoto))

	//  -----commentPhoto--------------------------
	rt.router.POST("/users/:id//comment", rt.wrap(rt.commentPhoto))
	//  -----deleteComment--------------------------
	rt.router.DELETE("/users/:id/comment/:comment_id", rt.wrap(rt.unCommentPhoto))

	//  -----likePhoto--------------------------
	rt.router.PUT("/users/:id/like", rt.wrap(rt.likePhoto))
	//  -----unlikePhoto--------------------------
	rt.router.DELETE("/users/:id/like/:photo_id", rt.wrap(rt.unlikePhoto))

	//  -----banUser--------------------------
	rt.router.PUT("/users/:id/ban", rt.wrap(rt.banUser))
	//  -----unbanUser--------------------------
	rt.router.DELETE("/users/:id/ban/:banned_id", rt.wrap(rt.unbanUser))

	//  -----followUser--------------------------
	rt.router.PUT("/users/:id/follow", rt.wrap(rt.followUser))
	//  -----unfollowUser--------------------------
	rt.router.DELETE("/users/:id/follow/:followed_id", rt.wrap(rt.unfollowUser))

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
