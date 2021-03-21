package model

import "fmt"

type Semester struct {
	Name        string
	Course      *Course
	Assignments []*Assignment
}

func NewSemester(course *Course, name string, url string) *Semester {
	if name == "" {
		fmt.Println("Provide a valid naming for semester")
		return nil
	}

	var assignments []*Assignment

	semester := &Semester{
		Course:      course,
		Name:        name,
		Assignments: assignments,
	}

	return semester
}

func (s *Semester) AddAssignmentToSemester(a *Assignment) *Semester {
	if a == nil {
		fmt.Println("No valid data for repository")
		return s
	}

	s.Assignments = append(s.Assignments, a)

	return s
}

func (s *Semester) DeleteAssignmentFromSemester(a *Assignment) *Semester {
	if a == nil {
		fmt.Println("No valid data for repository")
	}

	index := -1

	for i, v := range s.Assignments {
		if v.Name == a.Name {
			index = i
		}
	}

	if index == -1 {
		return s
	}

	s.Assignments[index] = s.Assignments[len(s.Assignments)-1]
	s.Assignments = s.Assignments[:len(s.Assignments)-1]

	return s
}
