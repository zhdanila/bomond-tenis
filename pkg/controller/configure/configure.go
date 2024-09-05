package configure

import (
	config "bomond-tenis/internal/config"
	"bomond-tenis/pkg/controller"
	"bomond-tenis/pkg/db/auth_db"
	"github.com/jmoiron/sqlx"
)

func ControllerInit(ctrlRegistry *controller.ControllerImpl, ctrl controller.Controller, pool *sqlx.DB, env config.Environment) (e error) {

	propogateErr := func(err error) {
		if err != nil {
			e = err
		}
	}
	// db
	// events
	propogateErr(ctrlRegistry.RegisterHandler(auth_db.NewLogoutHandler(pool)))

	return
}
