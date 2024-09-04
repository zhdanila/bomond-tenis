package courts

import (
	"bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/courts"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func PostCourtHandler() courts.PostV1BomondVnCourtIDBookHandlerFunc {
	return func(params courts.PostV1BomondVnCourtIDBookParams) middleware.Responder {
		fmt.Println("post court")

		return courts.NewPostV1BomondVnCourtIDBookOK().WithPayload(&models.SuccessResponse{
			Code:      "200",
			Data:      "nothing",
			Err:       "nothing",
			Message:   "hello",
			Status:    200,
			Timestamp: strfmt.DateTime{},
		})
	}
}
