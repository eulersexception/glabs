package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	CreateTables()
}

func TestJoinAssignmentRemoveFromAssignment(t *testing.T) {

	starter := &StarterCode{
		Url:             "git@gitlab.lrz.de:vss/startercode/assignmenttest.git",
		FromBranch:      "ws20.1",
		ProtectToBranch: true,
	}

	clone := &Clone{
		LocalPath: "/Users/assignment/test/vss/labs/gitlab/semester/ob-20ws/blatt1",
		Branch:    "develop",
	}

	assignment := &Assignment{
		AssignmentPath:    "blattAssignmentTest",
		SemesterPath:      "semester/xx-20ws",
		Per:               "group",
		Description:       "Blatt AA, Verteilte Softwaresysteme, WS Evergreen",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.Url,
	}

	want, _ := NewTeam("TestTeamAssignment")

	want.JoinAssignment(assignment.AssignmentPath)

	got := GetTeamsForAssignment(assignment.AssignmentPath)[0]

	if want.Name != got.Name {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	assignment.RemoveTeam(want.Name)

	wantTeams := make([]*Team, 0)
	gotTeams := GetTeamsForAssignment(assignment.AssignmentPath)

	if !cmp.Equal(wantTeams, gotTeams) {
		t.Errorf("want = '%v', got = '%v'", wantTeams, gotTeams)
	}

	DeleteTeam(want.Name)
	DeleteAssignment(assignment.AssignmentPath)
}

func TestDeleteAssignmentWithTeams(t *testing.T) {
	starter := &StarterCode{
		Url:             "git@gitlab.lrz.de:vss/startercode/assignmenttest.git",
		FromBranch:      "ws20.1",
		ProtectToBranch: true,
	}

	clone := &Clone{
		LocalPath: "/Users/assignment/test/vss/labs/gitlab/semester/ob-20ws/blatt1",
		Branch:    "develop",
	}

	assignment := &Assignment{
		AssignmentPath:    "blattAssignmentTest",
		SemesterPath:      "semester/xx-20ws",
		Per:               "group",
		Description:       "Blatt AA, Verteilte Softwaresysteme, WS Evergreen",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.Url,
	}

	team1, _ := NewTeam("TestTeamAssignment1")
	team2, _ := NewTeam("TestTeamAssignment2")
	team3, _ := NewTeam("TestTeamAssignment3")

	as, _ := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, clone.Branch, assignment.StarterUrl, starter.FromBranch,
		starter.ProtectToBranch)

	as.AddTeam(team1.Name)
	as.AddTeam(team2.Name)
	as.AddTeam(team3.Name)
	DeleteAssignment(as.AssignmentPath)

	want := make([]*Team, 0)
	got := GetTeamsForAssignment(assignment.AssignmentPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteTeam(team1.Name)
	DeleteTeam(team2.Name)
	DeleteTeam(team3.Name)
}

func TestDeleteTeamWithAssignments(t *testing.T) {
	starter := &StarterCode{
		Url:             "git@gitlab.lrz.de:vss/startercode/assignmenttest.git",
		FromBranch:      "ws20.1",
		ProtectToBranch: true,
	}

	clone := &Clone{
		LocalPath: "/Users/assignment/test/vss/labs/gitlab/semester/ob-20ws/blatt1",
		Branch:    "develop",
	}

	assignment1 := &Assignment{
		AssignmentPath:    "blattTest1",
		SemesterPath:      "semester/1",
		Per:               "group",
		Description:       "Blatt Sem1, Verteilte Softwaresysteme, WS 1990",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.Url,
	}

	assignment2 := &Assignment{
		AssignmentPath:    "blattTest2",
		SemesterPath:      "semester/2",
		Per:               "group",
		Description:       "Blatt Sem2, Verteilte Softwaresysteme, WS 1990",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.Url,
	}

	testTeam, _ := NewTeam("TestTeamAssignmentX")

	as1, _ := NewAssignment(assignment1.AssignmentPath, assignment1.SemesterPath,
		assignment1.Per, assignment1.Description, assignment1.ContainerRegistry,
		assignment1.LocalPath, clone.Branch, assignment1.StarterUrl, starter.FromBranch,
		starter.ProtectToBranch)

	as2, _ := NewAssignment(assignment2.AssignmentPath, assignment2.SemesterPath,
		assignment2.Per, assignment2.Description, assignment2.ContainerRegistry,
		assignment2.LocalPath, clone.Branch, assignment2.StarterUrl, starter.FromBranch,
		starter.ProtectToBranch)

	testTeam.JoinAssignment(as1.AssignmentPath)
	testTeam.JoinAssignment(as2.AssignmentPath)
	DeleteTeam(testTeam.Name)

	want := make([]*Assignment, 0)
	got := GetAssignmentsForTeam(testTeam.Name)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteAssignment(as1.AssignmentPath)
	DeleteAssignment(as2.AssignmentPath)
}
