package model

import (
	"fmt"

	DB "modernc.org/ql"
)

type StudentTeam struct {
	MatrikelTeam string `ql:"uindex xMatrikelTeam"`
	MatrikelNr   int64
	TeamName     string
}

func NewStudentTeam(matrikelNr int64, teamName string) {
	db := GetDB()
	defer FlushAndClose(db)

	id := fmt.Sprintf("%d%s", matrikelNr, teamName)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO StudentTeam IF NOT EXISTS (MatrikelTeam, MatrikelNr, TeamName) 
			VALUES ($1, $2, $3);
		COMMIT;
	`, id, matrikelNr, teamName)

	if err != nil {
		panic(err)
	}
}

func GetTeamsForStudent(matrikelNr int64) []Team {
	db := GetDB()

	rss, _, e := db.Run(DB.NewRWCtx(), `
			SELECT MatrikelTeam, MatrikelNr, TeamName 
			FROM StudentTeam 
			WHERE MatrikelNr = $1;
		`, matrikelNr)

	if e != nil {
		panic(e)
	}

	entries := make([]StudentTeam, 0)

	for _, rs := range rss {
		s := &StudentTeam{}

		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(s, data); err != nil {
				return false, err
			}

			entries = append(entries, *s)

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

func GetStudentsForTeam(team string) []Student {
	db := GetDB()

	rss, _, e := db.Run(DB.NewRWCtx(), `
			SELECT MatrikelTeam, MatrikelNr, TeamName 
			FROM StudentTeam 
			WHERE TeamName = $1;
		`, team)

	if e != nil {
		panic(e)
	}

	entries := make([]StudentTeam, 0)

	for _, rs := range rss {
		s := &StudentTeam{}

		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(s, data); err != nil {
				return false, err
			}

			entries = append(entries, *s)
			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	FlushAndClose(db)

	studs := make([]Student, 0)

	for _, v := range entries {
		s := GetStudent(v.MatrikelNr)
		studs = append(studs, *s)
	}

	return studs
}

func UpdateTeamNameForStudents(oldTeamName string, newTeamName string) {
	students := GetStudentsForTeam(oldTeamName)
	RemoveStudentsForTeam(oldTeamName)

	for _, v := range students {
		NewStudentTeam(v.MatrikelNr, newTeamName)
	}
}

func UpdateStudentMatrikel(oldNum int64, newNum int64) {
	teams := GetTeamsForStudent(oldNum)
	RemoveTeamsForStudent(oldNum)

	for _, v := range teams {
		NewStudentTeam(newNum, v.Name)
	}
}

func RemoveStudentsForTeam(team string) {
	db := GetDB()
	defer FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM StudentTeam WHERE TeamName = $1;
		COMMIT;
	`, team); err != nil {
		panic(err)
	}
}

func RemoveTeamsForStudent(matrikelNr int64) {
	db := GetDB()
	defer FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM StudentTeam WHERE MatrikelNr = $1;
		COMMIT;
	`, matrikelNr); err != nil {
		panic(err)
	}
}

func RemoveStudentFromTeam(matrikelNr int64, team string) {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM StudentTeam 
			WHERE MatrikelNr = $1 AND TeamName = $2;
		COMMIT;
	`, matrikelNr, team)

	if err != nil {
		panic(err)
	}
}
