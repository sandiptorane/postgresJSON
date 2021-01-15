package database

import (
	"fmt"
	"log"
)


//print JSON object into a set of key-value pairs
func (d *Datastore)JsonEach(){
	query := `SELECT json_each (data) FROM orders;`
	orders,err := d.DB.Queryx(query)
	if err!=nil{
		log.Fatal(err)
	}
	for orders.Next(){
		var data string
		_=orders.Scan(&data)
		fmt.Println("orders",data)
	}
}

func (d *Datastore)JsonOjectKeys(){
	//display keys of items field
     query :=`SELECT json_object_keys (data->'items')FROM orders;`
	var keys []string
	err := d.DB.Select(&keys,query)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("keys:",keys)
}

//get type of key
func (d *Datastore)JsonType(){
	//display key type of items field
	query :=`SELECT json_typeof (data->'items')FROM orders;`
	var types []string
	err := d.DB.Select(&types,query)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("types of items:",types)

	//gets types of qty field
	query =`SELECT json_typeof (data->'items'->'qty')FROM orders;`
	var qtyTypes []string
	err = d.DB.Select(&qtyTypes,query)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("types of quatity:",qtyTypes)
}
