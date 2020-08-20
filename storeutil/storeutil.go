package storeutil

import (
	"log"
)

var db store

type store interface {
	get([]byte) ([]byte, error)
	set([]byte, []byte) error
	has([]byte) (bool, error)
}

func init() {
	// bw := simpleMem{}
	bw := badgerWrapper{}
	err := bw.open()
	if err != nil {
		log.Fatalln(err)
	}

	db = bw
}

func Get(key string) []byte {
	data, _ := db.get([]byte(key))
	return data
}

func Set(key string, data []byte) {
	db.set([]byte(key), data)
}

func Has(key string) bool {
	val, _ := db.has([]byte(key))
	return val
}
