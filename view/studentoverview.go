package view

import (
	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne/widget"
)

type StudentOverview struct {
	TabContainer *widget.TabContainer
	Container    *widget.ScrollContainer
	Students     []*glabsmodel.Student
	Team         *glabsmodel.Team
}

func NewStudentOverview(team *glabsmodel.Team, tc *widget.TabContainer) *StudentOverview {

	return nil
}
