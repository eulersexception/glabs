package model

import (
	"fmt"

	DB "modernc.org/ql"
)

type Course struct {
	CourseID *int64 `ql:"index xID"`
	Path     string `ql:"uindex xPath, name CoursePath"`
}

func NewCourse(path string) *Course {
	if path == "" {
		fmt.Println("Enter valid course path.")
		return nil
	}

	c := &Course{
		Path: path,
	}

	c.setCourse()

	return c
}

func (c *Course) setCourse() {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO Course IF NOT EXISTS (CoursePath) VALUES ($1);
		COMMIT;
		`, c.Path)

	if err != nil {
		panic(err)
	}
}

func GetCourse(path string) *Course {
	db := GetDB()
	defer FlushAndClose(db)

	rss, _, e := db.Run(nil, `
		SELECT * FROM Course WHERE CoursePath = $1;
	`, path)

	if e != nil {
		panic(e)
	}

	c := &Course{}

	for _, rs := range rss {
		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(c, data); err != nil {
				return false, err
			}

			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	return c
}

func UpdateCourse(oldPath string, newPath string) {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			UPDATE Course CoursePath = $1 WHERE CoursePath = $2;
			UPDATE Semester CoursePath = $1 WHERE CoursePath = $2;
		COMMIT;
	`, newPath, oldPath)

	if err != nil {
		panic(err)
	}
}

func DeleteCourse(path string) {
	semesters := GetAllSemestersForCourse(path)

	if len(semesters) == 0 {
		db := GetDB()
		defer FlushAndClose(db)

		_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Course WHERE CoursePath = $1;
		COMMIT;
	`, path)

		if err != nil {
			panic(err)
		}
	}
}

func (c *Course) AddSemesterToCourse(path string) {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			UPDATE Semester CoursePath = $1 WHERE SemesterPath = $2;
		COMMIT;
	`, c.Path, path)

	if err != nil {
		panic(err)
	}
}

func GetAllCourses() []Course {
	db := GetDB()
	rss, _, e := db.Run(nil, `SELECT * FROM Course;`)

	if e != nil {
		panic(e)
	}

	courses := make([]Course, 0)

	for _, rs := range rss {
		c := &Course{}

		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(c, data); err != nil {
				return false, err
			}

			courses = append(courses, *c)

			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	FlushAndClose(db)

	return courses
}
