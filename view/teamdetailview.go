package view

import (
	"fmt"

	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type TeamDetailView struct {
	TabContainer *widget.TabContainer
	Container    *fyne.Container
	Team         *glabsmodel.Team
}

func NewTeamDetailView(team *glabsmodel.Team, tc *widget.TabContainer) *TeamDetailView {
	t := &TeamDetailView{
		TabContainer: tc,
		Team:         team,
	}

	//t.Container = glabsutil.MakeScrollableView(body, buttons)

	return t
}

func makeButtonForStudentsOverview(tc *widget.TabContainer, t *glabsmodel.Team) *widget.Button {
	overviewButton := widget.NewButton("Mitglieder", func() {
		studentOverview := NewStudentOverview(t, tc)
		item := widget.NewTabItem(fmt.Sprintf("Mitglieder von %s", t.Name), studentOverview.Container)
		tc.Append(item)
	})

	return overviewButton
}
