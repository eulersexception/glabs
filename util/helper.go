package util

import (
	"fmt"
	"log"
	"os"
	"regexp"

	database "modernc.org/ql"
)

var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

var DBNAME = os.Args[1]

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

var OPTIONS = &database.Options{
	CanCreate:      true,
	Headroom:       20000,
	RemoveEmptyWAL: true,
}

func InitLoggers() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
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
