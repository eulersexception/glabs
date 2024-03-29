package model

import (
	DB "modernc.org/ql"
)

func CreateTables() {
	db := GetDB()
	defer FlushAndClose(db)

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

	schemaTeamAssignment := DB.MustSchema((*TeamAssignment)(nil), "", nil)
	if _, _, e := db.Execute(DB.NewRWCtx(), schemaTeamAssignment); e != nil {
		panic(e)
	}

	schemaSemester := DB.MustSchema((*Semester)(nil), "", nil)
	if _, _, e := db.Execute(DB.NewRWCtx(), schemaSemester); e != nil {
		panic(e)
	}

	schemaCourse := DB.MustSchema((*Course)(nil), "", nil)
	if _, _, e := db.Execute(DB.NewRWCtx(), schemaCourse); e != nil {
		panic(e)
	}
}

func DropTables() {
	db := GetDB()

	if _, _, e := db.Run(DB.NewRWCtx(), `
		BEGIN TRANSACTION;
			DROP TABLE IF EXISTS Student;
			DROP TABLE IF EXISTS Team;
			DROP TABLE IF EXISTS StudentTeam;
			DROP TABLE IF EXISTS Assignment;
			DROP TABLE IF EXISTS TeamAssignment;
			DROP TABLE IF EXISTS StarterCode;
			DROP TABLE IF EXISTS Clone;
			DROP TABLE IF EXISTS Semester;
			DROP TABLE IF EXISTS Course;
		COMMIT;
	`); e != nil {
		panic(e)
	}

	FlushAndClose(db)
}

