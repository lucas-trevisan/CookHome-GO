package routes

import (
	"go/src/controllers"
	"net/http"
)

var routesGroups = []Route{
	{
		URI:         "/groups",
		Method:      http.MethodPost,
		Function:    controllers.CreateGroup,
		RequireAuth: false,
	},

	{
		URI:         "/groups",
		Method:      http.MethodGet,
		Function:    controllers.FindGroups,
		RequireAuth: false,
	},

	{
		URI:         "/groups/{groupId}",
		Method:      http.MethodGet,
		Function:    controllers.FindGroupById,
		RequireAuth: false,
	},

	{
		URI:         "/groups/{groupId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateGroup,
		RequireAuth: false,
	},

	{
		URI:         "/groups/{groupId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteGroup,
		RequireAuth: false,
	},
}
