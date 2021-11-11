package routes

import (
	"go/src/controllers"
	"net/http"
)

var routesPortions = []Route{
	{
		URI:         "/portions",
		Method:      http.MethodPost,
		Function:    controllers.CreatePortion,
		RequireAuth: false,
	},

	{
		URI:         "/portions",
		Method:      http.MethodGet,
		Function:    controllers.FindPortions,
		RequireAuth: false,
	},

	{
		URI:         "/portions/{portionId}",
		Method:      http.MethodGet,
		Function:    controllers.FindPortionById,
		RequireAuth: false,
	},

	{
		URI:         "/portions/{portionId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePortion,
		RequireAuth: false,
	},

	{
		URI:         "/portions/{portionId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePortion,
		RequireAuth: false,
	},
}
