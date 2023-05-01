package routers

import "board/controllers"

var userController = controllers.UserAction{}

var Apis = []controllers.Controller{
	{Method: "GET", Path: "/", Action: controllers.StatusAction{}.Status},

	// user
	{Method: "GET", Path: "/user/:id", Action: userController.Show},
	{Method: "POST", Path: "/user", Action: userController.Store},
}
