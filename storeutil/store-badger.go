package storeutil

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

type badgerWrapper struct {
	db *badger.DB
}

func (w badgerWrapper) dbCheck() {
	if w.db == nil {
		log.Fatalln(fmt.Errorf("badgerWrapper: db not set"))
	}
}

func (w *badgerWrapper) open() error {
	db, err := badger.Open(badger.DefaultOptions("./data/badger.db"))
	if err != nil {
		return err
	}
	w.db = db

	return nil
}

func (w badgerWrapper) get(key []byte) ([]byte, error) {
	w.dbCheck()
	var data []byte
	err := w.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		item.Value(func(val []byte) error {
			data = val
			return nil
		})
		return nil
	})
	return data, err
}

func (w badgerWrapper) set(key, data []byte) error {
	w.dbCheck()
	err := w.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, data)
		return err
	})

	return err
}

func (w badgerWrapper) has(key []byte) (bool, error) {
	val, err := w.get(key)
	if err != nil || val == nil {
		return false, err
	}

	return true, nil
}
