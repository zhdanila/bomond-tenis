package auth

import (
	models2 "bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/authentication"
	controller2 "bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/db/query"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type Logout struct {
	ctrl controller2.Controller
}

func NewLogout(ctrl controller2.Controller) *Logout {
	return &Logout{ctrl: ctrl}
}

type LogoutService interface {
	Handle(params authentication.PostV1BomondVnAuthLogoutParams) middleware.Responder
}

func (h *Logout) Handle(params authentication.PostV1BomondVnAuthLogoutParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.LogoutQuery{}

	err := h.ctrl.Exec(ctx, q)
	if err != nil {
		return authentication.NewPostV1BomondVnAuthLogoutBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error creating logout handler",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	return authentication.NewPostV1BomondVnAuthLogoutOK().WithPayload(&models2.SuccessResponse{
		Code:      "200",
		Message:   "Success",
		Status:    200,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	})
}
