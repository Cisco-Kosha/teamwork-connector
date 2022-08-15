package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kosha/teamwork-connector/pkg/config"
	"github.com/kosha/teamwork-connector/pkg/logger"
)

type App struct {
	Router *mux.Router
	Log    logger.Logger
	Cfg    *config.Config
}

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}

// Initialize creates the necessary scaffolding of the app
func (a *App) Initialize(log logger.Logger) {

	cfg := config.Get()

	a.Cfg = cfg
	a.Log = log

	a.Router = router()

	a.initializeRoutes()

}

// Run starts the app and serves on the specified addr
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
