package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	username string
	password string
	host     string
	port     int
	dbName   string
}

func New(username, password, host, dbName string, port int) *Database {
	db := &Database{
		username: username,
		password: password,
		host:     host,
		port:     port,
		dbName:   dbName,
	}

	// wait until docker postgres ready
	for {
		dbCon, err := db.Connect()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}

		err = dbCon.Ping()
		if err != nil {
			time.Sleep(time.Second)
			log.Println(err)
			continue
		}
		break
	}

	return db
}

func (d Database) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", d.username, d.password, d.host, d.port, d.dbName))
	if err != nil {
		return nil, err
	}
	return db, err
}
