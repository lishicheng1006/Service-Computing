package db

import "sync"

type SimpleDB struct {
	*sync.Map
}

var instance *SimpleDB

func init() {
	instance = new(SimpleDB)
	instance.Map = new(sync.Map)
}

func Get() *SimpleDB {
	return instance
}
