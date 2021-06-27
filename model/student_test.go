package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var stud = &Student{
	Name:       "Muster",
	FirstName:  "Max",
	NickName:   "Eminem",
	Email:      "mm@example.com",
	MatrikelNr: 9999,
}

func init() {
	CreateTables()
}

func TestMailValid(t *testing.T) {
	got := Mail("valid@mail.com")

	if !got {
		t.Errorf("want = true, got = false\n")
	}
}

func TestMailInvalid(t *testing.T) {
	got := Mail("in valid@mail#com")

	if got {
		t.Error("want = false, got = true\n")
	}
}

func TestNewStudent(t *testing.T) {
	want := stud
	NewStudent(want.Name, want.FirstName, want.NickName, want.Email, want.MatrikelNr)

	got := GetStudent(want.MatrikelNr)
	want.StudentID = got.StudentID

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'\n", want, got)
	}
}

func TestNewStudentAlreadyExists(t *testing.T) {
	want := stud
	wantString := "Student already exists - use update for changes"

	got, gotString := NewStudent(stud.Name, stud.FirstName, stud.NickName, stud.Email, stud.MatrikelNr)

	if !cmp.Equal(want, got) || wantString != gotString {
		t.Errorf("want = '%v', got = '%v'\nwantString = '%s', gotString = '%s'", want, got, wantString, gotString)
	}
}

func TestNewStudentFailMissingName(t *testing.T) {
	want := "\n+++ Enter valid name or first name.\n"

	_, got := NewStudent("", stud.FirstName, stud.NickName, stud.Email, stud.MatrikelNr)

	if want != got {
		t.Errorf("want = '%s', got = '%s'\n", want, got)
	}
}

func TestNewStudentFailMissingFirstName(t *testing.T) {
	want := "\n+++ Enter valid name or first name.\n"

	_, got := NewStudent(stud.Name, "", stud.NickName, stud.Email, stud.MatrikelNr)

	if want != got {
		t.Errorf("want = '%s', got = '%s'\n", want, got)
	}
}

func TestNewStudentFailMalformedMailAddress(t *testing.T) {
	malformed := "bla@blub@blub bla.com"
	want := "\n+++ Enter valid email address.\n"

	_, got := NewStudent(stud.Name, stud.FirstName, stud.NickName, malformed, stud.MatrikelNr)

	if want != got {
		t.Errorf("want = '%s', got = '%s'\n", want, got)
	}
}

func TestDeleteStudent(t *testing.T) {
	DeleteStudent(stud.MatrikelNr)

	want := &Student{}
	got := GetStudent(stud.MatrikelNr)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'\n", want, got)
	}

	NewStudent(stud.Name, stud.FirstName, stud.NickName, stud.Email, stud.MatrikelNr)
}

func TestUpdateStudent(t *testing.T) {
	want := stud
	want.Email = "updated@email.com"
	want.UpdateStudent()

	got := GetStudent(want.MatrikelNr)
	want.StudentID = got.StudentID

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'\n", want, got)
	}

	want.Email = "mm@example.com"
	want.UpdateStudent()
}
