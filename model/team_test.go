package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewTeamSuccess(t *testing.T) {
	want := &Team{Name: "TestTeam1"}

	NewTeam("TestTeam1")

	got := GetTeam("TestTeam1")

	if !cmp.Equal(want, got) {
		t.Errorf("want = %v, got = %v\n", want, got)
	}
}

func TestNewTeamFail(t *testing.T) {
	want := "\n+++ Please enter a valid team name."

	gotTeam, got := NewTeam("")

	if gotTeam != nil || want != got {
		t.Errorf("want = '%s', got = '%s'\n", want, got)
	}
}

func TestDeleteTeam(t *testing.T) {
	team := &Team{Name: "TestTeam2"}

	NewTeam(team.Name)
	DeleteTeam(team.Name)

	want := &Team{}
	got := GetTeam(team.Name)

	if !cmp.Equal(want, got) {
		t.Errorf("want = %v, got %v", want, got)
	}
}

func TestAddExistingStudent(t *testing.T) {

}
