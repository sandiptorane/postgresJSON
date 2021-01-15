package database

import (
	"fmt"
	"log"
)

func (d *Datastore)CreateTable(){
	query := `CREATE TABLE IF NOT EXISTS orders (
	id SERIAL NOT NULL PRIMARY KEY,
	data json NOT NULL
	);`
	_,err:= d.DB.Exec(query)
	if err!=nil{
		fmt.Println("error while creating table:",err)
		return
	}

	log.Println("Table created successfully")
}

func (d *Datastore)Drop(){
	_,err:=d.DB.Exec(`DROP TABLE IF EXISTS orders`)
	if err!=nil{
		log.Fatal(err)
	}
}

