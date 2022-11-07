package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/Bookings/internal/config"
	"github.com/Bookings/internal/driver"
	"github.com/Bookings/internal/model"
	"github.com/Bookings/internal/render"
	repostitory "github.com/Bookings/internal/repository"
	"github.com/Bookings/internal/repository/mysql_repo"
)

var Repo *Repostitory

type Repostitory struct {
	App    *config.AppConfig
	DBrepo repostitory.Repo
}

func InitRepo(ap *config.AppConfig, db *driver.SqlDB) *Repostitory {
	return &Repostitory{
		App:    ap,
		DBrepo: mysql_repo.NewMysqlRepo(ap, db.DB),
	}
}

func InitRepository(repo *Repostitory) {
	Repo = repo
}

func (m *Repostitory) Home(w http.ResponseWriter, r *http.Request) {
	err := render.SetTemplate(w, r, "home.page.html", &model.TemplateModel{})
	if err != nil {
		log.Println(err)
	}
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

func (m *Repostitory) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start-date")
	end := r.Form.Get("end-date")

	layout := "01/02/2006"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		log.Println("Error Parse Start Date", err)
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		log.Println("Error Parse end Date", err)
	}

	// Get Availability Rooms
	rooms, err := m.DBrepo.SearchAvailabilityRooms(startDate, endDate)
	if err != nil {
		log.Println(err)
	}

	if len(rooms) == 0 {
		m.App.Session.Put(r.Context(), "error", "Room Tidak Tersedia Sekarang!")
		http.Redirect(w, r, "/check-availability", http.StatusSeeOther)
		return
	}
	res := model.ReservationRoom{
		StartDate: startDate,
		EndDate:   endDate,
	}
	m.App.Session.Put(r.Context(), "reservation", res)
	data := make(map[string]any)
	data["rooms"] = rooms
	render.SetTemplate(w, r, "show-room.page.html", &model.TemplateModel{
		Data: data,
	})

}
