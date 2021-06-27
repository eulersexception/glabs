package model

import (
	"fmt"

	"github.com/eulersexception/glabs-ui/util"
	DB "modernc.org/ql"
)

type StudentTeam struct {
	MatrikelTeam string `ql:"uindex xMatrikelTeam"`
	MatrikelNr   int64
	TeamName     string
}

func NewStudentTeam(matrikelNr int64, teamName string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

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

func GetTeamsForStudent(matrikelNr int64) []*Team {
	db := util.GetDB()

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

	util.FlushAndClose(db)

	teams := make([]*Team, 0)

	for _, v := range entries {
		teams = append(teams, GetTeam(v.TeamName))
	}

	return teams
}

func GetStudentsForTeam(team string) []*Student {
	db := util.GetDB()

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

	util.FlushAndClose(db)

	studs := make([]*Student, 0)

	for _, v := range entries {
		studs = append(studs, GetStudent(v.MatrikelNr))
	}

	return studs
}

func RemoveStudentFromTeam(matrikelNr int64, team string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

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
