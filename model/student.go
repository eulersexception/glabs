package model

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	DB "modernc.org/ql"

	util "github.com/eulersexception/glabs-ui/util"
)

type Student struct {
	StudentID  *int64 `ql:"index xID"`
	MatrikelNr int64  `ql:"uindex xMatrikelNr"`
	Name       string
	FirstName  string
	NickName   string
	Email      string
}

func NewStudent(name string, firstName string, nickName string, email string, matrikelNr int64) (*Student, string) {
	if name == "" || firstName == "" {
		res := "\n+++ Enter valid name or first name.\n"
		return nil, res
	}

	if !util.IsValidMail(email) {
		res := "\n+++ Enter valid email address.\n"
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

func (s *Student) setStudent() {
	db := util.GetDB()

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO Student IF NOT EXISTS (MatrikelNr, Name, FirstName, NickName, Email) 
			VALUES ($1, $2, $3, $4, $5);
		COMMIT;
		`, s.MatrikelNr, s.Name, s.FirstName, s.NickName, s.Email)

	if err != nil {
		panic(err)
	}

	util.FlushAndClose(db)
}

func GetStudent(matrikelNr int64) *Student {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	rss, _, e := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				SELECT * FROM Student WHERE MatrikelNr = $1;
			COMMIT;
		`, matrikelNr)

	if e != nil {
		panic(e)
	}

	s := &Student{}

	for _, rs := range rss {

		if er := rs.Do(false, func(data []interface{}) (bool, error) {

			if err := DB.Unmarshal(s, data); err != nil {
				return false, err
			}

			return true, nil

		}); er != nil {
			panic(er)
		}
	}

	return s
}

func (s *Student) UpdateStudent() {
	db := util.GetDB()
	// defer util.FlushAndClose(db)

	// util.WarningLogger.Printf("Student before update: Matrikel %d, name = %s, nickname = %s\n", s.MatrikelNr, s.Name, s.NickName)

	if _, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE Student
					Name = $1, FirstName = $2, NickName = $3, Email = $4
				WHERE MatrikelNr = $5;
			COMMIT;
	`, s.Name, s.FirstName, s.NickName, s.Email, s.MatrikelNr); err != nil {
		panic(err)
	}

	util.FlushAndClose(db)

	// newS := GetStudent(s.MatrikelNr)

	// util.WarningLogger.Printf("Student after update: Matrikel %d, name = %s, nickname = %s\n", newS.MatrikelNr, newS.Name, newS.NickName)

}

func UpdateMatrikelNummer(oldNum int64, newNum int64) {
	if oldNum != newNum {
		UpdateStudentMatrikel(oldNum, newNum)
		db := util.GetDB()
		defer util.FlushAndClose(db)

		_, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE Student MatrikelNr = $1
				WHERE MatrikelNr = $2;
			COMMIT;
		`, newNum, oldNum)

		if err != nil {
			panic(err)
		}
	}
}

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

func (s Student) JoinTeam(team string) {
	NewStudentTeam(s.MatrikelNr, team)
}

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
