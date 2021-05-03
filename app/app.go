package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"huddlet/app/db"
	"huddlet/app/routes"
	"huddlet/config"
)

type Application struct {
	Name   string
	Env    string
	Port   int
	Router *chi.Mux
	DB     *db.DB
	Config *config.Config
}

var appInstance *Application

func Run() error {
	appInstance, err := initApp()
	if err != nil {
		return err
	}
	defer appInstance.Close()

	log.Println("Running application on port:", appInstance.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", appInstance.Port), appInstance.Router)
	return err
}

func initApp() (*Application, error) {
	cfg := config.Get()
	db, err := db.Connect(cfg.GetDBConnStr())
	if err != nil {
		log.Println("Error connecting to db")
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

func GetSecret() string {
	return appInstance.Config.GetSecret()
}

func (app *Application) Close() {
	app.DB.Close()
}
