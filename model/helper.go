package model

import (
	"fmt"
	"os"
	"regexp"

	database "modernc.org/ql"
)

var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

var DBNAME = os.Args[1]

var OPTIONS = &database.Options{
	CanCreate:      true,
	Headroom:       20000,
	RemoveEmptyWAL: true,
}

func IsValidMail(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}

	return EmailRegex.MatchString(e)
}

func GetDB() *database.DB {
	var db *database.DB
	var err error

	if DBNAME == "" {
		db, err = database.OpenFile("default_db", OPTIONS)
	} else {
		db, err = database.OpenFile(DBNAME, OPTIONS)
	}

	if err != nil {

		fmt.Printf("%v\n", err.Error())
		panic(err)
	}

	return db
}

func FlushAndClose(db *database.DB) {
	db.Flush()
	db.Close()
}
