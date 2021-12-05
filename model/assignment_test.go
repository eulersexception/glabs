package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	CreateTables()
}

var localStarter = &StarterCode{
	StarterUrl:      "git@gitlab.lrz.de:vss/startercode/startercodeC1.git",
	FromBranch:      "ws20.1",
	ProtectToBranch: true,
}

var localClone = &Clone{
	LocalPath: "/Users/obraun/lectures/vss/labs/gitlab/semester/ob-20ws/blatt10",
	Branch:    "develop",
}

var assignment = &Assignment{
	AssignmentPath:    "blatt1",
	SemesterPath:      "semester/ob-20ws",
	Per:               "group",
	Description:       "Blatt 1, Verteilte Softwaresysteme, WS 20/21",
	ContainerRegistry: true,
	LocalPath:         localClone.LocalPath,
	StarterUrl:        localStarter.StarterUrl,
}

func TestNewAssignmentSuccess(t *testing.T) {
	want, _ := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, assignment.StarterUrl)

	got := GetAssignment(want.AssignmentPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'\n", want, got)
	}

	DeleteAssignment(want.AssignmentPath)
}

func TestNewAssignmentFailMissingAssignmentPath(t *testing.T) {
	wantStr := "Enter valid assignment path."

	got, gotString := NewAssignment("", assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, assignment.StarterUrl)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestNewAssignmentFailMissingSemesterPath(t *testing.T) {
	wantStr := "Enter valid semester path."

	got, gotString := NewAssignment(assignment.AssignmentPath, "",
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, assignment.StarterUrl)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestNewAssignmentFailMissingPer(t *testing.T) {
	wantStr := "Enter valid per."

	got, gotString := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		"", assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, assignment.StarterUrl)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestNewAssignmentFailMissingDescription(t *testing.T) {
	wantStr := "Enter valid description."

	got, gotString := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		assignment.Per, "", assignment.ContainerRegistry,
		assignment.LocalPath, assignment.StarterUrl)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestDeleteAssignment(t *testing.T) {
	as := &Assignment{
		AssignmentPath:    "blattX",
		SemesterPath:      "semester/ob-2Xws",
		Per:               "group",
		Description:       "Blatt X, Verteilte Softwaresysteme, WS 2X/2X",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.StarterUrl,
	}

	NewAssignment(as.AssignmentPath, as.SemesterPath,
		as.Per, as.Description,
		as.ContainerRegistry, as.LocalPath,
		as.StarterUrl)

	DeleteAssignment(as.AssignmentPath)

	want := &Assignment{}
	got := GetAssignment(as.AssignmentPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}
}

func TestUpdateAssignment(t *testing.T) {
	as := &Assignment{
		AssignmentPath:    "blattX",
		SemesterPath:      "semester/ob-2Xws",
		Per:               "group",
		Description:       "Blatt X, Verteilte Softwaresysteme, WS 2X/2X",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.StarterUrl,
	}

	NewAssignment(as.AssignmentPath, as.SemesterPath,
		as.Per, as.Description,
		as.ContainerRegistry, as.LocalPath,
		as.StarterUrl)

	as.SemesterPath = "semester/ob-22ws"
	as.Per = "single"
	as.Description = "Blatt X Updated, Verteilte Softwaresysteme, WS 2X/2X"
	as.ContainerRegistry = false

	as.UpdateAssignment()

	want := as
	got := GetAssignment(as.AssignmentPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteAssignment(as.AssignmentPath)
}
