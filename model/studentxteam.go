package model

import (
	"github.com/eulersexception/glabs-ui/util"
	DB "modernc.org/ql"
)

type StudentTeam struct {
	MatrikelNr *int64
	TeamName   *string
}

func NewStudentTeam(matrikelNr int64, teamName string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	schema := DB.MustSchema((*StudentTeam)(nil), "", nil)

	if _, _, e := db.Execute(DB.NewRWCtx(), schema); e != nil {
		panic(e)
	}

}

func createNewEntry(matrikelNr int64, teamName string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO StudentTeam IF NOT EXISTS (MatrikelNr, Name) VALUES ($1, $2);
		COMMIT;
	`, matrikelNr, teamName)

	if err != nil {
		panic(err)
	}
}

func GetTeamsForStudent(matrikelNr int64) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	//teams := make([]string, 0)

	if _, _, err := db.Run(DB.NewRWCtx(), `
			SELECT TeamName FROM StudentTeam WHERE MatrikelNr = $1
		`, matrikelNr); err != nil {
		panic(err)
	}

	//for _, rs := range rss {

	//	if e := rs.Do(false, func(data []interface{}) (bool, error) {

	//		return true, nil

	//	}); e != nil {
	//		panic(e)
	//	}
	//}
}

