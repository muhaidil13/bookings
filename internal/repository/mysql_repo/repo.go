package mysql_repo

import (
	"database/sql"

	"github.com/Bookings/internal/config"
	repostitory "github.com/Bookings/internal/repository"
)

type MysqlDbrepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewMysqlRepo(a *config.AppConfig, con *sql.DB) repostitory.Repo {
	return &MysqlDbrepo{
		App: a,
		DB:  con,
	}
}
