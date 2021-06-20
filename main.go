package main

import (
	//view "github.com/eulersexception/glabs-ui/view"
	"fmt"

	"github.com/eulersexception/glabs-ui/model"
	//db "modernc.org/ql"
)

func main() {
	s1, _ := model.NewStudent("a", "b", "ab", "ab@mail.com", 33)
	s2, _ := model.NewStudent("c", "d", "cd", "cd@mail.com", 44)
	s3, _ := model.NewStudent("e", "f", "ef", "ef@mail.com", 55)

	t1, _ := model.NewTeam("Team1")
	t2, _ := model.NewTeam("Team2")
	t3, _ := model.NewTeam("Team3")

	model.NewStudentTeam(s1.MatrikelNr, t1.Name)
	model.NewStudentTeam(s2.MatrikelNr, t1.Name)
	model.NewStudentTeam(s3.MatrikelNr, t1.Name)

	model.NewStudentTeam(s1.MatrikelNr, t2.Name)
	model.NewStudentTeam(s1.MatrikelNr, t3.Name)

	teams := model.GetTeamsForStudent(s1.MatrikelNr)

	for _, v := range teams {
		fmt.Printf("teams - main.go 27 = %v\n", v)
	}

	studs := model.GetStudentsForTeam(t1.Name)

	for _, v := range studs {
		fmt.Printf("students - main.go 36 = %v\n", v)
	}

	model.DeleteStudent(s1.MatrikelNr)

	teams = model.GetTeamsForStudent(s1.MatrikelNr)

	fmt.Println("Teams array after deleting s1")
	for _, v := range teams {
		fmt.Printf("teams - main.go 34 = %v\n", v)
	}
}
