package model

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	CreateTables()
}

func TestJoinTeamAndRemoveFromTeam(t *testing.T) {
	testTeam, _ := NewTeam("TestStudentTeam")
	testStudent, _ := NewStudent("Join", "John", "JohnJoin", "John@join.com", 99999)
	RemoveStudentsForTeam(testTeam.Name)

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

	for _, s := range studs {
		s.PrintData()
	}

	DeleteTeam(testTeam.Name)
	DeleteStudent(testStudent.MatrikelNr)
}

func TestExistingStudentToTeam(t *testing.T) {
	testTeam, _ := NewTeam("TestStudentTeam")
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

	RemoveStudentsForTeam(testTeam.Name)
	DeleteTeam(testTeam.Name)
	DeleteStudent(testStudent.MatrikelNr)
}

func TestDeleteTeamWithStudents(t *testing.T) {
	testTeam, _ := NewTeam("TestStudentTeam")
	testStudent, _ := NewStudent("Join", "John", "JohnJoin", "john@join.com", 99999)
	testStudent2, _ := NewStudent("Pan", "Peter", "PeterPan", "peter@pan.com", 99998)

	testTeam.AddStudent(testStudent)
	testTeam.AddStudent(testStudent2)
	RemoveStudentsForTeam(testTeam.Name)
	DeleteTeam(testTeam.Name)

	want := make([]Student, 0)
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

	want := make([]Team, 0)
	got := GetTeamsForStudent(testStudent.MatrikelNr)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	RemoveTeamsForStudent(testStudent.MatrikelNr)
	DeleteTeam(testTeam.Name)
	DeleteTeam(testTeam1.Name)
	DeleteStudent(testStudent.MatrikelNr)
}

func TestUpdateTeamNameForStudents(t *testing.T) {
	testTeam, _ := NewTeam("TestStudentTeam")
	testStudent, _ := NewStudent("Join", "John", "JohnJoin", "john@join.com", 9995)
	testStudent2, _ := NewStudent("Pan", "Peter", "PeterPan", "peter@pan.com", 99996)

	testTeam.AddStudent(testStudent)
	testTeam.AddStudent(testStudent2)

	want := &Team{Name: "TestStudentTeam2"}
	testTeam.UpdateTeam(want.Name)

	got := GetTeam(want.Name)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	wantedStudents := make([]Student, 0)
	wantedStudents = append(wantedStudents, *testStudent, *testStudent2)
	gotStudents := GetStudentsForTeam(want.Name)
	sort.Slice(gotStudents, func(i int, j int) bool { return gotStudents[i].MatrikelNr < gotStudents[j].MatrikelNr })

	for i, v := range wantedStudents {
		if !cmp.Equal(v, gotStudents[i]) {
			t.Errorf("want = '%v', got = '%v'", v, gotStudents[i])
		}
	}

	DeleteStudent(testStudent.MatrikelNr)
	DeleteStudent(testStudent2.MatrikelNr)
	DeleteTeam(testTeam.Name)
}

func TestUpdateMatrikelNummer(t *testing.T) {
	testTeam1, _ := NewTeam("TestStudentTeam1")
	testTeam2, _ := NewTeam("TestStudentTeam2")
	testStudent, _ := NewStudent("Peter", "North", "PeNo", "peter@north.com", 9999)

	testTeam1.AddStudent(testStudent)
	testTeam2.AddStudent(testStudent)

	UpdateMatrikelNummer(testStudent.MatrikelNr, 11111)
	testStudent.MatrikelNr = 11111
	want := testStudent
	got := GetStudent(11111)

	if !cmp.Equal(want, got) {
		t.Errorf("want = %v, got = %v", want, got)
	}

	wantTeams := make([]Team, 0)
	wantTeams = append(wantTeams, *testTeam1, *testTeam2)
	gotTeams := GetTeamsForStudent(want.MatrikelNr)
	sort.Slice(gotTeams, func(i int, j int) bool { return gotTeams[i].Name < gotTeams[j].Name })

	for i, v := range wantTeams {
		if !cmp.Equal(v, gotTeams[i]) {
			t.Errorf("want = %v, got %v", v, gotTeams[i])
		}
	}

	DeleteStudent(want.MatrikelNr)
	DeleteTeam(testTeam1.Name)
	DeleteTeam(testTeam2.Name)
}
