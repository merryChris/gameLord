package api

import (
	"github.com/gorilla/mux"
	"github.com/merryChris/gameLord/types"
)

var v1Routes = types.Routes{
	types.Route{
		"POST",
		"Signup",
		"/v1/signup",
		RoutingUserHandler.Signup,
	},
	types.Route{
		"POST",
		"Login",
		"/v1/login",
		RoutingUserHandler.Login,
	},
	types.Route{
		"POST",
		"Logout",
		"/v1/logout",
		RoutingUserHandler.Logout,
	},
	types.Route{
		"POST",
		"LoadGame",
		"/v1/load_game",
		RoutingUserHandler.LoadGame,
	},
	types.Route{
		"POST",
		"LeaveGame",
		"/v1/leave_game",
		RoutingUserHandler.LeaveGame,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range v1Routes {
		handlerFunc := route.HandlerFunc
		handlerFunc = HandlerWrapper(handlerFunc, route.Name)

		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			HandlerFunc(handlerFunc)
	}
	return router
}

func AddRouter(router *mux.Router, routesList ...types.Routes) {
	for _, routes := range routesList {
		for _, route := range routes {
			handlerFunc := route.HandlerFunc
			handlerFunc = HandlerWrapper(handlerFunc, route.Name)

			router.
				Methods(route.Method).
				Name(route.Name).
				Path(route.Pattern).
				HandlerFunc(handlerFunc)
		}
	}
}
