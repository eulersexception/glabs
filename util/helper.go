package util

import (
	"os"
	"regexp"

	database "modernc.org/ql"
)

// EmailRegex is a regular expression to match emails.
var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// DB name passed as command line argument
var DBNAME = os.Args[1]

var OPTIONS = &database.Options{
	CanCreate:      true,
	Headroom:       20000,
	RemoveEmptyWAL: true,
}

// IsValidMail checks if a provided string is a valid mail address.
// Returns true or false.
func IsValidMail(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return EmailRegex.MatchString(e)
}

// GetDB opens a connection to the database with the given name. If there is no DB, it will be created.
// Argument is a string that creates a file for the embedded database. If the filename contains a path to a folder, the folder
// must be created before calling the function, i.e. "tmp/test".
// The database must be flushed and the connection must be closed after function is called.
// It is recommended to call 'defer db.Close()' immediately.
// Returns a pointer to a ql DB object.
func GetDB() *database.DB {
	var db *database.DB
	var err error

	if DBNAME == "" {
		db, err = database.OpenFile("default_db", OPTIONS)
	} else {
		db, err = database.OpenFile(DBNAME, OPTIONS)
	}

	if err != nil {
		panic(err)
	}

	return db
}

func FlushAndClose(db *database.DB) {
	db.Flush()
	db.Close()
}
