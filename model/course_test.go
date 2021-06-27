package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	CreateTables()
}

func TestNewCourseSuccess(t *testing.T) {
	want := &Course{
		Path: "TestCourse",
	}

	NewCourse(want.Path)

	got := GetCourse(want.Path)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}
}

func TestNewCourseFail(t *testing.T) {
	got := NewCourse("")

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}
}
