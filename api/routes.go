package api

import (
	"fmt"
	articleRoute "web-server/api/article"
	helloRoute "web-server/api/hello"
	userRoute "web-server/api/user"
	routeModel "web-server/model/route"
)

func GetRouteMatrix(appAPI *API) map[routeModel.Path]routeModel.Route {
	routeMatrix := make(map[routeModel.Path]routeModel.Route)
	helloRoute.InsertRoute(routeMatrix)
	articleRoute.InsertRoute(routeMatrix, appAPI.Article)
	userRoute.InsertRoute(routeMatrix, appAPI.User)
	fmt.Println("routeMatrix", routeMatrix)
	return routeMatrix
}
