package view

import (
	"fmt"
	"net/url"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
	glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
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
	body := widget.NewVScrollContainer(group)

	left := makeButtonForTeamOverview(tc, as.Assignment)
	right := glabsutil.MakeCloseButton(tc)
	buttons := glabsutil.MakeButtonGroup(left, right)

	as.Container = glabsutil.MakeScrollableView(body, buttons)

	return as
}

func makeButtonForTeamOverview(tc *widget.TabContainer, as *glabsmodel.Assignment) *widget.Button {
	overviewButton := widget.NewButton("Teamübersicht", func() {
		teamOverview := NewTeamOverview(as, tc)
		item := widget.NewTabItem(fmt.Sprintf("Teams %s", as.Name), teamOverview.Container)
		tc.Append(item)
	})

	return overviewButton
}
