package auth

import (
	"bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/authentication"
	"bomond-tenis/internal/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"time"
)

func SignInHandler(authService service.Authorization) authentication.PostV1BomondVnAuthSignInHandlerFunc {
	return func(params authentication.PostV1BomondVnAuthSignInParams) middleware.Responder {
		token, err := authService.SignIn(params.UserSignin.Email, params.UserSignin.Password)
		if err != nil {
			return authentication.NewPostV1BomondVnAuthSignInBadRequest().WithPayload(&models.ErrorResult{
				Code:      "500",
				DebugInfo: err.Error(),
				Message:   "sign in problem",
				Status:    500,
				Timestamp: strfmt.DateTime(time.Now().UTC()),
			})
		}

		return authentication.NewPostV1BomondVnAuthSignInOK().WithPayload(&models.SuccessResponse{
			Code:      "200",
			Data:      token,
			Err:       "nothing",
			Message:   "hello",
			Status:    200,
			Timestamp: strfmt.DateTime{},
		})
	}
}
