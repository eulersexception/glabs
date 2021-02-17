package main

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"
)

func main() {

	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	k := "Bahadir"
	v := "Suezer"
	updates := make(map[string]string)
	updates[k] = v

	err = db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(k), []byte(v))
		err = txn.SetEntry(e)
		return err
	})

	err = db.View(func(txn *badger.Txn) error {
		item, getErr := txn.Get([]byte(k))

		if getErr != nil {
			log.Fatal(getErr)
		}

		err = item.Value(func(val []byte) error {
			fmt.Println(fmt.Sprintf("Key = %s, Value = %s", item.String(), string(val)))
			return err
		})
		return getErr
	})

	//glabsView.NewHomeview().Window.ShowAndRun()

}
