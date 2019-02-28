package main

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func getDB(host string, port int, user string, password string, dbname string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Error(err)
		return nil
	}
	log.Info("successfull connected!")
	return db
}
