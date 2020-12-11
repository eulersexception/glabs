package model

import "fmt"

type Assignment struct {
	Name  string
	Url   string
	Teams []*Team
}

func NewAssignment(name string, url string) *Assignment {
	if name == "" {
		fmt.Println("Please enter a valid course name.")
		return nil
	}

	var teams []*Team

	assignment := &Assignment{
		Name:  name,
		Url:   url,
		Teams: teams,
	}

	return assignment
}

func (a *Assignment) AddTeamToAssignment(t *Team) *Assignment {
	if t == nil {
		fmt.Println("No valid data for team.")
		return a
	}

	a.Teams = append(a.Teams, t)

	return a
}

func (a *Assignment) DeleteTeamFromAssignment(t Team) *Assignment {
	index := -1

	for i, v := range a.Teams {
		if v.Name == t.Name {
			index = i
		}
	}

	if index == -1 {
		return a
	}

	a.Teams[index] = a.Teams[len(a.Teams)-1]
	a.Teams = a.Teams[:len(a.Teams)-1]

	return a
}
