package database

import (
	"fmt"
	"postgressPractice/internal/validate"
	"testing"
)

func TestAddCutData(t *testing.T){
	db:= DBConnect()
	d := &Datastore{
		DB: db,
	}
	data:= `{ "customer": "Mary Clark", "items": {"product": "Toy Train","qty": 2}}`
	if validate.IsJSON(data) {
		d.AddCustData(data)
		return
	}
	fmt.Println("invalid json input")
}
