package model

import (
	"fmt"

	DB "modernc.org/ql"

	util "github.com/eulersexception/glabs-ui/util"
)

type Assignment struct {
	AssignmentID      *int64 `ql:"index xID"`
	AssignmentPath    string `ql:"uindex xAssignmentPath"`
	SemesterPath      string
	Per               string
	Description       string
	ContainerRegistry bool
	LocalPath         string
	StarterUrl        string
}

type StarterCode struct {
	StarterCodeID   *int64 `ql:"index xID"`
	Url             string `ql:"uindex xUrl"`
	FromBranch      string
	ProtectToBranch bool
}

type Clone struct {
	CloneID   *int64 `ql:"index xID"`
	LocalPath string `ql:"uindex xLocalPath"`
	Branch    string
}

func NewAssignment(assignmentPath string, sem string, per string,
	desc string, conRegistry bool, localPath string,
	branch string, starterUrl string, fromBranch string,
	protectToBranch bool) (*Assignment, string) {

	if assignmentPath == "" {
		return nil, "Please enter valid assignment path."
	}

	if sem == "" {
		return nil, "Please enter valid semester path."
	}

	if per == "" {
		return nil, "Please enter valid per."
	}

	if desc == "" {
		return nil, "Please enter valid description."
	}

	if localPath == "" {
		return nil, "Please enter valid local path."
	}

	if branch == "" {
		return nil, "Please enter valid branch."
	}

	if starterUrl == "" {
		return nil, "Please enter valid starter url."
	}

	if fromBranch == "" {
		return nil, "Please enter valid from branch."
	}

	assignment := &Assignment{
		AssignmentPath:    assignmentPath,
		SemesterPath:      sem,
		Per:               per,
		Description:       desc,
		ContainerRegistry: conRegistry,
		LocalPath:         localPath,
		StarterUrl:        starterUrl,
	}

	assignment.setAssignment()

	starterCoder := &StarterCode{
		Url:             starterUrl,
		FromBranch:      fromBranch,
		ProtectToBranch: protectToBranch,
	}

	starterCoder.setStarterCode()

	clone := &Clone{
		LocalPath: localPath,
		Branch:    branch,
	}

	clone.setClone()

	return assignment, ""
}

func (a Assignment) setAssignment() {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO Assignment IF NOT EXISTS (AssignmentPath, SemesterPath, Per, Description, ContainerRegistry, LocalPath, StarterUrl) VALUES ($1, $2, $3, $4, $5, $6, $7);
		COMMIT;
		`, a.AssignmentPath, a.SemesterPath, a.Per, a.Description, a.ContainerRegistry, a.LocalPath, a.StarterUrl)

	if err != nil {
		panic(err)
	}
}

func GetAssignment(path string) *Assignment {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	rss, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				SELECT * FROM Assignment
				WHERE AssignmentPath = $1;
			COMMIT;
		`, path)

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

func DeleteAssignment(path string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Assignment WHERE AssignmentPath = $1;
		COMMIT;
	`, path); err != nil {
		panic(err)
	}
}

func (a *Assignment) UpdateAssignment() bool {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE Assignment
					AssignmentPath = $1, SemesterPath = $2, Per = $3, Description = $4, ContainerRegistry = $5, LocalPath = $6, StarterUrl = $7  
				WHERE AssignmentPath = $1;
			COMMIT;
	`, a.AssignmentPath, a.SemesterPath, a.Per, a.Description, a.ContainerRegistry, a.LocalPath, a.StarterUrl); err != nil {
		panic(err)
	}

	return true
}

func (s StarterCode) setStarterCode() {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, e := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO StarterCode IF NOT EXISTS (Url, FromBranch, ProtectToBranch) VALUES ($1, $2, $3);
		COMMIT;
	`, s.Url, s.FromBranch, s.ProtectToBranch); e != nil {
		panic(e)
	}
}

func (c Clone) setClone() {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, e := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO Clone IF NOT EXISTS (LocalPath, Branch) VALUES ($1, $2);
		COMMIT;
	`, c.LocalPath, c.Branch); e != nil {
		panic(e)
	}
}

func (s StarterCode) toString() string {
	return fmt.Sprintf("\tStarterCode:\n\t\tUrl:\t%s\n\t\tFromBranch:\t%s\n\t\tProtectToBranch:\t%v", s.Url, s.FromBranch, s.ProtectToBranch)
}

func (c Clone) toString() string {
	return fmt.Sprintf("\tClone:\n\t\tLocalPath:\t%s\n\t\tBranch:\t%s", c.LocalPath, c.Branch)
}
