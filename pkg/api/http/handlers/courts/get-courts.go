package courts

import (
	models2 "bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/courts"
	controller2 "bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/db/query"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"time"
)

type GetCourts struct {
	ctrl controller2.Controller
}

func NewGetCourts(ctrl controller2.Controller) *GetCourts {
	return &GetCourts{ctrl: ctrl}
}

type GetUserService interface {
	Handle(params courts.GetV1BomondVnCourtsParams) middleware.Responder
}

func (h *GetCourts) Handle(params courts.GetV1BomondVnCourtsParams, principal interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.GetCourtsQuery{}

	err := h.ctrl.Exec(ctx, q)
	if err != nil {
		return courts.NewGetV1BomondVnCourtsBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error creating get courts handler",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	return courts.NewGetV1BomondVnCourtsOK().WithPayload(&models2.SuccessResponse{
		Code:      "200",
		Message:   "Success",
		Status:    200,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	})
}
