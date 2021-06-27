package model

import (
	"fmt"

	"github.com/eulersexception/glabs-ui/util"

	DB "modernc.org/ql"
)

type Semester struct {
	SemesterID *int64 `ql:"index xID"`
	Path       string `ql:"uindex xPath, name SemesterPath"`
	CoursePath string
}

func NewSemester(coursePath string, path string) *Semester {
	if path == "" {
		fmt.Println("Enter valid path for semester.")
		return nil
	}

	if coursePath == "" {
		fmt.Println("Enter valid course path.")
		return nil
	}

	s := &Semester{
		Path:       path,
		CoursePath: coursePath,
	}

	s.setSemester()

	return s
}

func (s *Semester) setSemester() {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	_, _, e := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO Semester IF NOT EXISTS (SemesterPath, CoursePath) VALUES ($1, $2);
		COMMIT;
	`, s.Path, s.CoursePath)

	if e != nil {
		panic(e)
	}
}

func GetSemester(path string) *Semester {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	rss, _, e := db.Run(nil, `
		SELECT * FROM Semester WHERE SemesterPath = $1;
	`, path)

	if e != nil {
		panic(e)
	}

	s := &Semester{}

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

func DeleteSemester(path string) {
	db := util.GetDB()
	defer db.Close()

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Semester WHERE SemesterPath = $1;
		COMMIT;
	`, path); err != nil {
		panic(err)
	}
}

func (s *Semester) UpdateSemester(course string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			UPDATE Semester CoursePath = $1 WHERE SemesterPath = $2;
		COMMIT;
	`, course, s.Path)

	if err != nil {
		panic(err)
	}
}
