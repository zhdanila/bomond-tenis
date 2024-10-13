package configure

import (
	config "bomond-tenis/internal/config"
	"bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/db/auth_db"
	"bomond-tenis/pkg/db/courts"
	"bomond-tenis/pkg/db/users_db"
	"github.com/jmoiron/sqlx"
	redis2 "github.com/redis/go-redis/v9"
)

func ControllerInit(ctrlRegistry *controller.ControllerImpl, ctrl controller.Controller, pool *sqlx.DB, redis *redis2.Client, env config.Environment) (e error) {

	propogateErr := func(err error) {
		if err != nil {
			e = err
		}
	}
	// db
	// auth
	propogateErr(ctrlRegistry.RegisterHandler(auth_db.NewLogoutHandler(pool, redis)))
	propogateErr(ctrlRegistry.RegisterHandler(auth_db.NewSignInHandler(pool)))
	propogateErr(ctrlRegistry.RegisterHandler(auth_db.NewSignUpHandler(pool)))

	//users
	propogateErr(ctrlRegistry.RegisterHandler(users_db.NewGetUserHandler(pool)))
	propogateErr(ctrlRegistry.RegisterHandler(users_db.NewUpdateUserHandler(pool)))
	propogateErr(ctrlRegistry.RegisterHandler(users_db.NewDeleteUserHandler(pool)))

	//courts
	propogateErr(ctrlRegistry.RegisterHandler(courts.NewGetCourtsHandler(pool)))
	propogateErr(ctrlRegistry.RegisterHandler(courts.NewCancelCourtBookingQueryHandler(pool)))
	propogateErr(ctrlRegistry.RegisterHandler(courts.NewGetBookedCourtHandler(pool)))
	propogateErr(ctrlRegistry.RegisterHandler(courts.NewBookCourtHandler(pool)))

	return
}
