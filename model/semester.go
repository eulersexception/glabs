package model

import "fmt"

type Semester struct {
	Name  string
	Repos []Repository
}

func NewSemester(name string) *Semester {
	if name == "" {
		fmt.Println("Provide a valid naming for semester")
		return nil
	}

	var repos []Repository

	semester := &Semester{
		Name:  name,
		Repos: repos,
	}

	return semester
}

func (s *Semester) AddRepoToSemester(repo *Repository) *Semester {
	if repo == nil {
		fmt.Println("No valid data for repository")
		return s
	}

	s.Repos = append(s.Repos, *repo)

	return s
}

func (s *Semester) DeleteRepoFromSemester(repo *Repository) *Semester {
	if repo == nil {
		fmt.Println("No valid data for repository")
	}

	index := -1

	for i, v := range s.Repos {
		if v.Name == repo.Name {
			index = i
		}
	}

	if index == -1 {
		return s
	}

	s.Repos[index] = s.Repos[len(s.Repos)-1]
	s.Repos = s.Repos[:len(s.Repos)-1]

	return s
}
