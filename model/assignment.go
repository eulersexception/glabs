package model

import (
	"fmt"

	util "github.com/eulersexception/glabs-ui/util"
	DB "modernc.org/ql"
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
		return nil, "Enter valid assignment path."
	}

	if sem == "" {
		return nil, "Enter valid semester path."
	}

	if per == "" {
		return nil, "Enter valid per."
	}

	if desc == "" {
		return nil, "Enter valid description."
	}

	if localPath == "" {
		return nil, "Enter valid local path."
	}

	if branch == "" {
		return nil, "Enter valid branch."
	}

	if starterUrl == "" {
		return nil, "Enter valid starter url."
	}

	if fromBranch == "" {
		return nil, "Enter valid from branch."
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

	rss, _, e := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				SELECT * FROM Assignment
				WHERE AssignmentPath = $1;
			COMMIT;
		`, path)

	if e != nil {
		panic(e)
	}

	a := &Assignment{}

	for _, rs := range rss {

		if er := rs.Do(false, func(data []interface{}) (bool, error) {

			if err := DB.Unmarshal(a, data); err != nil {
				return false, err
			}

			return true, nil

		}); er != nil {
			panic(er)
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
			DELETE FROM TeamAssignment WHERE AssignmentPath = $1;
		COMMIT;
	`, path); err != nil {
		panic(err)
	}
}

func (a *Assignment) UpdateAssignment() {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE Assignment
					SemesterPath = $2, Per = $3, Description = $4, ContainerRegistry = $5, LocalPath = $6, StarterUrl = $7  
				WHERE AssignmentPath = $1;
			COMMIT;
	`, a.AssignmentPath, a.SemesterPath, a.Per, a.Description, a.ContainerRegistry, a.LocalPath, a.StarterUrl); err != nil {
		panic(err)
	}
}

func GetAllAssignmentsForSemester(semesterPath string) []Assignment {
	db := util.GetDB()

	rss, _, e := db.Run(DB.NewRWCtx(), `
		SELECT * FROM Assignment
		WHERE SemesterPath = $1;
	`, semesterPath)

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

	util.FlushAndClose(db)

	return assignments
}

func (as *Assignment) AddTeam(name string) {
	NewTeamAssignment(name, as.AssignmentPath)
}

func (as *Assignment) RemoveTeam(name string) {
	RemoveTeamFromAssignment(name, as.AssignmentPath)
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

func GetStarterCode(url string) *StarterCode {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	rss, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
					SELECT * FROM StarterCode
					WHERE  Url = $1;
				COMMIT;
			`, url)

	if err != nil {
		panic(err)
	}

	s := &StarterCode{}

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

func DeleteStarterCode(url string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM StarterCode WHERE Url = $1;
		COMMIT;
	`, url); err != nil {
		panic(err)
	}
}

func (s *StarterCode) UpdateStarterCode() {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE StarterCode
					FromBranch = $2, ProtectToBranch = $3
					WHERE Url = $1;
			COMMIT;
	`, s.Url, s.FromBranch, s.ProtectToBranch)

	if err != nil {
		panic(err)
	}
}

func (s StarterCode) toString() string {
	return fmt.Sprintf("\tStarterCode:\n\t\tUrl:\t%s\n\t\tFromBranch:\t%s\n\t\tProtectToBranch:\t%v", s.Url, s.FromBranch, s.ProtectToBranch)
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

func GetClone(localPath string) *Clone {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	rss, _, err := db.Run(DB.NewRWCtx(), `
				BEGIN TRANSACTION;
					SELECT * FROM Clone
					WHERE  LocalPath = $1;
				COMMIT;
			`, localPath)

	if err != nil {
		panic(err)
	}

	c := &Clone{}

	for _, rs := range rss {

		if err := rs.Do(false, func(data []interface{}) (bool, error) {

			if e := DB.Unmarshal(c, data); e != nil {
				return false, e
			}

			return true, nil
		}); err != nil {
			panic(err)
		}
	}

	return c
}

func DeleteClone(localPath string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Clone WHERE LocalPath = $1;
		COMMIT;
	`, localPath); err != nil {
		panic(err)
	}
}

func (c *Clone) UpdateClone() {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE Clone
					Branch = $1 
					WHERE LocalPath = $2;
			COMMIT;
	`, c.Branch, c.LocalPath)

	if err != nil {
		panic(err)
	}
}

func (c Clone) toString() string {
	return fmt.Sprintf("\tClone:\n\t\tLocalPath:\t%s\n\t\tBranch:\t%s", c.LocalPath, c.Branch)
}
