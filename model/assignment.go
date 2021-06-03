package model

import (
	"fmt"

	DB "modernc.org/ql"

	util "github.com/eulersexception/glabs-ui/util"
)

type StarterCode struct {
	Url             string
	FromBranch      string
	ProtectToBranch bool
}

func (s StarterCode) toString() string {
	return fmt.Sprintf("\tStarterCode:\n\t\tUrl:\t%s\n\t\tFromBranch:\t%s\n\t\tProtectToBranch:\t%v", s.Url, s.FromBranch, s.ProtectToBranch)
}

type Clone struct {
	LocalPath string
	Branch    string
}

func (c Clone) toString() string {
	return fmt.Sprintf("\tClone:\n\t\tLocalPath:\t%s\n\t\tBranch:\t%s", c.LocalPath, c.Branch)
}

type Assignment struct {
	AssignmentID      *int64       `ql:"index xID"`
	Name              string       `ql:"uindex xName, name AssignmentName"`
	Semester          string       `ql:"name SemesterName"`
	LocalClone        *Clone       `ql:"-"`
	Starter           *StarterCode `ql:"-"`
	ContainerRegistry bool         `ql:"-"`
}

func NewAssignment(name string, sem string, clone *Clone, starter *StarterCode) *Assignment {
	if name == "" {
		fmt.Println("Please enter a valid assignment name.")
		return nil
	}

	if sem == "" {
		fmt.Println("Please enter a valid semester name.")
		return nil
	}

	db := util.GetDB()

	schema := DB.MustSchema((*Assignment)(nil), "", nil)

	if _, _, e := db.Execute(DB.NewRWCtx(), schema); e != nil {
		panic(e)
	}

	db.Flush()
	db.Close()

	assignment := &Assignment{
		Semester:   sem,
		Name:       name,
		Starter:    starter,
		LocalClone: clone,
	}

	assignment.setAssignment()

	return assignment
}

func (a Assignment) setAssignment() {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO Assignment IF NOT EXISTS (AssignmentName, SemesterName) VALUES ($1, $2);
		COMMIT;
		`, a.Name, a.Semester)

	if DB.IsDuplicateUniqueIndexError(err) {
		fmt.Printf("Duplicate Index ------- %v\n", err)
	} else if err != nil {
		panic(err)
	}
}

func GetAssignment(name string) *Assignment {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	rss, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				SELECT AssignmentID, AssignmentName, SemesterName FROM Assignment
				WHERE AssignmentName = $1;
			COMMIT;
		`, name)

	if err != nil {
		panic(err)
	}

	a := &Assignment{}

	for _, rs := range rss {

		if err := rs.Do(false, func(data []interface{}) (bool, error) {
			if e := DB.Unmarshal(a, data); e != nil {
				return false, e
			}

			return true, nil

		}); err != nil {
			panic(err)
		}
	}

	return a
}

func DeleteAssignment(name string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Assignment WHERE AssignmentName == $1;
		COMMIT;
	`, name); err != nil {
		panic(err)
	}
}

func (a *Assignment) UpdateAssignment() bool {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE Assignment
					AssignmentName = $1, SemesterName = $2 
				WHERE AssignmentName = $1;
			COMMIT;
	`, a.Name, a.Semester); err != nil {
		panic(err)
	}

	return true
}
