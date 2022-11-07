package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Bookings/internal/config"
	"github.com/Bookings/internal/controller"
	"github.com/Bookings/internal/driver"
	"github.com/Bookings/internal/model"
	"github.com/Bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
)

var port = ":8080"
var env = "../.env"
var App config.AppConfig
var session *scs.SessionManager

func main() {
	gob.Register(model.ReservationRoom{})
	App.AppSecure = false
	App.UseCache = false
	App.AppSecureCookie = false

	// Setup Template for template cache
	tc, err := render.GetTemplate()
	if err != nil {
		log.Println(err)
	}

	App.TemplateCache = tc

	session = scs.New()
	session.Cookie.Persist = true
	session.Lifetime = 3 * time.Hour
	session.Cookie.Secure = App.AppSecure
	session.Cookie.SameSite = http.SameSiteLaxMode

	App.Session = session

	// Setup database
	err = godotenv.Load(env)
	if err != nil {
		log.Println("Failed to load File .env")
	}
	mysqdsn := os.Getenv("MYSQL_DSN")
	db, err := driver.ConnectSql(mysqdsn)
	if err != nil {
		log.Println(err)
	}
	repo := controller.InitRepo(&App, db)
	controller.InitRepository(repo)

	render.Init(&App)

	srv := &http.Server{
		Addr:    port,
		Handler: Router(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
