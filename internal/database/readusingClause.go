package database

import (
	"fmt"
	"log"
)

type customers struct{
	Name string `db:"customer"`
	Product string  `db:"product"`
}

//to find out who bought particular product
func (d *Datastore)WhoBoughtProduct() {
	query := `SELECT data ->> 'customer' AS customer FROM orders
			WHERE data -> 'items' ->> 'product' = 'Toy Train';`

	customer, err := d.DB.Queryx(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("cust who baught toy train")
	for customer.Next() {
		var cust string
		_ = customer.Scan(&cust)
		fmt.Println("	", cust)
	}


	query = `SELECT data ->> 'customer' AS customer,
		data -> 'items' ->> 'product' AS product
	FROM orders
	WHERE CAST ( data -> 'items' ->> 'qty' AS INTEGER) = 2`
	customer, err = d.DB.Queryx(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("cust who baught 2 product at a time")
	for customer.Next(){
		cust := customers{}
		err:= customer.StructScan(&cust)
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Println(cust)
	}

}
