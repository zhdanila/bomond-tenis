package auth

import (
	models2 "bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/authentication"
	controller2 "bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"time"
)

type SignUp struct {
	ctrl controller2.Controller
}

func NewSignUp(ctrl controller2.Controller) *SignUp {
	return &SignUp{ctrl: ctrl}
}

type SignUpService interface {
	Handle(params authentication.PostV1BomondVnAuthSignUpParams) middleware.Responder
}

func (h *SignUp) Handle(params authentication.PostV1BomondVnAuthSignUpParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.SignUpQuery{
		Email:    params.UserSignup.Email,
		Password: params.UserSignup.Password,
		Username: params.UserSignup.Username,
	}

	hashedPassword, err := utils.GeneratePasswordHash(q.Password)
	if err != nil {
		return authentication.NewPostV1BomondVnAuthSignUpBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error hashing password",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}
	q.Password = hashedPassword

	err = h.ctrl.Exec(ctx, q)
	if err != nil {
		return authentication.NewPostV1BomondVnAuthSignUpBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error creating sign up handler",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	return authentication.NewPostV1BomondVnAuthSignUpOK().WithPayload(&models2.SuccessResponse{
		Code:      "200",
		Message:   "Success",
		Status:    200,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	})
}
