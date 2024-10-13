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

type UpdateUser struct {
	ctrl controller2.Controller
}

func NewUpdateUser(ctrl controller2.Controller) *UpdateUser {
	return &UpdateUser{ctrl: ctrl}
}

type UpdateUserService interface {
	Handle(params users.PutV1BomondVnUsersUserIDParams) middleware.Responder
}

func (h *UpdateUser) Handle(params users.PutV1BomondVnUsersUserIDParams, principal interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.UpdateUserQuery{
		UserId:   params.UserID,
		Username: params.UserUpdate.Username,
		Email:    params.UserUpdate.Email,
		Password: params.UserUpdate.Password,
	}

	err := h.ctrl.Exec(ctx, q)
	if err != nil {
		return users.NewPutV1BomondVnUsersUserIDBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error creating update user handler",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	return users.NewPutV1BomondVnUsersUserIDOK().WithPayload(&models2.SuccessResponse{
		Code:      "200",
		Message:   "Success",
		Status:    200,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	})
}
