package application

import (
	"fmt"
	"huddlet/config"
	"huddlet/db"
	"huddlet/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	Name   string
	Env    string
	Port   int
	Router *chi.Mux
	DB     *db.DB
	Config *config.Config
}

func Run() error {
	application, err := initApp()
	if err != nil {
		return err
	}
	defer application.Close()

	log.Println("Running application on port:", application.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", application.Port), application.Router)
	return err
}

func initApp() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())
	if err != nil {
		return nil, err
	}

	return &Application{
		Name:   "Huddle",
		Env:    cfg.GetEnviourment(),
		Port:   cfg.GetPort(),
		Config: cfg,
		Router: routes.Get(cfg),
		DB:     db,
	}, nil
}

func (app *Application) Close() {
	app.DB.Close()
}
