package model

import (
	"log"
	"testing"

	"github.com/dgraph-io/badger/v3"

)

var testStarter = &StarterCode{
	Url:             "www.test.com",
	FromBranch:      "develop",
	ProtectToBranch: true,
}

var testClone = &Clone{
	LocalPath: "local",
	Branch:    "master",
}

var testAssignment = &Assignment{
	Name:              "TestAssignment",
	Semester:          nil,
	Teams:             nil,
	LocalClone:        testClone,
	Starter:           testStarter,
	ContainerRegistry: true,
}

var team = &Team{
	Name:       "TestTeam",
	Assignment: testAssignment,
	Students:   nil,
}

var studOne = &Student{
	Name:      "Minogue",
	FirstName: "Kylie",
	NickName:  "kymi",
	Email:     "kymi@example.com",
	Id:        10000,
}

var studTwo = &Student{
	Name:      "Simone",
	FirstName: "Nina",
	NickName:  "nisi",
	Email:     "nisi@example.com",
	Id:        10001,
}

func TestNewTeamSuccess(t *testing.T) {
	want := team
	NewTeam(team.Assignment, team.Name)
	got, _ := GetTeam(team.Name)

	if !want.Equals(got) {
		t.Errorf("Test failed, want %v but got %v", want, got)
	}
}

func TestNewTeamFail(t *testing.T) {
	want := "\n+++ Please enter a valid team name."
	team, got := NewTeam(nil, "")

	if want != got {
		t.Errorf("Test failed, expected %s but got %s", want, got)
	}

	if team != nil {
		t.Errorf("Expected result to be nil due to missing team name")
	}
}

func TestGetTeam(t *testing.T) {
	want := team
	got, err := GetTeam(team.Name)

	if err != nil {
		log.Fatal(err)
	}

	if !want.Equals(got) {
		t.Errorf("Test failed, want %v but got %v", want, got)
	}
}

func TestDeleteTeam(t *testing.T) {
	indb, _ := GetTeam(team.Name)
	got := DeleteTeam(team.Name)

	if indb == nil {
		t.Errorf("Test failed no db entry")
	}

	if got != nil {
		t.Errorf("Couldn't delete team %s.\nCheck error: %v", team.Name, got.Error())
	}
}

func TestAddExistingStudent(t *testing.T) {
	studentOne, _ := NewStudent(nil, studOne.Name, studOne.FirstName, studOne.NickName, studOne.Email, studOne.Id)
	got, _ := NewTeam(nil, team.Name)
	want := team
	studOne.Team = team

	want.Assignment = nil

	want.Students = append(want.Students, studOne)
	got.AddStudent(studentOne)

	if len(want.Students) != len(got.Students) {
		t.Errorf("Test failed for length comparison, want %d (number of students in team) but got %d", len(want.Students), len(got.Students))
	}

	for i, g := range got.Students {
		w := want.Students[i]
		if !w.Equals(g) {
			t.Errorf("Test failed for student comparison, want:\n\tname: %s, %s //\n\tfirst: %s, %s\n\tnick: %s, %s\n\temail: %s, %s\n\tid: %d, %d\n\tteam: %s, %s", w.Name, g.Name, w.FirstName, g.FirstName, w.NickName, g.NickName, w.Email, g.Email, w.Id, g.Id, w.Team.Name, g.Team.Name)
		}
	}

	if want.Name != got.Name {
		t.Errorf("Test failed for id comparison, want %s but got %s", want.Name, got.Name)
	}
}
