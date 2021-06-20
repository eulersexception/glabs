package model

import (
	"github.com/eulersexception/glabs-ui/util"
	DB "modernc.org/ql"
)

func CreateTables() {
	db := util.GetDB()
	defer util.FlushAndClose(db)

	// Create table Student
	schemaStudent := DB.MustSchema((*Student)(nil), "", nil)

	if _, _, e := db.Execute(DB.NewRWCtx(), schemaStudent); e != nil {
		panic(e)
	}

	schemaTeam := DB.MustSchema((*Team)(nil), "", nil)

	if _, _, e := db.Execute(DB.NewRWCtx(), schemaTeam); e != nil {
		panic(e)
	}

	schemaStudentTeam := DB.MustSchema((*StudentTeam)(nil), "", nil)

	if _, _, e := db.Execute(DB.NewRWCtx(), schemaStudentTeam); e != nil {
		panic(e)
	}

	schemaAssignment := DB.MustSchema((*Assignment)(nil), "", nil)

	if _, _, e := db.Execute(DB.NewRWCtx(), schemaAssignment); e != nil {
		panic(e)
	}

	schemaStarterCode := DB.MustSchema((*StarterCode)(nil), "", nil)

	if _, _, e := db.Execute(DB.NewRWCtx(), schemaStarterCode); e != nil {
		panic(e)
	}

	schemaClone := DB.MustSchema((*Clone)(nil), "", nil)

	if _, _, e := db.Execute(DB.NewRWCtx(), schemaClone); e != nil {
		panic(e)
	}

	//schemaSemester := DB.MustSchema((*Semester)(nil), "", nil)
	//schemaCourse := DB.MustSchema((*Course)(nil), "", nil)
}

func DropTables() {
	db := util.GetDB()

	if _, _, e := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DROP TABLE IF EXISTS Student;
			DROP TABLE IF EXISTS Team;
			DROP TABLE IF EXISTS StudentTeam;
			DROP TABLE IF EXISTS Assignment;
			DROP TABLE IF EXISTS StarterCode;
			DROP TABLE IF EXISTS Clone;
			DROP TABLE IF EXISTS Semester;
			DROP TABLE IF EXISTS Course;
		COMMIT;
	`); e != nil {
		panic(e)
	}

	util.FlushAndClose(db)
}
