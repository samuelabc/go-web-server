package articleController

import (
	routeModel "web-server/model/route"
)

func InsertRoute(routeMatrix map[routeModel.Path]routeModel.Route, rs *ArticleResource) {
	routeMatrix[routeModel.Path{MainPath: "/article", SubPath: "", Method: routeModel.GET}] =
		routeModel.Route{
			Controller: rs.get,
		}
	routeMatrix[routeModel.Path{MainPath: "/article", SubPath: "", Method: routeModel.POST}] =
		routeModel.Route{
			Controller: rs.create,
		}
	routeMatrix[routeModel.Path{MainPath: "/article", SubPath: "/list", Method: routeModel.GET}] =
		routeModel.Route{
			Controller: rs.list,
		}
}
