package routers

import "board/controllers"

var userController = controllers.UserAction{}
var chatController = controllers.ChatAction{}
var PusherController = controllers.PusherAction{}

var Apis = []controllers.Controller{
	{Method: "GET", Path: "/", Action: controllers.StatusAction{}.Status},

	// user
	{Method: "GET", Path: "/user/:id", Action: userController.Show},
	{Method: "POST", Path: "/user", Action: userController.Store},

	{Method: "GET", Path: "/chat", Action: chatController.Room},
	{Method: "GET", Path: "/app", Action: PusherController.Sub},
}
