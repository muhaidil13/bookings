package driver

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SqlDB struct {
	DB *sql.DB
}

const (
	maxidle     = 5
	maxcon      = 10
	maxlifetime = 2 * time.Minute
)

var con = &SqlDB{}

func ConnectSql(dsn string) (*SqlDB, error) {
	db, err := openCon(dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(maxidle)
	db.SetConnMaxLifetime(maxlifetime)
	db.SetMaxOpenConns(maxcon)
	con.DB = db

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("open")
	return con, nil
}

func openCon(dsn string) (*sql.DB, error) {
	count := 1
	for {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Println("Trying Connect database in ", count)
			count++
		} else {
			return db, nil
		}
		if count > 10 {
			return nil, err
		}
		time.Sleep(2 * time.Second)
		count++
		continue
	}

}
