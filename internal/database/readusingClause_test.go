package database

import "testing"

func TestDatastore_WhoBoughtProduct(t *testing.T) {
	d := &Datastore{
		DB: DBConnect(),
	}
	d.WhoBoughtProduct()
}
