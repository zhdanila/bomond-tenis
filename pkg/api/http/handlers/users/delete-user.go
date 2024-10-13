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

type DeleteUser struct {
	ctrl controller2.Controller
}

func NewDeleteUser(ctrl controller2.Controller) *DeleteUser {
	return &DeleteUser{ctrl: ctrl}
}

type DeleteUserService interface {
	Handle(params users.DeleteV1BomondVnUsersUserIDParams) middleware.Responder
}

func (h *DeleteUser) Handle(params users.DeleteV1BomondVnUsersUserIDParams, principal interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.DeleteUserQuery{
		UserId: params.UserID,
	}

	err := h.ctrl.Exec(ctx, q)
	if err != nil {
		return users.NewDeleteV1BomondVnUsersUserIDBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error creating delete user handler",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	return users.NewDeleteV1BomondVnUsersUserIDOK().WithPayload(&models2.SuccessResponse{
		Code:      "200",
		Message:   "Success",
		Status:    200,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	})
}
