package auth

import (
	models2 "bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/authentication"
	controller2 "bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type SignIn struct {
	ctrl controller2.Controller
}

func NewSignIn(ctrl controller2.Controller) *SignIn {
	return &SignIn{ctrl: ctrl}
}

type SignInService interface {
	Handle(params authentication.PostV1BomondVnAuthSignInParams) middleware.Responder
}

func (h *SignIn) Handle(params authentication.PostV1BomondVnAuthSignInParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.SignInQuery{
		Email: params.UserSignin.Email,
	}

	err := h.ctrl.Exec(ctx, q)
	if err != nil {
		return authentication.NewPostV1BomondVnAuthSignInBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error creating sign in handler",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(q.Out.Account.Password), []byte(params.UserSignin.Password))
	if err != nil {
		return authentication.NewPostV1BomondVnAuthSignInBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Incorrect password",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	token, err := utils.GenerateJWT(q.Out.Account.Email, q.Out.Account.ID)
	if err != nil {
		return authentication.NewPostV1BomondVnAuthSignInBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error generating jwt token",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	err = h.ctrl.Exec(ctx, q)
	if err != nil {
		return authentication.NewPostV1BomondVnAuthSignUpBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error creating sign in handler",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	return authentication.NewPostV1BomondVnAuthSignUpOK().WithPayload(&models2.SuccessResponse{
		Code:      "200",
		Message:   token,
		Status:    200,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	})
}
