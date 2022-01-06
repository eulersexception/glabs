package model

import (
	"fmt"

	DB "modernc.org/ql"
)

type Clone struct {
	CloneID   *int64 `ql:"index xID"`
	LocalPath string `ql:"uindex xLocalPath"`
	Branch    string
}

func NewClone(localPath string, branch string) (*Clone, string) {
	if localPath == "" {
		return nil, "Enter valid local path."
	}

	if branch == "" {
		return nil, "Enter valid branch."
	}

	clone := &Clone{
		LocalPath: localPath,
		Branch:    branch,
	}

	clone.setClone()

	return clone, ""
}

func (c Clone) setClone() {
	db := GetDB()
	defer FlushAndClose(db)

	if _, _, e := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO Clone IF NOT EXISTS (LocalPath, Branch) VALUES ($1, $2);
		COMMIT;
	`, c.LocalPath, c.Branch); e != nil {
		panic(e)
	}
}

func GetClone(localPath string) *Clone {
	db := GetDB()
	defer FlushAndClose(db)

	rss, _, e := db.Run(DB.NewRWCtx(), `
				BEGIN TRANSACTION;
					SELECT * FROM Clone WHERE  LocalPath = $1;
				COMMIT;
			`, localPath)

	if e != nil {
		panic(e)
	}

	c := &Clone{}

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

func GetAllClones() []Clone {
	db := GetDB()
	defer FlushAndClose(db)

	rss, _, e := db.Run(nil, `
	 	SELECT * FROM Clone;
	`)

	if e != nil {
		panic(e)
	}

	clones := make([]Clone, 0)

	for _, rs := range rss {
		c := &Clone{}

		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(c, data); err != nil {
				return false, err
			}

			clones = append(clones, *c)

			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	return clones
}

func GetAllAssignmentsForClone(path string) []Assignment {
	db := GetDB()
	defer FlushAndClose(db)

	rss, _, e := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			SELECT * FROM Assignment WHERE LocalPath = $1;
		COMMIT;
	`, path)

	if e != nil {
		panic(e)
	}

	assignments := make([]Assignment, 0)

	for _, rs := range rss {
		a := &Assignment{}

		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(a, data); err != nil {
				return false, err
			}

			assignments = append(assignments, *a)

			return true, nil
		}); er != nil {
			panic(er)
		}

	}

	return assignments
}

func (c *Clone) UpdateClone() {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE Clone Branch = $1 WHERE LocalPath = $2;
			COMMIT;
	`, c.Branch, c.LocalPath)

	if err != nil {
		panic(err)
	}
}

func UpdateClonePath(oldPath string, newPath string) {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE Clone LocalPath = $1 WHERE LocalPath = $2;
				UPDATE Assignment LocalPath = $1 WHERE LocalPath = $2;
			COMMIT;
		`, newPath, oldPath)

	if err != nil {
		panic(err)
	}
}

func DeleteClone(localPath string) {
	db := GetDB()
	defer FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Clone WHERE LocalPath = $1;
			UPDATE Assignment LocalPath = $2 WHERE LocalPath = $1;
		COMMIT;
	`, localPath, ""); err != nil {
		panic(err)
	}
}

func (c Clone) toString() string {
	return fmt.Sprintf("\tClone:\n\t\tLocalPath:\t%s\n\t\tBranch:\t%s", c.LocalPath, c.Branch)
}
