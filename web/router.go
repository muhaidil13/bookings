package main

import (
	"net/http"

	"github.com/Bookings/internal/controller"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(Nosurf)
	r.Use(sessionload)

	r.Get("/", controller.Repo.Home)
	r.Get("/room/clasic", controller.Repo.Clasik)
	r.Get("/room/generals", controller.Repo.Generals)
	r.Get("/room/vip", controller.Repo.Vip)
	r.Get("/check-availability", controller.Repo.CheckAvailability)
	r.Post("/post-availability", controller.Repo.PostAvailability)
	r.Route("/admin", func(mux chi.Router) {
		mux.Get("/dashboard", controller.Repo.AdminDashboard)
	})

	fileserver := http.FileServer(http.Dir("../static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileserver))
	return r
}
