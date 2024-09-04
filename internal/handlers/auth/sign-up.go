package auth

import (
	"bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/authentication"
	"bomond-tenis/internal/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func SignUpHandler(service service.Authorization) authentication.PostV1BomondVnAuthSignUpHandlerFunc {
	return func(params authentication.PostV1BomondVnAuthSignUpParams) middleware.Responder {
		fmt.Println("sign up")

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
