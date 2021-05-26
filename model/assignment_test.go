package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewAssignmentSuccess(t *testing.T) {
	want := &Assignment{Name: "TestAssignment1", Semester: "SS2018"}

	NewAssignment(want.Name, want.Semester, nil, nil)

	got := GetAssignment(want.Name)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'\n", want, got)
	}
}

func TestNewAssignmentFailMissingName(t *testing.T) {
	got := NewAssignment("", "SS2018", nil, nil)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}
}

func TestNewAssignmentFailMissingSemester(t *testing.T) {
	got := NewAssignment("TestAssignment", "", nil, nil)

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}
}

func TestDeleteAssignment(t *testing.T) {
	a := &Assignment{Name: "TestAssignment2", Semester: "WS2018"}

	NewAssignment(a.Name, a.Semester, nil, nil)
	DeleteAssignment(a.Name)

	want := &Assignment{}
	got := GetAssignment(a.Name)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'\n", want, got)
	}
}

func TestUpdateAssignment(t *testing.T) {
	want := &Assignment{Name: "TestAssignment1", Semester: "SS2019"}
	want.UpdateAssignment()
	got := GetAssignment(want.Name)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'\n", want, got)
	}

	want.Semester = "SS2018"
	want.UpdateAssignment()
}
