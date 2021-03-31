package model

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
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

	if !cmp.Equal(want, got) {
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

	if !cmp.Equal(want, got) {
		t.Errorf("Test failed, want %v but got %v", want, got)
	}
}

func TestDeleteTeam(t *testing.T) {
	got := DeleteTeam(team.Name)

	if got != nil {
		t.Errorf("Couldn't delete team %s.\nCheck error: %v", team.Name, got.Error())
	}
}