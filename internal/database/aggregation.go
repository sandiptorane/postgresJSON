package database

import (
	"fmt"
	"log"
)

func (d *Datastore)Aggregations(){
	query := `SELECT MAX (CAST (data -> 'items' ->> 'qty' AS INTEGER)) FROM orders;`
	var max int
	err :=d.DB.Get(&max,query)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("MAX of Products:",max)

	//minimum product quantity
	query = `SELECT MIN (CAST (data -> 'items' ->> 'qty' AS INTEGER)) FROM orders;`
	var min int
	err =d.DB.Get(&min,query)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("MIN of Products:",min)

	//average product quantity
	query = `SELECT AVG (CAST (data -> 'items' ->> 'qty' AS INTEGER)) FROM orders;`
	var avg float32
	err =d.DB.Get(&avg,query)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("average of Products quantity:",avg)

}
