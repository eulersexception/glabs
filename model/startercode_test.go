package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var starter = &StarterCode{
	StarterUrl:      "git@gitlab.lrz.de:vss/startercode/startercodeB1.git",
	FromBranch:      "ws20.1",
	ProtectToBranch: true,
}

func init() {
	CreateTables()
}

func TestNewStarterUrlFailMissingStarterUrl(t *testing.T) {
	wantStr := "Enter valid starter url."

	got, gotString := NewStarterCode("", starter.FromBranch, starter.ProtectToBranch)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestNewStarterUrlFailMissingFromBranch(t *testing.T) {
	wantStr := "Enter valid from branch."

	got, gotString := NewStarterCode(starter.StarterUrl, "", starter.ProtectToBranch)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestGetStarterCode(t *testing.T) {
	DeleteStarterCode(starter.StarterUrl)

	NewStarterCode(starter.StarterUrl, starter.FromBranch, starter.ProtectToBranch)

	want := starter
	got := GetStarterCode(starter.StarterUrl)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteStarterCode(starter.StarterUrl)
}

func TestDeleteStarterCode(t *testing.T) {
	NewStarterCode(starter.StarterUrl, starter.FromBranch, starter.ProtectToBranch)
	DeleteStarterCode(starter.StarterUrl)

	want := &StarterCode{}
	got := GetStarterCode(starter.StarterUrl)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}
}

func TestUpdateStarterCode(t *testing.T) {
	want, _ := NewStarterCode(starter.StarterUrl, starter.FromBranch, starter.ProtectToBranch)
	want.FromBranch = "github-dummy.com/starter"
	want.ProtectToBranch = false

	want.UpdateStarterCode()
	got := GetStarterCode(want.StarterUrl)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteStarterCode(want.StarterUrl)
}

func TestGetAllAssignments(t *testing.T) {
	s, _ := NewStarterCode("testurl.com", "develop", true)

	a1, _ := NewAssignment("PathA", "SemesterX", "group", "BlattA", true, "TestCloneA", s.StarterUrl)
	a2, _ := NewAssignment("PathB", "SemesterY", "single", "BlattB", false, "TestCloneB", s.StarterUrl)
	a3, _ := NewAssignment("PathC", "SemesterZ", "pair", "BlattC", true, "TestCloneC", s.StarterUrl)

	want := make([]Assignment, 0)
	want = append(want, *a1, *a2, *a3)
	got := GetAllAssignmentsForStarterCode(s.StarterUrl)

	if cmp.Equal(want, got) {
		t.Errorf("want = %v, got %v", want, got)
	}

	DeleteAssignment(a1.AssignmentPath)
	DeleteAssignment(a2.AssignmentPath)
	DeleteAssignment(a3.AssignmentPath)
	DeleteStarterCode(s.StarterUrl)
}

func TestUpdateStarterUrl(t *testing.T) {
	s, _ := NewStarterCode("testurl.com", "develop", true)

	a1, _ := NewAssignment("PathA", "SemesterX", "group", "BlattA", true, "TestCloneA", s.StarterUrl)
	a2, _ := NewAssignment("PathB", "SemesterY", "single", "BlattB", false, "TestCloneB", s.StarterUrl)
	a3, _ := NewAssignment("PathC", "SemesterZ", "pair", "BlattC", true, "TestCloneC", s.StarterUrl)

	UpdateStarterUrl(s.StarterUrl, "newtesturl.com")
	s.StarterUrl = "newtesturl.com"

	want := make([]Assignment, 0)
	want = append(want, *a1, *a2, *a3)
	got := GetAllAssignmentsForStarterCode(s.StarterUrl)

	if cmp.Equal(want, got) {
		t.Errorf("want = %v, got %v", want, got)
	}

	DeleteAssignment(a1.AssignmentPath)
	DeleteAssignment(a2.AssignmentPath)
	DeleteAssignment(a3.AssignmentPath)
	DeleteStarterCode(s.StarterUrl)
}
