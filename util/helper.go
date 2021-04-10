package util

import (
	"log"
	"os"
	"regexp"
	"time"

	"github.com/dgraph-io/badger/v3"
)

// EmailRegex is a regular expression to match emails.
var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// DB name passed as command line argument
var DBNAME = os.Args[1]

// IsValidMail checks if a provided string is a valid mail address.
// Returns true or false.
func IsValidMail(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return EmailRegex.MatchString(e)
}

// GetDB opens a connection to the database with the given name. If there is no DB, it will be created.
// Argument is a string that creates a path inside folder where function is called, i.e. "tmp/badger", "tmp/test".
// The connection must be closed after function is called. It is recommended to call 'defer db.Close()' immediately.
// Returns a pointer to a badger.DB object.
func GetDB() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions(DBNAME))
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func cleanUpDB(db *badger.DB) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
	again:
		err := db.RunValueLogGC(0.7)
		if err == nil {
			goto again
		}
	}
	time.Sleep(8 * time.Second)
}
