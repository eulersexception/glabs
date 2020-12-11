package model

import "fmt"

type Team struct {
	Name     string
	Url      string
	Students []*Student
}

func NewTeam(name string, url string) *Team {
	if name == "" {
		fmt.Println("Please enter a valid team name.")
		return nil
	}

	var students []*Student

	team := &Team{
		Name:     name,
		Url:      url,
		Students: students,
	}

	return team
}

func (t *Team) AddStudentToTeam(s *Student) *Team {
	t.Students = append(t.Students, s)

	return t
}

func (t *Team) RemoveStudentFromTeam(s Student) *Team {
	index := -1

	for i, v := range t.Students {
		if s.Id == v.Id {
			index = i
		}
	}

	if index == -1 {
		return t
	}

	t.Students[index] = t.Students[len(t.Students)-1]
	t.Students = t.Students[:len(t.Students)-1]

	return t
}
