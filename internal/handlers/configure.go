package handlers

import (
	"bomond-tenis/internal/api/restapi/operations"
	"bomond-tenis/internal/handlers/auth"
	"bomond-tenis/internal/handlers/courts"
	"bomond-tenis/internal/handlers/users"
	"bomond-tenis/internal/service"
)

func ConfigureHandlers(api *operations.BomondTenisAPI, service *service.Service) {
	//Users
	api.UsersGetV1BomondVnUsersUserIDHandler = users.GetUserHandler(service.Users)
	api.UsersPutV1BomondVnUsersUserIDHandler = users.PutUserHandler(service.Users)
	api.UsersDeleteV1BomondVnUsersUserIDHandler = users.DeleteUserHandler(service.Users)

	//Authentication
	api.AuthenticationPostV1BomondVnAuthSignInHandler = auth.SignInHandler(service.Authorization)
	api.AuthenticationPostV1BomondVnAuthSignUpHandler = auth.SignUpHandler(service.Authorization)
	api.AuthenticationPostV1BomondVnAuthLogoutHandler = auth.LogoutHandler(service.Authorization)

	//Courts
	api.CourtsGetV1BomondVnCourtsHandler = courts.GetAllCourtsHandler(service.Courts)
	api.CourtsGetV1BomondVnCourtIDBookHandler = courts.GetCourtByIdHandler(service.Courts)
	api.CourtsPostV1BomondVnCourtIDBookHandler = courts.PostCourtHandler(service.Courts)
	api.CourtsDeleteV1BomondVnCourtIDBookBookIDHandler = courts.DeleteCourtHandler(service.Courts)
}
