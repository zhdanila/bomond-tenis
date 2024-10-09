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

type BookCourt struct {
	ctrl controller2.Controller
}

func NewBookCourt(ctrl controller2.Controller) *BookCourt {
	return &BookCourt{ctrl: ctrl}
}

type BookCourtService interface {
	Handle(params courts.PostV1BomondVnCourtIDBookParams) middleware.Responder
}

func (h *BookCourt) Handle(params courts.PostV1BomondVnCourtIDBookParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.BookCourtQuery{
		CourtId:  params.CourtID,
		Date:     params.BookingRequest.Date,
		Duration: params.BookingRequest.Duration,
		Time:     params.BookingRequest.Time,
		UserID:   params.BookingRequest.UserID,
	}

	err := h.ctrl.Exec(ctx, q)
	if err != nil {
		return courts.NewPostV1BomondVnCourtIDBookBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error creating book court handler",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	return courts.NewPostV1BomondVnCourtIDBookOK().WithPayload(&models2.SuccessResponse{
		Code:      "200",
		Message:   "Success",
		Status:    200,
		Timestamp: strfmt.DateTime(time.Now().UTC()),
	})
}
