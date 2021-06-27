package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	CreateTables()
}

func TestJoinTeamAndRemoveFromTeam(t *testing.T) {
	testTeam := &Team{Name: "TestStudentTeam"}
	testStudent, _ := NewStudent("Join", "John", "JohnJoin", "John@join.com", 99999)

	want := &Student{Name: testStudent.Name, FirstName: testStudent.FirstName, NickName: testStudent.NickName, Email: testStudent.Email, MatrikelNr: testStudent.MatrikelNr}

	NewStudentTeam(want.MatrikelNr, testTeam.Name)

	got := GetStudentsForTeam(testTeam.Name)[0]

	if want.MatrikelNr != got.MatrikelNr && want.Name != got.Name && want.FirstName != got.FirstName && want.Email != got.Email && want.NickName != got.NickName {
		t.Errorf("want = '%v', got = '%v'\n", want, got)
	}

	RemoveStudentFromTeam(testStudent.MatrikelNr, testTeam.Name)
	studs := GetStudentsForTeam(testTeam.Name)

	for _, v := range studs {
		if want.MatrikelNr == v.MatrikelNr && want.Name == v.Name && want.FirstName == v.FirstName && want.Email == v.Email && want.NickName == v.NickName {
			t.Errorf("want = <nil>, got = '%v'\n", got)
		}
	}

	DeleteTeam(testTeam.Name)
	DeleteStudent(testStudent.MatrikelNr)
}

func TestExistingStudentToTeam(t *testing.T) {
	testTeam := &Team{Name: "TestStudentTeam"}
	testStudent, _ := NewStudent("Join", "John", "JohnJoin", "John@join.com", 99999)

	want := &Student{Name: testStudent.Name, FirstName: testStudent.FirstName, NickName: testStudent.NickName, Email: testStudent.Email, MatrikelNr: testStudent.MatrikelNr}

	NewStudentTeam(want.MatrikelNr, testTeam.Name)
	NewStudentTeam(want.MatrikelNr, testTeam.Name)

	studs := GetStudentsForTeam(testTeam.Name)

	if len(studs) != 1 {
		t.Errorf("want = 1, got =%d\n", len(studs))
	}

	got := studs[0]

	if want.MatrikelNr != got.MatrikelNr && want.Name != got.Name && want.FirstName != got.FirstName && want.Email != got.Email && want.NickName != got.NickName {
		t.Errorf("want = '%v', got = '%v'\n", want, got)
	}

	DeleteTeam(testTeam.Name)
	DeleteStudent(testStudent.MatrikelNr)
}

func TestDeleteTeamWithStudents(t *testing.T) {
	testTeam, _ := NewTeam("TestStudentTeam")
	testStudent, _ := NewStudent("Join", "John", "JohnJoin", "john@join.com", 99999)
	testStudent2, _ := NewStudent("Pan", "Peter", "PeterPan", "peter@pan.com", 99998)

	testTeam.AddStudent(testStudent)
	testTeam.AddStudent(testStudent2)
	DeleteTeam(testTeam.Name)

	want := make([]*Student, 0)
	got := GetStudentsForTeam(testTeam.Name)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteStudent(testStudent.MatrikelNr)
	DeleteStudent(testStudent2.MatrikelNr)
}

func TestDeleteStudentExistingInMultipleTeams(t *testing.T) {
	testTeam, _ := NewTeam("TestStudentTeam")
	testTeam1, _ := NewTeam("TestStudentTeam1")
	testStudent, _ := NewStudent("Join", "John", "JohnJoin", "john@join.com", 99999)

	testTeam.AddStudent(testStudent)
	testTeam1.AddStudent(testStudent)
	DeleteStudent(testStudent.MatrikelNr)

	want := make([]*Team, 0)
	got := GetTeamsForStudent(testStudent.MatrikelNr)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteTeam(testTeam.Name)
	DeleteTeam(testTeam1.Name)
	DeleteStudent(testStudent.MatrikelNr)
}
