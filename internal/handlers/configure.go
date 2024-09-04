package handlers

import (
	"bomond-tenis/internal/api/restapi/operations"
	"bomond-tenis/internal/handlers/auth"
	"bomond-tenis/internal/handlers/courts"
	"bomond-tenis/internal/handlers/users"
)

// func ConfigureHandlers(api *operations.BomondTenisAPI, userService *service.UserService) {
func ConfigureHandlers(api *operations.BomondTenisAPI) {
	//Users
	api.UsersGetV1BomondVnUsersUserIDHandler = users.GetUserHandler()
	api.UsersPutV1BomondVnUsersUserIDHandler = users.PutUserHandler()
	api.UsersDeleteV1BomondVnUsersUserIDHandler = users.DeleteUserHandler()

	//Authentication
	api.AuthenticationPostV1BomondVnAuthSignInHandler = auth.SignInHandler()
	api.AuthenticationPostV1BomondVnAuthSignUpHandler = auth.SignUpHandler()
	api.AuthenticationPostV1BomondVnAuthLogoutHandler = auth.LogoutHandler()

	//Courts
	api.CourtsGetV1BomondVnCourtsHandler = courts.GetAllCourtsHandler()
	api.CourtsGetV1BomondVnCourtIDBookHandler = courts.GetCourtByIdHandler()
	api.CourtsPostV1BomondVnCourtIDBookHandler = courts.PostCourtHandler()
	api.CourtsDeleteV1BomondVnCourtIDBookBookIDHandler = courts.DeleteCourtHandler()
}
