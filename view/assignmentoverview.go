package view

import (
	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne/widget"
)

type AssignmentOverview struct {
	TabContainer *widget.TabContainer
	Container    *widget.ScrollContainer
	Assignments  []*glabsmodel.Assignment
	Semester     *glabsmodel.Semester
}

func NewAssignmentOverview(semester *glabsmodel.Semester, tc *widget.TabContainer) *AssignmentOverview {

	return nil
}

func addTeams(n int, as *glabsmodel.Assignment, s *glabsmodel.Semester) {

}
