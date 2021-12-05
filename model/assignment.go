package model

import (
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

func NewAssignment(assignmentPath string, sem string, per string, desc string,
	conRegistry bool, localPath string, starterUrl string) (*Assignment, string) {

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

	return assignment, ""
}

func (a Assignment) setAssignment() {
	db := GetDB()
	defer FlushAndClose(db)

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
	db := GetDB()
	defer FlushAndClose(db)

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

func (a *Assignment) UpdateAssignment() {
	db := GetDB()
	defer FlushAndClose(db)

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

func UpdateAssignmentPath(oldPath string, newPath string) {
	UpdateAssignmentForTeams(oldPath, newPath)
	db := GetDB()
	defer FlushAndClose(db)

	_, _, err := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			UPDATE Assignment AssignmentPath = $1 WHERE AssignmentPath = $2;
		COMMIT;
	`, newPath, oldPath)

	if err != nil {
		panic(err)
	}
}

func DeleteAssignment(path string) {
	db := GetDB()
	defer FlushAndClose(db)

	if _, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;			
				DELETE FROM Assignment WHERE AssignmentPath = $1;
				DELETE FROM TeamAssignment WHERE AssignmentPath = $1;
			COMMIT;
		`, path); err != nil {
		panic(err)
	}
}

func GetAllAssignmentsForSemester(semesterPath string) []Assignment {
	db := GetDB()

	rss, _, e := db.Run(DB.NewRWCtx(), `
		SELECT * FROM Assignment WHERE SemesterPath = $1;
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

	FlushAndClose(db)

	return assignments
}

func (as *Assignment) AddTeam(name string) {
	NewTeamAssignment(name, as.AssignmentPath)
}

func (as *Assignment) RemoveTeam(name string) {
	RemoveTeamFromAssignment(name, as.AssignmentPath)
}
