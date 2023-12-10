package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//   Endpoints:
	//   -----------

	//  -----doLogin--------------------------
	rt.router.POST("/session", rt.wrap(rt.doLogin, false))

	//  ------searchUser--------------------------------
	rt.router.GET("/users", rt.wrap(rt.searchUser,true))

	//  -----setMyUserName-------------------
	rt.router.PUT("/users/:id", rt.wrap(rt.setMyUserName, true))
	//  -----getProfile--------------------------
	rt.router.GET("/users/:id", rt.wrap(rt.getProfile, true))

	//  -----myStream--------------------------
	rt.router.GET("/users/:id/mystream", rt.wrap(rt.myStream , true))

	//  -----uploadPhoto--------------------------
	rt.router.POST("/users/:id/photos", rt.wrap(rt.uploadPhoto, true))
	//  -----deletePhoto--------------------------
	rt.router.DELETE("/users/:id/photos/:photo_id", rt.wrap(rt.deletePhoto, true))

	//  -----commentPhoto--------------------------
	rt.router.POST("/users/:id/photos/:photo_id/comments", rt.wrap(rt.commentPhoto , true))
	rt.router.DELETE("/users/:id/photos/:photo_id/comments/:comment_id", rt.wrap(rt.unCommentPhoto, true))

	//  -----likePhoto--------------------------
	rt.router.PUT("/users/:id/photos/:photo_id/likes/:like_id", rt.wrap(rt.likePhoto, true))
	rt.router.DELETE("/users/:id/photos/:photo_id/likes/:like_id", rt.wrap(rt.unlikePhoto, true))

	//  -----banUser--------------------------
	rt.router.PUT("/users/:id/ban_users/:banned_id", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:id/ban_users/:banned_id", rt.wrap(rt.unbanUser))

	//  -----followUser--------------------------
	rt.router.PUT("/users/:id/followers/:follow_id", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:id/followers/:follow_id", rt.wrap(rt.unfollowUser))

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
