package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type AppConfig struct {
	DB *sqlx.DB
}

func (appConfig *AppConfig) Configure() {

	appConfig.configureDatabase()
}

func (appConfig *AppConfig) configureDatabase() {

	var db *sqlx.DB

	db, err := sqlx.Connect("mysql", "root:mysecretpassword@(localhost:3306)/ticktap")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()

	appConfig.DB = db

}
