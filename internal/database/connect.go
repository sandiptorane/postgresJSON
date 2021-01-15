package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "perennial"
	password = "perennial"
	dbname   = "supermart"
)


type Orders struct {
	Id   int  `json:"id,omitempty"`
	Data string
}

type Datastore struct {
	DB *sqlx.DB
}



func DBConnect() *sqlx.DB {
	postgresInfo :=fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Db, err := sqlx.Connect("postgres", postgresInfo)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected successfully")
	return Db
}

