package database

import (
	"testing"
)

func TestDatastore_CreateTable(t *testing.T) {
	Db:= DBConnect()
	datastore := &Datastore{
		DB: Db,
	}
	datastore.CreateTable()
}

func TestDatastore_Drop(t *testing.T) {
	Db:= DBConnect()
	datastore := &Datastore{
		DB: Db,
	}
	datastore.Drop()
}