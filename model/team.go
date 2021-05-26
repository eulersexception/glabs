package model

import (
	"fmt"

	util "github.com/eulersexception/glabs-ui/util"
	DB "modernc.org/ql"
)

// Team - Name is the primary key. All fields are public an

// Getter or Setter functions relate to database operations.
type Team struct {
	TeamID *int64 `ql:"index xID"`
	Name   string `ql:"uindex xName, name TeamName"`
}

// NewTeam creates a new team and stores the object in DB.
// String argument for name must not be empty.
// If a team with given name exists already in DB, the existing dataset will be overwritten.
// Returns a new teamo.
func NewTeam(name string) (*Team, string) {
	if name == "" {
		res := "\n+++ Please enter a valid team name."
		return nil, res
	}

	db := util.GetDB()
	schema := DB.MustSchema((*Team)(nil), "", nil)

	if _, _, e := db.Execute(DB.NewRWCtx(), schema); e != nil {
		panic(e)
	}

	util.FlushAndClose(db)

	team := &Team{
		Name: name,
	}

	team.setTeam()

	return team, ""
}

func (t *Team) AddStudent(s *Student) error {

	return nil
}

func (t Team) RemoveStudent(s Student) error {

	return nil
}

// UpdateTeam changes a teams record in DB.
// Returns an error if the update fails.
func (t *Team) UpdateTeam(newName string) bool {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				UPDATE Team	TeamName = $1 WHERE TeamName = $2;
			COMMIT;
	`, newName, t.Name); err != nil {
		panic(err)
	}

	return true
}

// This function updates team record in DB.
// An update could be a creation or edition of a record.
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

// GetTeam fetches team from DB with an argument of type string as name.
// Returns an error if fetch fails or a pointer to the Team.
func GetTeam(name string) *Team {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	t := &Team{Name: name}

	rss, _, err := db.Run(DB.NewRWCtx(), `
				BEGIN TRANSACTION;
					SELECT TeamID, TeamName FROM Team WHERE TeamName = $1;
				COMMIT;
			`, name)

	if err != nil {
		panic(err)
	}

	t = &Team{}

	for _, rs := range rss {

		if err := rs.Do(false, func(data []interface{}) (bool, error) {

			if err := DB.Unmarshal(t, data); err != nil {
				return false, err
			}

			return true, nil
		}); err != nil {
			panic(err)
		}
	}

	return t
}

// DeleteTeam removes a team by name (string) from DB.
// Returns an error if operation fails.
func DeleteTeam(name string) {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DELETE FROM Team WHERE TeamName == $1;
		COMMIT;
	`, name); err != nil {
		panic(err)
	}
}

func (fst *Team) Equals(scd *Team) bool {
	if scd == nil || fst.Name != scd.Name {
		return false
	}

	return true
}
