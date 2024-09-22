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

type CancelCourtBooking struct {
	ctrl controller2.Controller
}

func NewCancelCourtBooking(ctrl controller2.Controller) *CancelCourtBooking {
	return &CancelCourtBooking{ctrl: ctrl}
}

type CancelCourtBookingService interface {
	Handle(params courts.DeleteV1BomondVnCourtIDBookBookIDParams) middleware.Responder
}

func (h *CancelCourtBooking) Handle(params courts.DeleteV1BomondVnCourtIDBookBookIDParams, principal interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.CancelCourtBookingQuery{
		CourtId: params.CourtID,
		BookId:  params.BookID,
	}

	err := h.ctrl.Exec(ctx, q)
	if err != nil {
		return courts.NewDeleteV1BomondVnCourtIDBookBookIDBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error creating cancel booked court handler",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	return courts.NewDeleteV1BomondVnCourtIDBookBookIDOK().WithPayload(&models2.SuccessResponse{
		Code:      "200",
		Message:   "Success",
		Status:    200,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	})
}
