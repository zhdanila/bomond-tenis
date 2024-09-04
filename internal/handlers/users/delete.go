package users

import (
	"bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/users"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func DeleteUserHandler() users.DeleteV1BomondVnUsersUserIDHandlerFunc {
	return func(params users.DeleteV1BomondVnUsersUserIDParams) middleware.Responder {
		fmt.Println("delete")

		// Приклад: повертаємо просту відповідь
		return users.NewDeleteV1BomondVnUsersUserIDOK().WithPayload(&models.SuccessResponse{
			Code:      "200",
			Data:      "nothing",
			Err:       "nothing",
			Message:   "hello",
			Status:    200,
			Timestamp: strfmt.DateTime{},
		})
	}
}
