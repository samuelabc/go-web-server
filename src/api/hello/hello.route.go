package helloRoute

import (
	routeModel "web-server/model/route"
)

func InsertRoute(routeMatrix map[routeModel.Path]routeModel.Route) {
	routeMatrix[routeModel.Path{MainPath: "/hello", SubPath: ""}] =
		routeModel.Route{
			Controller: PostHello,
			Method:     "POST",
		}
}
