package users

import (
	models2 "bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/users"
	controller2 "bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/db/query"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"time"
)

type GetUser struct {
	ctrl controller2.Controller
}

func NewGetUser(ctrl controller2.Controller) *GetUser {
	return &GetUser{ctrl: ctrl}
}

type GetUserService interface {
	Handle(params users.GetV1BomondVnUsersUserIDParams) middleware.Responder
}

func (h *GetUser) Handle(params users.GetV1BomondVnUsersUserIDParams, principal interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.GetUserQuery{
		UserId: params.UserID,
	}

	err := h.ctrl.Exec(ctx, q)
	if err != nil {
		return users.NewGetV1BomondVnUsersUserIDBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error executing get user",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	return users.NewGetV1BomondVnUsersUserIDOK().WithPayload(&models2.SuccessResponse{
		Code:      "200",
		Message:   "Success",
		Data:      q.Out,
		Status:    200,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	})
}