func InitData() {

	// Dummy courses
	vss := NewCourse("vss")
	se := NewCourse("se")
	algoDat := NewCourse("algodat-i")

	// Dummy semesters
	vssSem1 := NewSemester(vss.Path, "semester/ob-20ws")
	vssSem2 := NewSemester(vss.Path, "semester/ob-21ss")
	seSem1 := NewSemester(se.Path, "semester/rs-20ws")
	seSem2 := NewSemester(se.Path, "semester/rs-21ss")
	algoDatSem1 := NewSemester(algoDat.Path, "semester/ob-19ws")

	// Dummy starter codes
	starter1, _ := NewStarterCode("git@gitlab.lrz.de:algodat/startercode/startercodeA1.git", "ws20.1", true)
	starter2, _ := NewStarterCode("git@gitlab.lrz.de:vss/startercode/startercodeA2.git", "ws20.2", true)
	starter3, _ := NewStarterCode("git@gitlab.lrz.de:se1/startercode/startercodeB1.git", "ss18.1", false)
	starter4, _ := NewStarterCode("git@gitlab.lrz.de:se2/startercode/startercodeB1.git", "ss19.1", false)

	// Dummy clones
	clone1, _ := NewClone("/Users/obraun/lectures/algodat/labs/gitlab/semester/ob-20ws/blatt1", "develop")
	clone2, _ := NewClone("/Users/obraun/lectures/vss/labs/gitlab/semester/ob-21ws/blatt1", "develop")
	clone3, _ := NewClone("/Users/rs/lectures/se1/labs/gitlab/semester/ob-18ss/blatt1", "master")
	clone4, _ := NewClone("/Users/rs/lectures/se2/labs/gitlab/semester/ob-19ss/blatt1", "master")

	// Dummy assignments
	vssSem1Blatt1, _ := NewAssignment("vssSem1Blatt1",
		vssSem1.Path,
		"group",
		"Blatt 1, Verteilte Softwaresysteme, WS 20/21",
		true,
		clone1.LocalPath,
		starter1.StarterUrl,
	)

	vssSem1Blatt2, _ := NewAssignment("vssSem1Blatt2",
		vssSem1.Path,
		"group",
		"Blatt 2, Verteilte Softwaresysteme, WS 20/21",
		true,
		clone2.LocalPath,
		starter2.StarterUrl,
	)

	vssSem2Blatt1, _ := NewAssignment("vssSem2Blatt1",
		vssSem2.Path,
		"group",
		"Blatt 1, Verteilte Softwaresysteme, SoSe 21",
		true,
		clone3.LocalPath,
		starter3.StarterUrl,
	)

	vssSem2Blatt2, _ := NewAssignment("vssSem2Blatt2",
		vssSem2.Path,
		"group",
		"Blatt 2, Verteilte Softwaresysteme, SoSe 21",
		true,
		clone4.LocalPath,
		starter4.StarterUrl,
	)

	seSem1Blatt1, _ := NewAssignment("seSem1Blatt1",
		seSem1.Path,
		"group",
		"Blatt 1, Softwareentwicklung, WS 20/21",
		true,
		clone1.LocalPath,
		starter1.StarterUrl,
	)

	seSem1Blatt2, _ := NewAssignment("seSem1Blatt2",
		seSem1.Path,
		"group",
		"Blatt 2, Softwareentwicklung, WS 20/21",
		true,
		clone2.LocalPath,
		starter2.StarterUrl,
	)

	seSem2Blatt1, _ := NewAssignment("seSem2Blatt1",
		seSem2.Path,
		"group",
		"Blatt 1, Softwareentwicklung, SoSe 21",
		true,
		clone3.LocalPath,
		starter3.StarterUrl,
	)

	seSem2Blatt2, _ := NewAssignment("seSem2Blatt2",
		seSem2.Path,
		"group",
		"Blatt 2, Verteilte Softwaresysteme, SoSE 21",
		true,
		clone4.LocalPath,
		starter4.StarterUrl,
	)

	algoDatBlatt1, _ := NewAssignment("algodatBlatt1",
		algoDatSem1.Path,
		"group",
		"Blatt 1, Algorithmen und Datenstrukturen, WS 20/21",
		true,
		clone1.LocalPath,
		starter1.StarterUrl,
	)

	algoDatBlatt2, _ := NewAssignment("algodatBlatt2",
		algoDatSem1.Path,
		"group",
		"Blatt 2, Algorithmen und Datenstrukturen, WS 20/21",
		true,
		clone4.LocalPath,
		starter4.StarterUrl,
	)

	team1, _ := NewTeam("Team1")
	team2, _ := NewTeam("Team2")
	team3, _ := NewTeam("Team3")
	team4, _ := NewTeam("Team4")

	stud1, _ := NewStudent("Tom", "Tailor", "Tick", "tom@hm.de", 1111)
	stud2, _ := NewStudent("Tim", "Tailor", "Trick", "tim@hm.de", 2222)
	stud3, _ := NewStudent("Tina", "Tailor", "Track", "tinam@hm.de", 3333)
	stud4, _ := NewStudent("Ying", "Yong", "yin", "ying@hm.de", 4444)
	stud5, _ := NewStudent("Yang", "Yong", "yan", "yang@hm.de", 5555)
	stud6, _ := NewStudent("Lili", "Lang", "Li", "lilly@hm.de", 6666)
	stud7, _ := NewStudent("Lala", "Lang", "La", "lala@hm.de", 7777)
	stud8, _ := NewStudent("Lulu", "Lang", "Lu", "lulu@hm.de", 8888)

	vssSem1Blatt1.AddTeam("Team1")
	vssSem1Blatt1.AddTeam("Team2")
	vssSem1Blatt1.AddTeam("Team3")
	vssSem1Blatt1.AddTeam("Team4")

	vssSem1Blatt2.AddTeam("Team1")
	vssSem1Blatt2.AddTeam("Team2")
	vssSem1Blatt2.AddTeam("Team3")
	vssSem1Blatt2.AddTeam("Team4")

	vssSem2Blatt1.AddTeam("Team1")
	vssSem2Blatt1.AddTeam("Team2")
	vssSem2Blatt1.AddTeam("Team3")
	vssSem2Blatt1.AddTeam("Team4")

	vssSem2Blatt2.AddTeam("Team1")
	vssSem2Blatt2.AddTeam("Team2")
	vssSem2Blatt2.AddTeam("Team3")
	vssSem2Blatt2.AddTeam("Team4")

	seSem1Blatt1.AddTeam("Team1")
	seSem1Blatt1.AddTeam("Team2")
	seSem1Blatt1.AddTeam("Team3")
	seSem1Blatt1.AddTeam("Team4")

	seSem1Blatt2.AddTeam("Team1")
	seSem1Blatt2.AddTeam("Team2")
	seSem1Blatt2.AddTeam("Team3")
	seSem1Blatt2.AddTeam("Team4")

	seSem2Blatt1.AddTeam("Team1")
	seSem2Blatt1.AddTeam("Team2")
	seSem2Blatt1.AddTeam("Team3")
	seSem2Blatt1.AddTeam("Team4")

	seSem2Blatt2.AddTeam("Team1")
	seSem2Blatt2.AddTeam("Team2")
	seSem2Blatt2.AddTeam("Team3")
	seSem2Blatt2.AddTeam("Team4")

	algoDatBlatt1.AddTeam("Team1")
	algoDatBlatt1.AddTeam("Team2")

	algoDatBlatt2.AddTeam("Team1")
	algoDatBlatt2.AddTeam("Team2")

	team1.AddStudent(stud1)
	team1.AddStudent(stud2)

	team2.AddStudent(stud3)
	team2.AddStudent(stud4)

	team3.AddStudent(stud5)
	team3.AddStudent(stud6)

	team4.AddStudent(stud7)
	team4.AddStudent(stud8)
}
