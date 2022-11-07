package main

import (
	"log"
	"net/http"
	"time"

	"github.com/MRP/config"
	"github.com/MRP/controller"
	"github.com/MRP/render"
	"github.com/alexedwards/scs/v2"
)

var port = ":8080"

var App config.AppConfig
var session *scs.SessionManager

func main() {

	App.AppSecure = false
	App.UseCache = false
	App.AppSecureCookie = false

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

	repo := controller.InitRepo(&App)
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
