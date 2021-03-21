package model

import (
	"fmt"
)

type Course struct {
	Name        string
	Description string
	Semesters   []*Semester
}

func NewCourse(name string, description string) *Course {

	if name == "" {
		fmt.Println("Please enter a valid course name.")
		return nil
	}

	var semesters []*Semester

	c := &Course{
		Name:      name,
		Semesters: semesters,
	}

	if description != "" {
		c.Description = description
	}

	return c
}

func (c *Course) AddSemesterToCourse(s *Semester) *Course {
	if s == nil {
		fmt.Println("No valid argument for student")
		return c
	}

	c.Semesters = append(c.Semesters, s)

	return c
}

func (c *Course) DeleteSemesterFromCourse(s *Semester) *Course {
	if s == nil {
		fmt.Println("No valid argument for student")
	}

	index := -1

	for i, v := range c.Semesters {
		if v.Name == s.Name {
			index = i
		}
	}

	if index == -1 {
		return c
	}

	c.Semesters[index] = c.Semesters[len(c.Semesters)-1]
	c.Semesters = c.Semesters[:len(c.Semesters)-1]

	return c
}
