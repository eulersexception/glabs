package model

import (
	"fmt"
)

type StarterCode struct {
	Url             string
	FromBranch      string
	ProtectToBranch bool
}

func (s StarterCode) toString() string {
	return fmt.Sprintf("\tStarterCode:\n\t\tUrl:\t%s\n\t\tFromBranch:\t%s\n\t\tProtectToBranch:\t%v", s.Url, s.FromBranch, s.ProtectToBranch)
}

type Clone struct {
	LocalPath string
	Branch    string
}

func (c Clone) toString() string {
	return fmt.Sprintf("\tClone:\n\t\tLocalPath:\t%s\n\t\tBranch:\t%s", c.LocalPath, c.Branch)
}

type Assignment struct {
	AssignmentID      *int64       `ql:"index xID"`
	Name              string       `ql:"uindex xName, name AssignmentName"`
	Semester          Semester     `ql:"SemesterName"`
	LocalClone        *Clone       `ql:"-"`
	Starter           *StarterCode `ql:"-"`
	ContainerRegistry bool         `ql:"-"`
}

func NewAssignment(name string, sem *Semester, clone *Clone, starter *StarterCode) *Assignment {
	if name == "" {
		fmt.Println("Please enter a valid course name.")
		return nil
	}

	assignment := &Assignment{
		Semester:   *sem,
		Name:       name,
		Starter:    starter,
		LocalClone: clone,
	}

	return assignment
}

func (a *Assignment) AddTeamToAssignment(t *Team) *Assignment {
	return nil
}

func (a *Assignment) DeleteTeamFromAssignment(t Team) *Assignment {
	return nil
}

func (a Assignment) SetAssignment() {
}

func GetAssignment(name string) *Assignment {
	return nil

}

func DeleteAssignment(name string) error {

	return nil
}

func (a Assignment) PrintData() {

}
