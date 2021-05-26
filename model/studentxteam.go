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
	schema := DB.MustSchema((*StudentTeam)(nil), "", nil)

	if _, _, e := db.Execute(DB.NewRWCtx(), schema); e != nil {
		panic(e)
	}

	util.FlushAndClose(db)

	createNewEntry(matrikelNr, teamName)
}

func createNewEntry(matrikelNr int64, teamName string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	id := fmt.Sprintf("%d%s", matrikelNr, teamName)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO StudentTeam IF NOT EXISTS (MatrikelTeam, MatrikelNr, TeamName) VALUES ($1, $2, $3);
		COMMIT;
	`, id, matrikelNr, teamName)

	if err != nil {
		panic(err)
	}
}

func GetTeamsForStudent(matrikelNr int64) []Team {
	db := util.GetDB()

	rss, _, err := db.Run(DB.NewRWCtx(), `
			SELECT * FROM StudentTeam WHERE MatrikelNr = $1
		`, matrikelNr)

	if err != nil {
		panic(err)
	}

	entries := make([]*StudentTeam, 0)

	for _, rs := range rss {
		s := &StudentTeam{}

		if e := rs.Do(false, func(data []interface{}) (bool, error) {

			if err = DB.Unmarshal(s, data); err != nil {
				return false, err
			}

			entries = append(entries, s)
			return true, nil
		}); e != nil {
			panic(e)
		}
	}

	util.FlushAndClose(db)

	teams := make([]Team, 0)

	for _, v := range entries {
		teams = append(teams, *GetTeam(v.TeamName))
	}

	return teams
}
