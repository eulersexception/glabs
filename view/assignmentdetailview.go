package view

import (
	glabsmodel "github.com/eulersexception/glabs-ui/model"
	//glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type AssignmentDetailView struct {
	TabContainer *widget.TabContainer
	Container    *fyne.Container
	Assignment   *glabsmodel.Assignment
}

func NewAssignmentDetailView(assignment *glabsmodel.Assignment, tc *widget.TabContainer) *AssignmentDetailView {
	return nil
}

func makeButtonForTeamOverview(tc *widget.TabContainer, as *glabsmodel.Assignment) *widget.Button {
	return nil
}
