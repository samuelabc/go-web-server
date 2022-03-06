package articleController

import (
	routeModel "web-server/model/route"
)

func InsertRoute(routeMatrix map[routeModel.Path]routeModel.Route, rs *ArticleResource) {
	routeMatrix[routeModel.Path{MainPath: "/article", SubPath: "/fetch", Method: routeModel.GET}] =
		routeModel.Route{
			Controller: rs.get,
		}
	routeMatrix[routeModel.Path{MainPath: "/article", SubPath: "/create", Method: routeModel.POST}] =
		routeModel.Route{
			Controller: rs.create,
		}
	routeMatrix[routeModel.Path{MainPath: "/article", SubPath: "/list", Method: routeModel.GET}] =
		routeModel.Route{
			Controller: rs.list,
		}
	routeMatrix[routeModel.Path{MainPath: "/article", SubPath: "/update", Method: routeModel.POST}] =
		routeModel.Route{
			Controller: rs.update,
		}
	routeMatrix[routeModel.Path{MainPath: "/article", SubPath: "/delete", Method: routeModel.DELETE}] =
		routeModel.Route{
			Controller: rs.delete,
		}
}
