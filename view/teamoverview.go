package view

import (
	"fmt"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
	glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type TeamOverview struct {
	TabContainer *widget.TabContainer
	Container    *widget.ScrollContainer
	Teams        []*glabsmodel.Team
	Assignment   *glabsmodel.Assignment
}

func NewTeamOverview(assignment *glabsmodel.Assignment, tc *widget.TabContainer) *TeamOverview {
	a := &TeamOverview{
		Teams:        assignment.Teams,
		TabContainer: tc,
		Assignment:   assignment,
	}

	group := widget.NewGroup(fmt.Sprintf("Teamübersicht %s", assignment.Name))

	if a.Teams != nil {
		for _, v := range a.Teams {
			currentTeam := &glabsmodel.Team{
				Assignment: assignment,
				Name:       v.Name,
				Students:   v.Students,
			}

			addStudents(5, currentTeam)

			button := widget.NewButton("Details", func() {
				tc.Append(widget.NewTabItem(currentTeam.Name, NewTeamDetailView(currentTeam, tc).Container))
				tc.Remove(tc.CurrentTab())
			})
			label := widget.NewLabel(v.Name)
			line := widget.NewHBox(label, layout.NewSpacer(), button)
			group.Append(line)
		}

		a.Container = widget.NewVScrollContainer(group)

	} else {
		text := widget.NewLabel("Aktuell keine Teams vorhanden")
		left := MakeButtonForCourseCreation(tc)
		right := glabsutil.MakeCloseButton(tc)
		buttons := glabsutil.MakeButtonGroup(left, right)

		a.Container = widget.NewVScrollContainer(fyne.NewContainerWithLayout(layout.NewHBoxLayout(), text, layout.NewSpacer(), buttons))
	}

	return a
}

func addStudents(n int, t *glabsmodel.Team) {
	for i := 0; i < n; i++ {
		nick := fmt.Sprintf("max_payne_%d", i)
		mail := fmt.Sprintf("einsteinNo_%d_@fantasiaschool.edu", i)
		s, _ := glabsmodel.NewStudent(t, "Mustermann", fmt.Sprintf("Max der %d.", i), nick, mail, uint32(i))
		//t.AddStudentToTeam(s)
		s.PrintData()
	}
}
