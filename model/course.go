package model

import (
	"fmt"

	util "github.com/eulersexception/glabs-ui/util"
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
	db := util.GetDB()
	defer util.FlushAndClose(db)

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
	db := util.GetDB()
	defer util.FlushAndClose(db)

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

func DeleteSemesterFromCourse(path string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Course WHERE CoursePath = $1;
		COMMIT;
	`, path)

	if err != nil {
		panic(err)
	}
}
