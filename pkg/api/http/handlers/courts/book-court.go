package courts

import (
	models2 "bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/courts"
	controller2 "bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/db/query"
	"bomond-tenis/pkg/utils"
	"errors"
	"fmt"
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

func (h *BookCourt) Handle(params courts.PostV1BomondVnCourtIDBookParams, principal interface{}) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	q := &query.GetBookedCourtQuery{
		CourtId: params.CourtID,
	}

	err := h.ctrl.Exec(ctx, q)
	if err != nil && errors.Is(err, fmt.Errorf("sql: no rows in result set")) {
		return courts.NewPostV1BomondVnCourtIDBookBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "ErrorRA",
			Status:    400,
			Timestamp: strfmt.DateTime(time.Now().UTC()),
		})
	}

	for _, court := range q.Out.BookedCourts {
		check, err := utils.CheckTimeBookingBusy(court.Court, query.BookedCourt{
			Date:     params.BookingRequest.Date,
			Duration: params.BookingRequest.Duration,
		})
		if err != nil {
			return courts.NewPostV1BomondVnCourtIDBookBadRequest().WithPayload(&models2.ErrorResult{
				Code:      "400",
				DebugInfo: err.Error(),
				Message:   "Error",
				Status:    400,
				Timestamp: strfmt.DateTime(time.Now().UTC()),
			})
		}

		if check {
			return courts.NewPostV1BomondVnCourtIDBookBadRequest().WithPayload(&models2.ErrorResult{
				Code:      "400",
				Message:   "Error, time is busy",
				Status:    400,
				Timestamp: strfmt.DateTime(time.Now().UTC()),
			})
		}
	}

	q2 := &query.BookCourtQuery{
		BookCourt: query.BookedCourt{
			CourtId:  params.CourtID,
			UserID:   params.BookingRequest.UserID,
			Date:     params.BookingRequest.Date,
			Duration: params.BookingRequest.Duration,
		},
	}

	err = h.ctrl.Exec(ctx, q2)
	if err != nil {
		return courts.NewPostV1BomondVnCourtIDBookBadRequest().WithPayload(&models2.ErrorResult{
			Code:      "400",
			DebugInfo: err.Error(),
			Message:   "Error exec booking",
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
