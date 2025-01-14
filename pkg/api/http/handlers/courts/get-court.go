package courts

import (
	models2 "bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/courts"
	controller2 "bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/db/query"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"log"
	"time"
)

type GetBookedCourt struct {
	ctrl controller2.Controller
}

func NewGetBookedCourt(ctrl controller2.Controller) *GetBookedCourt {
	return &GetBookedCourt{ctrl: ctrl}
}

type GetBookedCourtService interface {
	Handle(params courts.GetV1BomondVnCourtIDBookParams) middleware.Responder
}

func (h *GetBookedCourt) Handle(params courts.GetV1BomondVnCourtIDBookParams, principal interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.GetBookedCourtQuery{
		CourtId: params.CourtID,
	}

	err := h.ctrl.Exec(ctx, q)
	if err != nil {
		return courts.NewGetV1BomondVnCourtIDBookBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error creating get booked court handler",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	log.Printf("q.Out: %+v\n", q.Out)

	return courts.NewGetV1BomondVnCourtIDBookOK().WithPayload(&models2.SuccessResponse{
		Code:      "200",
		Message:   "Success",
		Data:      q.Out,
		Status:    200,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	})
}
