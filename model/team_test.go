package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	CreateTables()
}

func TestNewTeamSuccess(t *testing.T) {
	want := &Team{Name: "TestTeam1"}

	NewTeam("TestTeam1")

	got := GetTeam("TestTeam1")
	want.TeamID = got.TeamID

	if !cmp.Equal(want, got) {
		t.Errorf("want = %v, got = %v\n", want, got)
	}
}

func TestNewTeamFail(t *testing.T) {
	want := "\n+++ Enter valid team name."

	gotTeam, got := NewTeam("")

	if gotTeam != nil || want != got {
		t.Errorf("want = '%s', got = '%s'\n", want, got)
	}
}

func TestNewTeamAlreadyExists(t *testing.T) {
	want := GetTeam("TestTeam1")
	wantString := "Team already exists - use update for changes"

	got, gotString := NewTeam("TestTeam1")

	if !cmp.Equal(want, got) || wantString != gotString {
		t.Errorf("want = '%v', got = '%v'\nwantString = '%s', gotString = '%s'", want, got, wantString, gotString)
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
