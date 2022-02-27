package healthRoute

import (
	routeModel "web-server/model/route"
)

func InsertRoute(routeMatrix map[routeModel.Path]routeModel.Route) {
	routeMatrix[routeModel.Path{MainPath: "/health", SubPath: ""}] =
		routeModel.Route{
			Controller: PostHealth,
			Validation: GetPostHealthSchemaObject,
			Method:     "POST",
		}
}
