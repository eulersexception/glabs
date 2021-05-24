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

var studs = make([]*Student, 0)

var teamOne = &Team{
	Name: "TestTeam",
}

func TestNewStudent(t *testing.T) {
	want := stud
	NewStudent(want.Name, want.FirstName, want.NickName, want.Email, want.MatrikelNr)
	got := GetStudent(want.MatrikelNr)

	if !cmp.Equal(want, got) {
		t.Errorf("NewStudent:\nName = '%s', want 'Muster'\nFirstName = '%s', want 'Max'\nNickname = '%s', want 'Eminem'\nEmail = '%s', want 'mm@example.com'\nId = %d, want 9999\n",
			got.Name, got.FirstName, got.NickName, got.Email, got.MatrikelNr)
	}
}

func TestNewStudentFailMissingName(t *testing.T) {
	want := "\n+++ Please provide valid name or first name.\n"
	s, got := NewStudent("", stud.FirstName, stud.NickName, stud.Email, stud.MatrikelNr)

	if want != got {
		t.Errorf("Expected fail due to missing name")
	}

	if s != nil {
		t.Errorf("Expected result to be nil due to missing name")
	}
}

func TestNewStudentFailMissingFirstName(t *testing.T) {
	want := "\n+++ Please provide valid name or first name.\n"
	s, got := NewStudent(stud.Name, "", stud.NickName, stud.Email, stud.MatrikelNr)

	if want != got {
		t.Errorf("Expected fail due to missing first name")
	}

	if s != nil {
		t.Errorf("Expected result to be nil due to missing first name")
	}
}

func TestNewStudentFailMalformedMailAddress(t *testing.T) {
	malformed := "bla@blub@blub bla.com"
	want := "\n+++ Please provide valid email address.\n"

	s, got := NewStudent(stud.Name, stud.FirstName, stud.NickName, malformed, stud.MatrikelNr)

	if want != got {
		t.Errorf("Expected fail due to malformed mail address '%s'", malformed)
	}

	if s != nil {
		t.Errorf("Expected result to be nil due to malformed mail address '%s'", malformed)
	}
}

func TestDeleteStudent(t *testing.T) {
	DeleteStudent(stud.MatrikelNr)

	got := GetStudent(stud.MatrikelNr)

	if got.StudentID != nil {
		t.Errorf("Expected nil as result of get after delete operation but got %v", got)
	}

	NewStudent(stud.Name, stud.FirstName, stud.NickName, stud.Email, stud.MatrikelNr)
}

func TestUpdateStudent(t *testing.T) {
	want := &Student{
		Name:       stud.Name,
		FirstName:  stud.FirstName,
		NickName:   stud.NickName,
		Email:      "updated@email.com",
		MatrikelNr: stud.MatrikelNr,
	}

	testObj, _ := NewStudent(stud.Name, stud.FirstName, stud.NickName, stud.Email, stud.MatrikelNr)
	testObj.Email = "updated@email.com"
	result := testObj.UpdateStudent()

	if !result {
		t.Errorf("Test failed while update - return value = false")
	}

	got := GetStudent(stud.MatrikelNr)

	if !want.Equals(got) {
		t.Errorf("Want %v but got %v", want, got)
	}

	testObj.Email = "mm@example.com"
	testObj.UpdateStudent()
}

func TestJoinTeam(t *testing.T) {

}

func TestMailValid(t *testing.T) {
	got := Mail("valid@mail.com")

	if !got {
		t.Errorf("Test failed for mail check. Want 'true' but got 'false'")
	}
}

func TestMailInvalid(t *testing.T) {
	got := Mail("in valid@mail#com")

	if got {
		t.Errorf("Test failed for mail check. Want 'false' but got 'true'")
	}
}
