package database

import (
	"fmt"
	"log"
)

func (d *Datastore)Read(){
	query := `SELECT * FROM orders;`
	orders := []Orders{}
	err :=d.DB.Select(&orders,query)
	if err!=nil{
		log.Fatal(err)
	}
	for i,_:=range orders {
		fmt.Println("orders", orders[i])
	}
}

func (d *Datastore)ReadAll(){
	query := `SELECT data FROM orders;`
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

func (d *Datastore)GetAllCustomers(){
	query := `SELECT data -> 'customer' AS customer FROM orders;`  //-> is returns JSON object field by key

	customer,err := d.DB.Queryx(query)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("get all customers in form of JSON:")
	for customer.Next(){
		var data string
		_=customer.Scan(&data)
		fmt.Println("	",data)
	}

	//to get all customers in form of text
	query = `SELECT data ->> 'customer' AS customer FROM orders;`   //->> returns JSON object field by tex
	customer,err = d.DB.Queryx(query)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("get all customers in form of text:")
	for customer.Next(){
		var data string
		_=customer.Scan(&data)
		fmt.Println("	",data)
	}
}

//get list of all sold products
func (d *Datastore)ProductSold(){
	query := `SELECT data -> 'items' ->> 'product' as product FROM orders  ORDER BY product;`
	products,err := d.DB.Queryx(query)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("all sold products:")
	for products.Next(){
		var product string
		_=products.Scan(&product)
		fmt.Println("	",product)
	}
}





