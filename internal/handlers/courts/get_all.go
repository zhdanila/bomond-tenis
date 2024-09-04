package courts

import (
	"bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/courts"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func GetAllCourtsHandler() courts.GetV1BomondVnCourtsHandlerFunc {
	return func(params courts.GetV1BomondVnCourtsParams) middleware.Responder {
		fmt.Println("get all court")

		return courts.NewGetV1BomondVnCourtsOK().WithPayload(&models.SuccessResponse{
			Code:      "200",
			Data:      "nothing",
			Err:       "nothing",
			Message:   "hello",
			Status:    200,
			Timestamp: strfmt.DateTime{},
		})
	}
}
