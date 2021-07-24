package view

import (
	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne/widget"
)

type TeamOverview struct {
	TabContainer *widget.TabContainer
	Container    *widget.ScrollContainer
	Teams        []*glabsmodel.Team
	Assignment   *glabsmodel.Assignment
}

func NewTeamOverview(assignment *glabsmodel.Assignment, tc *widget.TabContainer) *TeamOverview {

	return nil
}

func addStudents(n int, t *glabsmodel.Team) {

}
