package model

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	CreateTables()
}

func TestJoinAssignmentRemoveFromAssignment(t *testing.T) {
	assignment := &Assignment{
		AssignmentPath:    "blattAssignmentTest",
		SemesterPath:      "semester/xx-20ws",
		Per:               "group",
		Description:       "Blatt AA, Verteilte Softwaresysteme, WS Evergreen",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.StarterUrl,
	}

	RemoveTeamsForAssignment(assignment.AssignmentPath)

	want, _ := NewTeam("TestTeamAssignment")
	want.JoinAssignment(assignment.AssignmentPath)

	got := GetTeamsForAssignment(assignment.AssignmentPath)[0]

	if want.Name != got.Name {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	assignment.RemoveTeam(want.Name)

	wantTeams := make([]Team, 0)
	gotTeams := GetTeamsForAssignment(assignment.AssignmentPath)

	if !cmp.Equal(wantTeams, gotTeams) {
		t.Errorf("want = '%v', got = '%v'", wantTeams, gotTeams)
	}

	DeleteTeam(want.Name)
	DeleteAssignment(assignment.AssignmentPath)
}

func TestDeleteAssignmentWithTeams(t *testing.T) {
	assignment := &Assignment{
		AssignmentPath:    "blattAssignmentTest",
		SemesterPath:      "semester/xx-20ws",
		Per:               "group",
		Description:       "Blatt AA, Verteilte Softwaresysteme, WS Evergreen",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.StarterUrl,
	}

	RemoveTeamsForAssignment(assignment.AssignmentPath)

	team1, _ := NewTeam("TestTeamAssignment1")
	team2, _ := NewTeam("TestTeamAssignment2")
	team3, _ := NewTeam("TestTeamAssignment3")

	as, _ := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, assignment.StarterUrl)

	as.AddTeam(team1.Name)
	as.AddTeam(team2.Name)
	as.AddTeam(team3.Name)

	RemoveTeamsForAssignment(as.AssignmentPath)
	DeleteAssignment(as.AssignmentPath)

	want := make([]Team, 0)
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
		StarterUrl:      "git@gitlab.lrz.de:vss/startercode/assignmenttest.git",
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
		StarterUrl:        starter.StarterUrl,
	}

	assignment2 := &Assignment{
		AssignmentPath:    "blattTest2",
		SemesterPath:      "semester/2",
		Per:               "group",
		Description:       "Blatt Sem2, Verteilte Softwaresysteme, WS 1990",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.StarterUrl,
	}

	testTeam, _ := NewTeam("TestTeamAssignmentX")

	as1, _ := NewAssignment(assignment1.AssignmentPath, assignment1.SemesterPath,
		assignment1.Per, assignment1.Description, assignment1.ContainerRegistry,
		assignment1.LocalPath, assignment1.StarterUrl)

	as2, _ := NewAssignment(assignment2.AssignmentPath, assignment2.SemesterPath,
		assignment2.Per, assignment2.Description, assignment2.ContainerRegistry,
		assignment2.LocalPath, assignment2.StarterUrl)

	testTeam.JoinAssignment(as1.AssignmentPath)
	testTeam.JoinAssignment(as2.AssignmentPath)
	DeleteTeam(testTeam.Name)

	want := make([]Assignment, 0)
	got := GetAssignmentsForTeam(testTeam.Name)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteAssignment(as1.AssignmentPath)
	DeleteAssignment(as2.AssignmentPath)
}

