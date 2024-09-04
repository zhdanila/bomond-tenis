package auth

import (
	"bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/authentication"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func LogoutHandler() authentication.PostV1BomondVnAuthLogoutHandlerFunc {
	return func(params authentication.PostV1BomondVnAuthLogoutParams) middleware.Responder {
		fmt.Println("logout")

		return authentication.NewPostV1BomondVnAuthSignInOK().WithPayload(&models.SuccessResponse{
			Code:      "200",
			Data:      "nothing",
			Err:       "nothing",
			Message:   "hello",
			Status:    200,
			Timestamp: strfmt.DateTime{},
		})
	}
}
