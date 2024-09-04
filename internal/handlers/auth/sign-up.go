package auth

import (
	"bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/authentication"
	"bomond-tenis/internal/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"time"
)

func SignUpHandler(service service.Authorization) authentication.PostV1BomondVnAuthSignUpHandlerFunc {
	return func(params authentication.PostV1BomondVnAuthSignUpParams) middleware.Responder {
		id, err := service.SignUp(params.UserSignup.Username, params.UserSignup.Email, params.UserSignup.Password)
		if err != nil {
			return authentication.NewPostV1BomondVnAuthSignInBadRequest().WithPayload(&models.ErrorResult{
				Code:      "500",
				DebugInfo: err.Error(),
				Message:   "sign up problem",
				Status:    500,
				Timestamp: strfmt.DateTime(time.Now().UTC()),
			})
		}

		return authentication.NewPostV1BomondVnAuthSignInOK().WithPayload(&models.SuccessResponse{
			Code:      "200",
			Data:      id,
			Message:   "Success",
			Status:    200,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}
}
