package database

import "testing"

func TestDatastore_Aggregations(t *testing.T) {
	d := &Datastore{
		DB: DBConnect(),
	}
	d.Aggregations()
}
