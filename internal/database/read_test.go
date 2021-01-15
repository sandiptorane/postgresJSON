package database

import "testing"

func TestDatastore_Read(t *testing.T) {
	d := &Datastore{
		DB: DBConnect(),
	}
	//d.Read()
	d.ReadAll()
}


func TestDatastore_GetAllCustomers(t *testing.T) {
	d := &Datastore{
		DB: DBConnect(),
	}
	d.GetAllCustomers()
}

func TestDatastore_ProductSold(t *testing.T) {
	d := &Datastore{
		DB: DBConnect(),
	}
	d.ProductSold()
}
