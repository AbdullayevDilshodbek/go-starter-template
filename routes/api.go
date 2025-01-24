package routes

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Routes(r *mux.Router, db *sqlx.DB) {
	UserRoutes(r, db)
}