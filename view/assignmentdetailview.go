package view

import (
	"net/url"

	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type AssignmentDetailView struct {
	TabContainer *widget.TabContainer
	Container    *fyne.Container
	Assignment   *glabsmodel.Assignment
}

func NewAssignmentDetailView(assignment *glabsmodel.Assignment, tc *widget.TabContainer) *AssignmentDetailView {
	as := &AssignmentDetailView{
		TabContainer: tc,
		Assignment:   assignment,
	}

	group := widget.NewGroup(as.Assignment.Name)
	desc := widget.NewLabel(as.Assignment.Name)
	url := widget.NewHyperlink("Repo", &url.URL{Scheme: "https", Host: as.Assignment.Url})
	group.Append(desc)
	group.Append(url)

	as.Container = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), group)

	return as
}
