package database

import "testing"



func TestDatastore_JsonEach(t *testing.T) {
	d := &Datastore{
		DB: DBConnect(),
	}
	d.JsonEach()
}

func TestDatastore_JsonOjectKeys(t *testing.T) {
	d := &Datastore{
		DB: DBConnect(),
	}
	d.JsonOjectKeys()
}

func TestDatastore_JsonType(t *testing.T) {
	d := &Datastore{
		DB: DBConnect(),
	}
	d.JsonType()
}


