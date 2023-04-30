package routers

import "board/controllers"

var Apis = []controllers.Controller{
	{Method: "GET", Path: "/", Action: controllers.StatusAction{}.Status},
	{Method: "POST", Path: "/user", Action: controllers.UserAction{}.Store},
}
