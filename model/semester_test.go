package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func init() {
	CreateTables()
}

func TestNewSemesterSuccess(t *testing.T) {
	want := &Semester{
		Path:       "test_semester/hm.edu",
		CoursePath: "test_course",
	}

	NewSemester(want.CoursePath, want.Path)

	got := GetSemester(want.Path)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteSemester(want.Path)
}

func TestNewSemesterFailNoPath(t *testing.T) {
	got := NewSemester("test_course", "")

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}
}

func TestNewSemesterFailNoCourse(t *testing.T) {
	got := NewSemester("", "test_path")

	if got != nil {
		t.Errorf("want = <nil>, got = '%v'", got)
	}
}

func TestDeleteSemester(t *testing.T) {
	s := &Semester{
		Path:       "test_semester/hm.edu",
		CoursePath: "test_course",
	}
	NewSemester(s.CoursePath, s.Path)
	DeleteSemester(s.Path)

	want := &Semester{}

	got := GetSemester(s.Path)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}
}

func TestUpdateSemester(t *testing.T) {
	s := &Semester{
		Path:       "test_semester/hm.edu",
		CoursePath: "test_course",
	}
	NewSemester(s.CoursePath, s.Path)
	newCourse := "test_course_2"
	s.UpdateSemester(newCourse)
	s.CoursePath = newCourse

	want := s
	
	got := GetSemester(s.Path)

	if !cmp.Equal(want, got) {
		t.Errorf("want = '%v', got = '%v'", want, got)
	}

	DeleteSemester(s.Path)
}
