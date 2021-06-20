package model

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	DB "modernc.org/ql"

	util "github.com/eulersexception/glabs-ui/util"
)

// Student - Id (Matrikelnummer) is the primary key. All fields are public and
// Getter or Setter functions relate to database operations.
type Student struct {
	StudentID  int64 `ql:"index xID"`
	MatrikelNr int64 `ql:"uindex xMatrikelNr"`
	Name       string
	FirstName  string
	NickName   string
	Email      string
}

// NewStudent creates a new student and stores the object in DB.
// Arguments id of type uint32 and strings for name and firstName must not be empty and a well formed email must be provided.
// If a student with given id exists already in DB, the existing dataset will be overwritten.
// Returns a pointer to a new student and a message string. If provided arguments are invalid the message will not be empty.
func NewStudent(name string, firstName string, nickName string, email string, matrikelNr int64) (*Student, string) {
	if name == "" || firstName == "" {
		res := "\n+++ Please provide valid name or first name.\n"
		return nil, res
	}

	if Mail(email) == false {
		res := "\n+++ Please provide valid email address.\n"
		return nil, res
	}

	existing := GetStudent(matrikelNr)
	empty := &Student{}

	if !cmp.Equal(existing, empty) {
		return existing, "Student already exists - use update for changes"
	}

	stud := &Student{
		MatrikelNr: matrikelNr,
		NickName:   nickName,
		Email:      email,
		Name:       name,
		FirstName:  firstName,
	}

	stud.setStudent()

	return stud, ""
}

// Mail checks if the given string is a well formed email address.
// Returns true or false.
func Mail(email string) bool {
	if util.IsValidMail(email) {
		return true
	} else {
		return false
	}
}

// GetMail of student.
// Returns email string.
func (s *Student) GetMail() string {
	return s.Email
}

// This function updates student record in DB. An update could be a creation or edition of a record.
func (s *Student) setStudent() {
	db := util.GetDB()

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO Student IF NOT EXISTS (MatrikelNr, Name, FirstName, NickName, Email) VALUES ($1, $2, $3, $4, $5);
		COMMIT;
		`, s.MatrikelNr, s.Name, s.FirstName, s.NickName, s.Email)

	if err != nil {
		panic(err)
	}

	util.FlushAndClose(db)
}

// GetStudent fetches student from DB with an argument of type int64 as Matrikelnr.
// Returns an error if fetch fails or a pointer to the student.
func GetStudent(matrikelNr int64) *Student {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	rss, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				SELECT id(), MatrikelNr, Name, FirstName, NickName, Email FROM Student
				WHERE MatrikelNr = $1;
			COMMIT;
		`, matrikelNr)

	if err != nil {
		panic(err)
	}

	s := &Student{}

	for _, rs := range rss {

		if err := rs.Do(false, func(data []interface{}) (bool, error) {

			if e := DB.Unmarshal(s, data); e != nil {
				return false, e
			}

			return true, nil

		}); err != nil {
			panic(err)
		}
	}

	return s
}

// DeleteStudent removes a student by id (uint64) from DB.
// Returns an error if operation fails.
func DeleteStudent(matrikelNr int64) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Student WHERE MatrikelNr = $1;
			DELETE FROM StudentTeam WHERE MatrikelNr = $1;
		COMMIT;
	`, matrikelNr); err != nil {
		panic(err)
	}
}

// UpdateStudent changes a students record in DB.
// Returns an error if the update fails.
func (s *Student) UpdateStudent() bool {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE Student
					Name = $1, FirstName = $2, NickName = $3, Email = $4
				WHERE MatrikelNr = $5;
			COMMIT;
	`, s.Name, s.FirstName, s.NickName, s.Email, s.MatrikelNr); err != nil {
		panic(err)
	}

	return true
}

// PrintData outputs a human readable string for students data.
func (s *Student) PrintData() {
	fmt.Printf("\n-------------------\nName:\t\t%s %s\nNick:\t\t%s\nMailTo:\t\t%s\nId:\t\t%d\n",
		s.FirstName, s.Name, s.NickName, s.Email, s.MatrikelNr)
}

func (fst *Student) Equals(scd *Student) bool {
	if scd == nil {
		return false
	}
	return fst.MatrikelNr == scd.MatrikelNr && fst.Email == scd.Email && fst.Name == scd.Name && fst.FirstName == scd.FirstName
}

// JoinTeam adds student to team.
// Expects the team name (string).
// Returns an error if the team doesn't exist otherwise nil (succesful operation).
func (s Student) JoinTeam(team string) {
	NewStudentTeam(s.MatrikelNr, team)
}
