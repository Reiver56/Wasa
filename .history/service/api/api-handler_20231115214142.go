package api

import (
	"net/http"
)

//   Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//   Endpoints:
	//   -----------
	rt.router.POST("/login",)

	return rt.router 
}