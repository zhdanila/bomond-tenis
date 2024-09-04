package courts

import (
	"bomond-tenis/internal/api/models"
	"bomond-tenis/internal/api/restapi/operations/courts"
	"bomond-tenis/internal/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func GetCourtByIdHandler(service service.Courts) courts.GetV1BomondVnCourtIDBookHandlerFunc {
	return func(params courts.GetV1BomondVnCourtIDBookParams) middleware.Responder {
		fmt.Println("get by id court")

		return courts.NewGetV1BomondVnCourtIDBookOK().WithPayload(&models.SuccessResponse{
			Code:      "200",
			Data:      "nothing",
			Err:       "nothing",
			Message:   "hello",
			Status:    200,
			Timestamp: strfmt.DateTime{},
		})
	}
}
