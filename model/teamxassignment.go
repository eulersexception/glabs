package model

import (
	"fmt"

	DB "modernc.org/ql"
)

type TeamAssignment struct {
	NamePath       string `ql:"uindex xNamePath"`
	TeamName       string
	AssignmentPath string
}

func NewTeamAssignment(name string, path string) {
	db := GetDB()
	defer FlushAndClose(db)

	id := fmt.Sprintf("%s%s", name, path)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO TeamAssignment IF NOT EXISTS (NamePath, TeamName, AssignmentPath) 
			VALUES ($1, $2, $3);
		COMMIT;
	`, id, name, path)

	if err != nil {
		panic(err)
	}
}

func GetAssignmentsForTeam(name string) []Assignment {
	db := GetDB()

	rss, _, e := db.Run(DB.NewRWCtx(), `
			SELECT NamePath, TeamName, AssignmentPath  
			FROM TeamAssignment 
			WHERE TeamName = $1;
		`, name)

	if e != nil {
		panic(e)
	}

	entries := make([]TeamAssignment, 0)

	for _, rs := range rss {
		t := &TeamAssignment{}

		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(t, data); err != nil {
				return false, err
			}

			entries = append(entries, *t)
			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	FlushAndClose(db)
	assignments := make([]Assignment, 0)

	for _, v := range entries {
		a := GetAssignment(v.AssignmentPath)
		assignments = append(assignments, *a)
	}

	return assignments
}

func GetTeamsForAssignment(path string) []Team {
	db := GetDB()

	rss, _, e := db.Run(DB.NewRWCtx(), `
			SELECT * FROM TeamAssignment 
			WHERE AssignmentPath = $1;
		`, path)

	if e != nil {
		panic(e)
	}

	entries := make([]TeamAssignment, 0)

	for _, rs := range rss {
		t := &TeamAssignment{}

		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(t, data); err != nil {
				return false, err
			}

			entries = append(entries, *t)
			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	FlushAndClose(db)
	teams := make([]Team, 0)

	for _, v := range entries {
		team := GetTeam(v.TeamName)
		teams = append(teams, *team)
	}

	return teams
}

func RemoveTeamFromAssignment(name string, path string) {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM TeamAssignment 
			WHERE TeamName = $1 AND AssignmentPath = $2;
		COMMIT;
	`, name, path)

	if err != nil {
		panic(err)
	}
}

func RemoveAssignmentsForTeam(name string) {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM TeamAssignment WHERE TeamName = $1;
		COMMIT;
	`, name)

	if err != nil {
		panic(err)
	}
}

func UpdateTeamForAssignments(oldName string, newName string) {
	assignments := GetAssignmentsForTeam(oldName)
	RemoveAssignmentsForTeam(oldName)

	for _, v := range assignments {
		NewTeamAssignment(newName, v.AssignmentPath)
	}
}

func RemoveTeamsForAssignment(path string) {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM TeamAssignment WHERE AssignmentPath = $1;
		COMMIT;
	`, path)

	if err != nil {
		panic(err)
	}
}

func UpdateAssignmentForTeams(oldPath string, newPath string) {
	teams := GetTeamsForAssignment(oldPath)
	RemoveTeamsForAssignment(oldPath)

	for _, v := range teams {
		NewTeamAssignment(v.Name, newPath)
	}
}
