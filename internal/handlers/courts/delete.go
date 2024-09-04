package courts

import (
	"bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/courts"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func DeleteCourtHandler() courts.DeleteV1BomondVnCourtIDBookBookIDHandlerFunc {
	return func(params courts.DeleteV1BomondVnCourtIDBookBookIDParams) middleware.Responder {
		fmt.Println("delete court")

		return courts.NewDeleteV1BomondVnCourtIDBookBookIDOK().WithPayload(&models.SuccessResponse{
			Code:      "200",
			Data:      "nothing",
			Err:       "nothing",
			Message:   "hello",
			Status:    200,
			Timestamp: strfmt.DateTime{},
		})
	}
}
