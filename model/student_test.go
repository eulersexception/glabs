package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var stud = &Student{
	Name:      "Muster",
	FirstName: "Max",
	NickName:  "Eminem",
	Email:     "mm@example.com",
	Id:        9999,
}

var studs = make([]*Student, 0)

var t = &Team{
	Name:       "TestTeam",
	Assignment: nil,
	Students:   studs,
}

func TestNewStudent(t *testing.T) {
	want := stud
	NewStudent(nil, want.Name, want.FirstName, want.NickName, want.Email, want.Id)
	got, err := GetStudent(want.Id)

	if err != nil {
		t.Errorf("Test failed while fetching data.")
	}

	if !cmp.Equal(want, got) {
		t.Errorf("NewStudent:\nName = '%s', want 'Muster'\nFirstName = '%s', want 'Max'\nNickname = '%s', want 'Eminem'\nEmail = '%s', want 'mm@example.com'\nId = %d, want 9999\n", got.Name, got.FirstName, got.NickName, got.Email, got.Id)
	}
}

func TestNewStudentFailMissingName(t *testing.T) {
	want := "\n+++ Please provide valid name or first name.\n"
	s, got := NewStudent(nil, "", stud.FirstName, stud.NickName, stud.Email, stud.Id)

	if want != got {
		t.Errorf("Expected fail due to missing name")
	}

	if s != nil {
		t.Errorf("Expected result to be nil due to missing name")
	}
}

func TestNewStudentFailMissingFirstName(t *testing.T) {
	want := "\n+++ Please provide valid name or first name.\n"
	s, got := NewStudent(nil, stud.Name, "", stud.NickName, stud.Email, stud.Id)

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

	s, got := NewStudent(nil, stud.Name, stud.FirstName, stud.NickName, malformed, stud.Id)

	if want != got {
		t.Errorf("Expected fail due to malformed mail address '%s'", malformed)
	}

	if s != nil {
		t.Errorf("Expected result to be nil due to malformed mail address '%s'", malformed)
	}
}

func TestDeleteStudent(t *testing.T) {
	got := DeleteStudent(stud.Id)

	if got != nil {
		t.Errorf("Expected nil error as result of delete operation but got %v", got.Error())
	}
}

func TestUpdateStudent(t *testing.T) {
	want := &Student{
		Name:      stud.Name,
		FirstName: stud.FirstName,
		NickName:  stud.NickName,
		Email:     "updated@email.com",
		Id:        stud.Id,
	}
	testObj, _ := NewStudent(nil, stud.Name, stud.FirstName, stud.NickName, stud.Email, stud.Id)
	testObj.Email = "updated@email.com"
	err := testObj.UpdateStudent()

	if err != nil {
		t.Errorf("Test failed while update due to %v", err.Error())
	}

	got, e := GetStudent(stud.Id)

	if e != nil {
		t.Errorf("Test failed while fetching data after update due to %v", e.Error())
	}

	if !cmp.Equal(want, got) {
		t.Errorf("Want %v but got %v", want, got)
	}
}

func TestJoinTeam(t *testing.T) {
	want := t
	team, _ := NewTeam(nil, "TestTeam")

	stud.JoinTeam(team.Name)
	got, _ := GetTeam(team.Name)

	if !cmp.Equal(want, team) {
		t.Errorf("Test failed, want %v but got %v", want, got)
	}

	DeleteTeam(team.Name)
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
