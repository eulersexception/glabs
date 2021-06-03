package model

import (
	"fmt"
)

type Course struct {
	CourseID    int64  `ql:"index xID"`
	Name        string `ql:"uindex xName, name CourseName"`
	Description string
}

func NewCourse(name string, description string) *Course {

	if name == "" {
		fmt.Println("Please enter a valid course name.")
		return nil
	}

	c := &Course{
		Name: name,
	}

	if description != "" {
		c.Description = description
	}

	return c
}

func (c *Course) AddSemesterToCourse(s *Semester) *Course {
	return nil
}

func (c *Course) DeleteSemesterFromCourse(s *Semester) *Course {
	return nil
}
