package model

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	util "github.com/eulersexception/glabs-ui/util"
	DB "modernc.org/ql"
)

type Team struct {
	TeamID *int64 `ql:"index xID"`
	Name   string `ql:"uindex xName, name TeamName"`
}

func NewTeam(name string) (*Team, string) {
	if name == "" {
		res := "\n+++ Enter valid team name."
		return nil, res
	}

	existing := GetTeam(name)
	empty := &Team{}

	if !cmp.Equal(existing, empty) {
		return existing, "Team already exists - use update for changes"
	}

	team := &Team{
		Name: name,
	}

	team.setTeam()

	return team, ""
}

func (t *Team) setTeam() {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO Team IF NOT EXISTS (TeamName) VALUES ($1);
		COMMIT;
		`, t.Name)

	if DB.IsDuplicateUniqueIndexError(err) {
		fmt.Printf("Duplicate Index ------- %v\n", err)
	} else if err != nil {
		panic(err)
	}
}

func GetTeam(name string) *Team {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	rss, _, e := db.Run(DB.NewRWCtx(), `
				BEGIN TRANSACTION;
					SELECT * FROM Team WHERE TeamName = $1;
				COMMIT;
			`, name)

	if e != nil {
		panic(e)
	}

	t := &Team{}

	for _, rs := range rss {

		if er := rs.Do(false, func(data []interface{}) (bool, error) {

			if err := DB.Unmarshal(t, data); err != nil {
				return false, err
			}

			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	return t
}

func (t *Team) UpdateTeam(newName string) {
	//check := GetTeam(newName)

	//if check.Name == newName {
	//	util.WarningLogger.Printf("Team with name %s already exists.\n", newName)
	//} else {

	UpdateTeamNameForStudents(t.Name, newName)
	UpdateTeamForAssignments(t.Name, newName)

	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
				BEGIN TRANSACTION;
					UPDATE Team	TeamName = $1 WHERE TeamName = $2;
				COMMIT;
		`, newName, t.Name); err != nil {
		panic(err)
	}
	//}
}

func DeleteTeam(name string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Team WHERE TeamName = $1;
			DELETE FROM StudentTeam WHERE TeamName = $1;
			DELETE FROM TeamAssignment WHERE TeamName = $1;
		COMMIT;
	`, name); err != nil {
		panic(err)
	}
}

func (t Team) JoinAssignment(assignmentPath string) {
	NewTeamAssignment(t.Name, assignmentPath)
}

func (t *Team) AddStudent(s *Student) {
	NewStudentTeam(s.MatrikelNr, t.Name)
}

func (t *Team) RemoveStudent(s *Student) {
	RemoveStudentFromTeam(s.MatrikelNr, t.Name)
}

func (fst *Team) Equals(scd *Team) bool {
	if scd == nil || fst.Name != scd.Name {
		return false
	}

	return true
}
