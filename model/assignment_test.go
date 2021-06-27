package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	CreateTables()
}

var starter = &StarterCode{
	Url:             "git@gitlab.lrz.de:vss/startercode/startercodeB1.git",
	FromBranch:      "ws20.1",
	ProtectToBranch: true,
}

var clone = &Clone{
	LocalPath: "/Users/obraun/lectures/vss/labs/gitlab/semester/ob-20ws/blatt1",
	Branch:    "develop",
}

var assignment = &Assignment{
	AssignmentPath:    "blatt1",
	SemesterPath:      "semester/ob-20ws",
	Per:               "group",
	Description:       "Blatt 1, Verteilte Softwaresysteme, WS 20/21",
	ContainerRegistry: true,
	LocalPath:         clone.LocalPath,
	StarterUrl:        starter.Url,
}

func TestNewAssignmentSuccess(t *testing.T) {
	want, _ := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, clone.Branch, assignment.StarterUrl, starter.FromBranch,
		starter.ProtectToBranch)

	got := GetAssignment(want.AssignmentPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'\n", want, got)
	}
}

func TestNewAssignmentFailMissingAssignmentPath(t *testing.T) {
	wantStr := "Enter valid assignment path."

	got, gotString := NewAssignment("", assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, clone.Branch, assignment.StarterUrl, starter.FromBranch,
		starter.ProtectToBranch)

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
		assignment.LocalPath, clone.Branch, assignment.StarterUrl, starter.FromBranch,
		starter.ProtectToBranch)

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
		assignment.LocalPath, clone.Branch, assignment.StarterUrl, starter.FromBranch,
		starter.ProtectToBranch)

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
		assignment.LocalPath, clone.Branch, assignment.StarterUrl, starter.FromBranch,
		starter.ProtectToBranch)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestNewAssignmentFailMissingLocalPath(t *testing.T) {
	wantStr := "Enter valid local path."

	got, gotString := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		"", clone.Branch, assignment.StarterUrl, starter.FromBranch,
		starter.ProtectToBranch)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestNewAssignmentFailMissingBranch(t *testing.T) {
	wantStr := "Enter valid branch."

	got, gotString := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, "", assignment.StarterUrl, starter.FromBranch,
		starter.ProtectToBranch)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestNewAssignmentFailMissingStarterUrl(t *testing.T) {
	wantStr := "Enter valid starter url."

	got, gotString := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, clone.Branch, "", starter.FromBranch,
		starter.ProtectToBranch)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestNewAssignmentFailMissingFromBranch(t *testing.T) {
	wantStr := "Enter valid from branch."

	got, gotString := NewAssignment(assignment.AssignmentPath, assignment.SemesterPath,
		assignment.Per, assignment.Description, assignment.ContainerRegistry,
		assignment.LocalPath, clone.Branch, assignment.StarterUrl, "",
		starter.ProtectToBranch)

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
		StarterUrl:        starter.Url,
	}

	NewAssignment(as.AssignmentPath, as.SemesterPath,
		as.Per, as.Description,
		as.ContainerRegistry, as.LocalPath,
		clone.Branch, as.StarterUrl,
		starter.FromBranch, starter.ProtectToBranch)

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
		StarterUrl:        starter.Url,
	}

	NewAssignment(as.AssignmentPath, as.SemesterPath,
		as.Per, as.Description,
		as.ContainerRegistry, as.LocalPath,
		clone.Branch, as.StarterUrl,
		starter.FromBranch, starter.ProtectToBranch)

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

func TestGetStarterCode(t *testing.T) {
	DeleteStarterCode(starter.Url)
	DeleteClone(clone.LocalPath)

	as := &Assignment{
		AssignmentPath:    "blattX",
		SemesterPath:      "semester/ob-2Xws",
		Per:               "group",
		Description:       "Blatt X, Verteilte Softwaresysteme, WS 2X/2X",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.Url,
	}

	NewAssignment(as.AssignmentPath, as.SemesterPath,
		as.Per, as.Description,
		as.ContainerRegistry, as.LocalPath,
		clone.Branch, as.StarterUrl,
		starter.FromBranch, starter.ProtectToBranch)

	want := starter
	got := GetStarterCode(as.StarterUrl)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteAssignment(as.AssignmentPath)
}

func TestUpdateStarterCode(t *testing.T) {
	want := starter
	want.FromBranch = "github-dummy.com/starter"
	want.ProtectToBranch = false

	want.UpdateStarterCode()
	got := GetStarterCode(want.Url)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteStarterCode(want.Url)
}

func TestDeleteStarterCode(t *testing.T) {
	as := &Assignment{
		AssignmentPath:    "blattY",
		SemesterPath:      "semester/ob-2Yws",
		Per:               "group",
		Description:       "Blatt Y, Verteilte Softwaresysteme, WS 2Y/2Y",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.Url,
	}

	NewAssignment(as.AssignmentPath, as.SemesterPath,
		as.Per, as.Description,
		as.ContainerRegistry, as.LocalPath,
		clone.Branch, as.StarterUrl,
		starter.FromBranch, starter.ProtectToBranch)

	DeleteStarterCode(as.StarterUrl)

	want := &StarterCode{}
	got := GetStarterCode(as.StarterUrl)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteAssignment(as.AssignmentPath)
}

func TestGetClone(t *testing.T) {
	as := &Assignment{
		AssignmentPath:    "blattX",
		SemesterPath:      "semester/ob-2Xws",
		Per:               "group",
		Description:       "Blatt X, Verteilte Softwaresysteme, WS 2X/2X",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.Url,
	}

	NewAssignment(as.AssignmentPath, as.SemesterPath,
		as.Per, as.Description,
		as.ContainerRegistry, as.LocalPath,
		clone.Branch, as.StarterUrl,
		starter.FromBranch, starter.ProtectToBranch)

	want := clone
	got := GetClone(as.LocalPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteAssignment(as.AssignmentPath)
}

func TestUpdateClone(t *testing.T) {
	want := clone
	want.Branch = "master"

	want.UpdateClone()
	got := GetClone(want.LocalPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteClone(want.LocalPath)
}

func TestDeleteClone(t *testing.T) {
	as := &Assignment{
		AssignmentPath:    "blattY",
		SemesterPath:      "semester/ob-2Yws",
		Per:               "group",
		Description:       "Blatt Y, Verteilte Softwaresysteme, WS 2Y/2Y",
		ContainerRegistry: true,
		LocalPath:         clone.LocalPath,
		StarterUrl:        starter.Url,
	}

	NewAssignment(as.AssignmentPath, as.SemesterPath,
		as.Per, as.Description,
		as.ContainerRegistry, as.LocalPath,
		clone.Branch, as.StarterUrl,
		starter.FromBranch, starter.ProtectToBranch)

	DeleteClone(as.LocalPath)

	want := &Clone{}
	got := GetClone(as.LocalPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteClone(as.LocalPath)
	DeleteAssignment(as.AssignmentPath)
}
