package model

import "fmt"

type Repository struct {
	Name  string
	Url   string
	Teams []Team
}

func NewRepository(name string, url string) *Repository {
	if name == "" {
		fmt.Println("Please enter a valid course name.")
		return nil
	}

	var teams []Team

	repo := &Repository{
		Name:  name,
		Url:   url,
		Teams: teams,
	}

	return repo
}

func (r *Repository) AddTeamToRepository(t *Team) *Repository {
	if t == nil {
		fmt.Println("Nod valid data for team.")
		return r
	}

	r.Teams = append(r.Teams, *t)

	return r
}

func (r *Repository) DeleteTeamFromRepository(t Team) *Repository {
	index := -1

	for i, v := range r.Teams {
		if v.Name == t.Name {
			index = i
		}
	}

	if index == -1 {
		return r
	}

	r.Teams[index] = r.Teams[len(r.Teams)-1]
	r.Teams = r.Teams[:len(r.Teams)-1]

	return r
}