func TestUpdateAssignmentPath(t *testing.T) {
	assignment := &Assignment{
		AssignmentPath:    "blattAssignmentTest",
		SemesterPath:      "semester/xx-20ws",
		Per:               "group",
		Description:       "Blatt AA, Verteilte Softwaresysteme, WS Evergreen",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.StarterUrl,
	}

	DeleteAssignment(assignment.AssignmentPath)
	DeleteAssignment("NewPath")

	team1, _ := NewTeam("TestTeamAssignment1")
	team2, _ := NewTeam("TestTeamAssignment2")
	team3, _ := NewTeam("TestTeamAssignment3")

	want, _ := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, assignment.StarterUrl)

	want.AddTeam(team1.Name)
	want.AddTeam(team2.Name)
	want.AddTeam(team3.Name)

	UpdateAssignmentPath(want.AssignmentPath, "NewPath")
	want.AssignmentPath = "NewPath"
	got := GetAssignment(want.AssignmentPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = %v, got = %v", want, got)
	}

	wantTeams := make([]Team, 0)
	wantTeams = append(wantTeams, *team1, *team2, *team3)
	gotTeams := GetTeamsForAssignment(want.AssignmentPath)
	sort.Slice(gotTeams, func(i int, j int) bool { return gotTeams[i].Name < gotTeams[j].Name })

	for i, v := range wantTeams {
		if !cmp.Equal(v, gotTeams[i]) {
			t.Errorf("want = %v, got = %v", v, gotTeams[i])
		}
	}

	DeleteTeam(team1.Name)
	DeleteTeam(team2.Name)
	DeleteTeam(team3.Name)
	DeleteAssignment(want.AssignmentPath)
}

func TestGetAssignmentsForTeam(t *testing.T) {
	a1, _ := NewAssignment("PathA", "SemesterX", "group", "BlattA", true, "TestCloneA", "TestStarterA")
	a2, _ := NewAssignment("PathB", "SemesterY", "single", "BlattB", false, "TestCloneB", "TestStarterB")
	a3, _ := NewAssignment("PathC", "SemesterZ", "pair", "BlattC", true, "TestCloneC", "TestStarterC")

	team, _ := NewTeam("TestGetAssignments")

	team.JoinAssignment(a1.AssignmentPath)
	team.JoinAssignment(a2.AssignmentPath)
	team.JoinAssignment(a3.AssignmentPath)

	want := make([]Assignment, 0)
	want = append(want, *a1, *a2, *a3)
	got := GetAssignmentsForTeam(team.Name)
	sort.Slice(got, func(i int, j int) bool { return got[i].AssignmentPath < got[j].AssignmentPath })

	for i, v := range want {
		if !cmp.Equal(v, got[i]) {
			t.Errorf("want = %v, got = %v", v, got[i])
		}
	}

	DeleteAssignment(a1.AssignmentPath)
	DeleteAssignment(a2.AssignmentPath)
	DeleteAssignment(a3.AssignmentPath)
	DeleteTeam(team.Name)
}

func TestUpdateTeam(t *testing.T) {
	a1, _ := NewAssignment("PathA", "SemesterX", "group", "BlattA", true, "TestCloneA", "TestStarterA")
	a2, _ := NewAssignment("PathB", "SemesterY", "single", "BlattB", false, "TestCloneB", "TestStarterB")
	a3, _ := NewAssignment("PathC", "SemesterZ", "pair", "BlattC", true, "TestCloneC", "TestStarterC")

	team, _ := NewTeam("TestUpdateTeamBefore")

	team.JoinAssignment(a1.AssignmentPath)
	team.JoinAssignment(a2.AssignmentPath)
	team.JoinAssignment(a3.AssignmentPath)
	newName := "TestUpdateTeamAfter"
	team.UpdateTeam(newName)
	team.Name = newName

	want := make([]Assignment, 0)
	want = append(want, *a1, *a2, *a3)
	got := GetAssignmentsForTeam(team.Name)
	sort.Slice(got, func(i int, j int) bool { return got[i].AssignmentPath < got[j].AssignmentPath })

	for i, v := range want {
		if !cmp.Equal(v, got[i]) {
			t.Errorf("want = %v, got = %v", v, got[i])
		}
	}

	DeleteAssignment(a1.AssignmentPath)
	DeleteAssignment(a2.AssignmentPath)
	DeleteAssignment(a3.AssignmentPath)
	DeleteTeam(team.Name)
}
