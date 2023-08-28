package routing

import (
	"blog/pkg/sessions"
	"blog/pkg/static"
)

func RouteBuilder() {
	Init()
	route := GetRouter()
	sessions.Init(route)
	static.LoadStatic(route)
	RegisterRoutes(route)
	Serve(route)

}
