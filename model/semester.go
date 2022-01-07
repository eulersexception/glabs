package model

import (
	"fmt"

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
	db := GetDB()
	defer FlushAndClose(db)

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
	db := GetDB()
	defer FlushAndClose(db)

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

func (s *Semester) UpdateSemester(course string) {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			UPDATE Semester CoursePath = $1 WHERE SemesterPath = $2;
		COMMIT;
	`, course, s.Path)

	if err != nil {
		panic(err)
	}
}

func (s *Semester) UpdateSemesterPath(semPath string) {
	assignments := GetAllAssignmentsForSemester(s.Path)

	for _, v := range assignments {
		v.SemesterPath = semPath
		v.UpdateAssignment()
	}
}

func DeleteSemester(path string) {
	db := GetDB()
	defer db.Close()

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Semester WHERE SemesterPath = $1;
		COMMIT;
	`, path); err != nil {
		panic(err)
	}
}

func GetAllSemestersForCourse(coursePath string) []Semester {
	db := GetDB()

	rss, _, e := db.Run(DB.NewRWCtx(), `
			SELECT * FROM Semester WHERE CoursePath = $1;
		`, coursePath)

	if e != nil {
		panic(e)
	}

	semesters := make([]Semester, 0)

	for _, rs := range rss {
		s := &Semester{}

		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(s, data); err != nil {
				return false, nil
			}

			semesters = append(semesters, *s)

			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	FlushAndClose(db)

	return semesters
}
