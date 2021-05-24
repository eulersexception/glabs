package model

import "fmt"

type Semester struct {
	SemesterID *int64  `ql:"index xID"`
	Name       string  `ql:"uindex xName, name SemesterName"`
	Course     *Course `ql:"CourseName"`
}

func NewSemester(course *Course, name string, url string) *Semester {
	if name == "" {
		fmt.Println("Provide a valid naming for semester")
		return nil
	}

	semester := &Semester{
		Course: course,
		Name:   name,
	}

	return semester
}

func (s *Semester) AddAssignmentToSemester(a *Assignment) *Semester {

	return nil
}

func (s *Semester) DeleteAssignmentFromSemester(a *Assignment) *Semester {

	return nil
}
