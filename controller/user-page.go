package controller

import (
	"net/http"

	"github.com/MRP/config"
	"github.com/MRP/model"
	"github.com/MRP/render"
)

var Repo *Repostitory

type Repostitory struct {
	App *config.AppConfig
}

func InitRepo(ap *config.AppConfig) *Repostitory {
	return &Repostitory{
		App: ap,
	}
}

func InitRepository(repo *Repostitory) {
	Repo = repo
}

func (m *Repostitory) Home(w http.ResponseWriter, r *http.Request) {
	render.SetTemplate(w, r, "home.page.html", &model.TemplateModel{})
}

func (m *Repostitory) Generals(w http.ResponseWriter, r *http.Request) {
	render.SetTemplate(w, r, "generals.page.html", &model.TemplateModel{})
}

func (m *Repostitory) Vip(w http.ResponseWriter, r *http.Request) {
	render.SetTemplate(w, r, "vip.page.html", &model.TemplateModel{})
}

func (m *Repostitory) Clasik(w http.ResponseWriter, r *http.Request) {
	render.SetTemplate(w, r, "clasik.page.html", &model.TemplateModel{})
}
func (m *Repostitory) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	render.SetTemplate(w, r, "check-availability.page.html", &model.TemplateModel{})
}
