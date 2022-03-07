package accountController

import (
	routeModel "web-server/model/route"
)

const accountMainPath = "/account"

func InsertRoute(routeMatrix map[routeModel.Path]routeModel.Route, rs *AccountResource) {
	routeMatrix[routeModel.Path{MainPath: accountMainPath, SubPath: "/fetch", Method: routeModel.GET}] =
		routeModel.Route{
			Controller: rs.get,
		}
	routeMatrix[routeModel.Path{MainPath: accountMainPath, SubPath: "/register", Method: routeModel.POST}] =
		routeModel.Route{
			Controller: rs.register,
		}
	routeMatrix[routeModel.Path{MainPath: accountMainPath, SubPath: "/login", Method: routeModel.POST}] =
		routeModel.Route{
			Controller: rs.login,
		}
	routeMatrix[routeModel.Path{MainPath: accountMainPath, SubPath: "/list", Method: routeModel.GET}] =
		routeModel.Route{
			Controller: rs.list,
		}
	routeMatrix[routeModel.Path{MainPath: accountMainPath, SubPath: "/update", Method: routeModel.POST}] =
		routeModel.Route{
			Controller: rs.update,
		}
	routeMatrix[routeModel.Path{MainPath: accountMainPath, SubPath: "/delete", Method: routeModel.DELETE}] =
		routeModel.Route{
			Controller: rs.delete,
		}
}
