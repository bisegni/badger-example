package main

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v2"
)

// A ...
type A struct {
	db *badger.DB
}

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Your code here…

	// Start a writable transaction.
	txn := db.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	err = txn.Set([]byte("answer"), []byte("42"))
	if err != nil {

	}

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
	}

	err = db.View(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte("answer"))

		var valNot, valCopy []byte
		_ = item.Value(func(val []byte) error {
			// This func with val would only be called if item.Value encounters no error.

			// Accessing val here is valid.
			fmt.Printf("The answer is: %s\n", val)

			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)

			// Assigning val slice to another variable is NOT OK.
			valNot = val // Do not do this.
			return nil
		})
		return nil
	})
}
