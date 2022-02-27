package apiMain

import (
	"fmt"
	routeModel "web-server/model/route"
	healthRoute "web-server/src/api/health"
	helloRoute "web-server/src/api/hello"
)

func GetRouteMatrix() map[routeModel.Path]routeModel.Route {
	routeMatrix := make(map[routeModel.Path]routeModel.Route)
	healthRoute.InsertRoute(routeMatrix)
	helloRoute.InsertRoute(routeMatrix)
	fmt.Println("routeMatrix", routeMatrix)
	return routeMatrix
}
