package main

import (
	"fmt"

	//view "github.com/eulersexception/glabs-ui/view"
	"github.com/eulersexception/glabs-ui/model"
	//db "modernc.org/ql"
)

func main() {

	fmt.Println("1. main.go - Creating Student:")
	stud, _ := model.NewStudent("Muster", "Max", "MaxMustermann", "max@muster.com", 99)
	fmt.Printf("\n2. main.go - Fetching Student with Matrikel %d after creation and printing Data:\n", stud.MatrikelNr)
	stud = model.GetStudent(99)
	stud.PrintData()

	fmt.Printf("\n3. main.go - Trying to update Student with Matrikel %d after insertion and printing data:\n", stud.MatrikelNr)
	stud.FirstName = "Maniac"
	stud.UpdateStudent()
	stud = model.GetStudent(99)
	stud.PrintData()

	fmt.Printf("\n4. main.go - Deleting student with Matrikel %d and printing data after get from DB:\n", stud.MatrikelNr)
	model.DeleteStudent(99)
	stud = model.GetStudent(99)

}
