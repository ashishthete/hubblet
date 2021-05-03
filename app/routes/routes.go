package routes

import (
	"huddlet/pkg/api/users"
	"huddlet/pkg/ui"
	"net/http"

	"huddlet/config"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Get(conf *config.Config) *chi.Mux {
	router := chi.NewRouter()
	// router.Use(middleware.RequestID)
	// router.Use(middleware.Logger)
	// router.Use(middleware.Recoverer)
	//authenticated api

	router.HandleFunc("/login", ui.Login)
	router.HandleFunc("/signup", ui.SignUp)
	router.HandleFunc("/dashboard", ui.DashBoardPageHandler)
	router.HandleFunc("/logout", ui.LogoutHandler)
	router.HandleFunc("/like", ui.LikePost)
	router.HandleFunc("/dislike", ui.DislikePost)
	router.HandleFunc("/posts/{id}/comments", ui.AddComment)

	router.Mount("/api", apiRouter())
	//public api
	return router
}

// A completely separate router for administrator routes
func apiRouter() http.Handler {
	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		userController := users.Controller{}
		r.Group(func(r chi.Router) {
			// Stop processing after 2.5 seconds.
			r.Use(middleware.Timeout(2500 * time.Millisecond))

			r.Get("/users", userController.List)
		})
	})
	return router
}
