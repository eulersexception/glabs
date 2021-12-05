package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var clone = &Clone{
	LocalPath: "/Users/obraun/lectures/vss/labs/gitlab/semester/ob-20ws/blatt1",
	Branch:    "develop",
}

func init() {
	CreateTables()
}

func TestNewCloneFailMissingLocalPath(t *testing.T) {
	wantStr := "Enter valid local path."

	got, gotString := NewClone("", clone.Branch)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestNewCloneFailMissingBranch(t *testing.T) {
	wantStr := "Enter valid branch."

	got, gotString := NewClone(clone.LocalPath, "")

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}

	if wantStr != gotString {
		t.Errorf("want = <nil>, got = <nil>\nwant message = '%s'\ngot message = '%s'", wantStr, gotString)
	}
}

func TestGetClone(t *testing.T) {
	want, _ := NewClone(clone.LocalPath, clone.Branch)

	got := GetClone(clone.LocalPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteClone(clone.LocalPath)
}

func TestUpdateClone(t *testing.T) {
	want, _ := NewClone(clone.LocalPath, clone.Branch)
	want.Branch = "master"
	want.UpdateClone()

	got := GetClone(want.LocalPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteClone(want.LocalPath)
}

func TestDeleteClone(t *testing.T) {
	NewClone(clone.LocalPath, clone.Branch)
	DeleteClone(clone.LocalPath)

	want := &Clone{}

	got := GetClone(clone.LocalPath)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteClone(clone.LocalPath)
}

func TestGetAllAssignmentsForClone(t *testing.T) {
	c, _ := NewClone("/Users/obraun/lectures/statistik/labs/gitlab/semester/ob-20ws/blatt1", "develop")
	a1, _ := NewAssignment("PathA", "SemesterX", "group", "BlattA", true, c.LocalPath, "blabla.com")
	a2, _ := NewAssignment("PathB", "SemesterY", "single", "BlattB", false, c.LocalPath, "blabla.com")
	a3, _ := NewAssignment("PathC", "SemesterZ", "pair", "BlattC", true, c.LocalPath, "blabla.com")

	want := make([]Assignment, 0)
	want = append(want, *a1, *a2, *a3)

	got := GetAllAssignmentsForClone(c.LocalPath)

	if cmp.Equal(want, got) {
		t.Errorf("want = %v, got %v", want, got)
	}

	DeleteAssignment(a1.AssignmentPath)
	DeleteAssignment(a2.AssignmentPath)
	DeleteAssignment(a3.AssignmentPath)
	DeleteClone(c.LocalPath)
}

func TestUpdateClonePath(t *testing.T) {
	c, _ := NewClone("/Users/obraun/lectures/statistik/labs/gitlab/semester/ob-20ws/blatt1", "develop")
	a1, _ := NewAssignment("PathA", "SemesterX", "group", "BlattA", true, c.LocalPath, "blabla.com")
	a2, _ := NewAssignment("PathB", "SemesterY", "single", "BlattB", false, c.LocalPath, "blabla.com")
	a3, _ := NewAssignment("PathC", "SemesterZ", "pair", "BlattC", true, c.LocalPath, "blabla.com")
	UpdateClonePath(c.LocalPath, "new/clone/path")
	c.LocalPath = "new/clone/path"

	want := make([]Assignment, 0)
	want = append(want, *a1, *a2, *a3)

	got := GetAllAssignmentsForClone(c.LocalPath)

	if cmp.Equal(want, got) {
		t.Errorf("want = %v, got %v", want, got)
	}

	DeleteAssignment(a1.AssignmentPath)
	DeleteAssignment(a2.AssignmentPath)
	DeleteAssignment(a3.AssignmentPath)
	DeleteClone(c.LocalPath)
}
