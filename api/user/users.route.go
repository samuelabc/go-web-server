package userController

import (
	routeModel "web-server/model/route"
)

const userMainPath = "/user"

func InsertRoute(routeMatrix map[routeModel.Path]routeModel.Route, rs *UserResource) {
	routeMatrix[routeModel.Path{MainPath: userMainPath, SubPath: "/fetch", Method: routeModel.GET}] =
		routeModel.Route{
			Controller: rs.get,
		}
	routeMatrix[routeModel.Path{MainPath: userMainPath, SubPath: "/register", Method: routeModel.POST}] =
		routeModel.Route{
			Controller: rs.register,
		}
	routeMatrix[routeModel.Path{MainPath: userMainPath, SubPath: "/login", Method: routeModel.POST}] =
		routeModel.Route{
			Controller: rs.login,
		}
	routeMatrix[routeModel.Path{MainPath: userMainPath, SubPath: "/list", Method: routeModel.GET}] =
		routeModel.Route{
			Controller: rs.list,
		}
	routeMatrix[routeModel.Path{MainPath: userMainPath, SubPath: "/update", Method: routeModel.POST}] =
		routeModel.Route{
			Controller: rs.update,
		}
	routeMatrix[routeModel.Path{MainPath: userMainPath, SubPath: "/delete", Method: routeModel.DELETE}] =
		routeModel.Route{
			Controller: rs.delete,
		}
}
