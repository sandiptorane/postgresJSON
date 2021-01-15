package database

import (
	"fmt"
	"log"
)

func (d *Datastore)AddCustData(dataToAdd string){
	query := `INSERT INTO orders (data) VALUES($1)`;
	result,err:= d.DB.Exec(query,dataToAdd)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("result after inserting:",result)
}

