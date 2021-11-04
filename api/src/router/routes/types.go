package routes

import (
	"go/src/controllers"
	"net/http"
)

var routesTypes = []Route{
	{
		URI:         "/types",
		Method:      http.MethodPost,
		Function:    controllers.CreateType,
		RequireAuth: false,
	},

	{
		URI:         "/types",
		Method:      http.MethodGet,
		Function:    controllers.FindTypes,
		RequireAuth: false,
	},

	{
		URI:         "/types/{typeId}",
		Method:      http.MethodGet,
		Function:    controllers.FindTypeById,
		RequireAuth: false,
	},

	{
		URI:         "/types/{typeId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateType,
		RequireAuth: false,
	},

	{
		URI:         "/types/{typeId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteType,
		RequireAuth: false,
	},
}
