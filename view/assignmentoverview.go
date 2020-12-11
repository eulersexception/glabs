package view

import (
	"fmt"
	"net/url"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
	glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type AssignmentOverview struct {
	TabContainer *widget.TabContainer
	Container    *widget.ScrollContainer
	Assignments  []*glabsmodel.Assignment
	Semester     *glabsmodel.Semester
}

func NewAssignmentOverview(semester *glabsmodel.Semester, tc *widget.TabContainer) *AssignmentOverview {
	a := &AssignmentOverview{
		Assignments:  semester.Assignments,
		TabContainer: tc,
		Semester:     semester,
	}

	group := widget.NewGroup(fmt.Sprintf("Assignmentübersicht %s", semester.Name))

	if a.Assignments != nil {
		for _, v := range a.Assignments {
			currentAssignment := &glabsmodel.Assignment{
				Name:  v.Name,
				Url:   v.Url,
				Teams: v.Teams,
			}

			addTeams(5, currentAssignment, semester)

			button := widget.NewButton("Details", func() {
				tc.Append(widget.NewTabItem(currentAssignment.Name, NewAssignmentDetailView(currentAssignment, tc).Container))
				tc.Remove(tc.CurrentTab())
			})
			label := widget.NewLabel(v.Name)
			url := widget.NewHyperlink("Repo", &url.URL{Scheme: "https", Host: v.Url})
			url.Wrapping = fyne.TextTruncate
			line := widget.NewHBox(label, url, layout.NewSpacer(), button)
			group.Append(line)
		}

		a.Container = widget.NewVScrollContainer(group)

	} else {
		text := widget.NewLabel("Aktuell keine Assignments vorhanden")
		left := MakeButtonForCourseCreation(tc)
		right := glabsutil.MakeCloseButton(tc)
		buttons := glabsutil.MakeButtonGroup(left, right)

		a.Container = widget.NewVScrollContainer(fyne.NewContainerWithLayout(layout.NewHBoxLayout(), text, layout.NewSpacer(), buttons))
	}

	return a
}

func addTeams(n int, as *glabsmodel.Assignment, s *glabsmodel.Semester) {
	temp := fmt.Sprintf("%s %s %s", s.Course.Name, s.Name, as.Name)

	for i := 0; i < n; i++ {

		name := fmt.Sprintf("Team %d", i)
		url := fmt.Sprintf("%s-%s", temp, name)
		t := glabsmodel.NewTeam(name, url)
		as.AddTeamToAssignment(t)
	}
}
