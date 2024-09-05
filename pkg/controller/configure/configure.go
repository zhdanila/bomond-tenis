package configure

import (
	config "bomond-tenis/internal/config"
	"bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/db/auth_db"
	"bomond-tenis/pkg/db/courts"
	"bomond-tenis/pkg/db/users_db"
	"github.com/jmoiron/sqlx"
)

func ControllerInit(ctrlRegistry *controller.ControllerImpl, ctrl controller.Controller, postgres *sqlx.DB, env config.Environment) (e error) {

	propogateErr := func(err error) {
		if err != nil {
			e = err
		}
	}
	// db
	// auth
	propogateErr(ctrlRegistry.RegisterHandler(auth_db.NewLogoutHandler(postgres)))
	propogateErr(ctrlRegistry.RegisterHandler(auth_db.NewSignInHandler(postgres)))
	propogateErr(ctrlRegistry.RegisterHandler(auth_db.NewSignUpHandler(postgres)))

	//users
	propogateErr(ctrlRegistry.RegisterHandler(users_db.NewGetUserHandler(postgres)))
	propogateErr(ctrlRegistry.RegisterHandler(users_db.NewUpdateUserHandler(postgres)))
	propogateErr(ctrlRegistry.RegisterHandler(users_db.NewDeleteUserHandler(postgres)))

	//courts
	propogateErr(ctrlRegistry.RegisterHandler(courts.NewGetCourtsHandler(postgres)))
	propogateErr(ctrlRegistry.RegisterHandler(courts.NewCancelCourtBookingQueryHandler(postgres)))
	propogateErr(ctrlRegistry.RegisterHandler(courts.NewGetBookedCourtHandler(postgres)))
	propogateErr(ctrlRegistry.RegisterHandler(courts.NewBookCourtHandler(postgres)))

	return
}
