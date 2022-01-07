package model

import (
	"fmt"

	DB "modernc.org/ql"
)

type StarterCode struct {
	StarterCodeID   *int64 `ql:"index xID"`
	StarterUrl      string `ql:"uindex xStarterUrl"`
	FromBranch      string
	ProtectToBranch bool
}

func NewStarterCode(url string, fromBranch string, protectToBranch bool) (*StarterCode, string) {
	if url == "" {
		return nil, "Enter valid starter url."
	}

	if fromBranch == "" {
		return nil, "Enter valid from branch."
	}

	starter := &StarterCode{
		StarterUrl:      url,
		FromBranch:      fromBranch,
		ProtectToBranch: protectToBranch,
	}

	starter.setStarterCode()

	return starter, ""
}

func (s StarterCode) setStarterCode() {
	db := GetDB()
	defer FlushAndClose(db)

	if _, _, e := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			INSERT INTO StarterCode IF NOT EXISTS (StarterUrl, FromBranch, ProtectToBranch) VALUES ($1, $2, $3);
		COMMIT;
	`, s.StarterUrl, s.FromBranch, s.ProtectToBranch); e != nil {
		panic(e)
	}
}

func GetStarterCode(url string) *StarterCode {
	db := GetDB()
	defer FlushAndClose(db)

	rss, _, e := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				SELECT * FROM StarterCode WHERE  StarterUrl = $1;
			COMMIT;
		`, url)

	if e != nil {
		panic(e)
	}

	s := &StarterCode{}

	for _, rs := range rss {
		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(s, data); err != nil {
				return false, err
			}

			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	return s
}

func GetAllStarterCodes() []StarterCode {
	db := GetDB()
	defer FlushAndClose(db)

	rss, _, e := db.Run(nil, `
	SELECT * FROM StarterCode;
`)

	if e != nil {
		panic(e)
	}

	starters := make([]StarterCode, 0)

	for _, rs := range rss {
		s := &StarterCode{}

		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(s, data); err != nil {
				return false, err
			}

			starters = append(starters, *s)

			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	return starters

}


func GetAllAssignmentsForStarterCode(url string) []Assignment {
	db := GetDB()
	defer FlushAndClose(db)

	rss, _, e := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			SELECT * FROM Assignment WHERE StarterUrl = $1;
		COMMIT;
	`, url)

	if e != nil {
		panic(e)
	}

	assignments := make([]Assignment, 0)

	for _, rs := range rss {
		assignment := &Assignment{}

		if er := rs.Do(false, func(data []interface{}) (bool, error) {
			if err := DB.Unmarshal(assignment, data); err != nil {
				return false, err
			}

			assignments = append(assignments, *assignment)

			return true, nil
		}); er != nil {
			panic(er)
		}
	}

	return assignments
}

func (s *StarterCode) UpdateStarterCode() {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE StarterCode FromBranch = $2, ProtectToBranch = $3
				WHERE StarterUrl = $1;
			COMMIT;
	`, s.StarterUrl, s.FromBranch, s.ProtectToBranch)

	if err != nil {
		panic(err)
	}
}

func UpdateStarterUrl(oldUrl string, newUrl string) {
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			UPDATE StarterCode StarterUrl = $1 WHERE StarterUrl = $2;
			UPDATE Assignment StarterUrl = $1 WHERE StarterUrl = $2;
		COMMIT;
	`, newUrl, oldUrl)

	if err != nil {
		panic(err)
	}
}

func DeleteStarterCode(url string) {
	db := GetDB()
	defer FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM StarterCode WHERE StarterUrl = $1;
			UPDATE Assignment StarterUrl = $2 WHERE StarterUrl = $1;
		COMMIT;
	`, url, ""); err != nil {
		panic(err)
	}
}

func (s StarterCode) toString() string {
	return fmt.Sprintf("\tStarterCode:\n\t\tUrl:\t%s\n\t\tFromBranch:\t%s\n\t\tProtectToBranch:\t%v", s.StarterUrl, s.FromBranch, s.ProtectToBranch)
}
