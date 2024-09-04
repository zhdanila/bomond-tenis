package auth

import (
	"bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/authentication"
	"bomond-tenis/internal/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func SignInHandler(service service.Authorization) authentication.PostV1BomondVnAuthSignInHandlerFunc {
	return func(params authentication.PostV1BomondVnAuthSignInParams) middleware.Responder {
		fmt.Println("sign in")

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
