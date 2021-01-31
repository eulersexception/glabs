package view

import (
	"fmt"
	"net/url"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
	glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type SemesterDetailView struct {
	TabContainer *widget.TabContainer
	Container    *fyne.Container
	Semester     *glabsmodel.Semester
}

func NewSemesterDetailView(semester *glabsmodel.Semester, tc *widget.TabContainer) *SemesterDetailView {
	s := &SemesterDetailView{
		TabContainer: tc,
		Semester:     semester,
	}

	group := widget.NewGroup(fmt.Sprintf("%s", s.Semester.Name))
	name := widget.NewLabel(s.Semester.Name)
	url := widget.NewHyperlink("Repo", &url.URL{Scheme: "https", Host: s.Semester.Url})
	group.Append(name)
	group.Append(url)
	body := widget.NewVScrollContainer(group)

	left := makeButtonForAssignmentOverview(tc, s.Semester)
	right := glabsutil.MakeCloseButton(s.TabContainer)
	buttons := glabsutil.MakeButtonGroup(left, right)

	s.Container = glabsutil.MakeScrollableView(body, buttons)

	return s
}

func makeButtonForAssignmentOverview(tc *widget.TabContainer, s *glabsmodel.Semester) *widget.Button {
	overviewButton := widget.NewButton("Assignment-Übersicht", func() {
		assignmentOverview := NewAssignmentOverview(s, tc)
		item := widget.NewTabItem(fmt.Sprintf("Assignments %s", s.Name), assignmentOverview.Container)
		tc.Append(item)
	})

	return overviewButton
}
