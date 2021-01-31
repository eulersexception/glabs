package view

import (
	"fmt"

	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type SemesterOverview struct {
	TabContainer *widget.TabContainer
	Container    *widget.ScrollContainer
	Semesters    []*glabsmodel.Semester
	Course       *glabsmodel.Course
}

func NewSemesterOverview(tc *widget.TabContainer, course *glabsmodel.Course) *SemesterOverview {
	s := &SemesterOverview{
		TabContainer: tc,
		Course:       course,
		Semesters:    course.Semesters,
	}

	group := widget.NewGroup(fmt.Sprintf("Semesterübersicht %s", s.Course.Name))

	if s.Semesters != nil {

		for _, v := range s.Semesters {
			currentSemester := &glabsmodel.Semester{
				Name:        v.Name,
				Assignments: v.Assignments,
				Course:      course,
			}

			addAssignments(20, currentSemester)

			button := widget.NewButton("Details", func() {
				tc.Append(widget.NewTabItem(currentSemester.Name, NewSemesterDetailView(currentSemester, s.TabContainer).Container))
				tc.Remove(tc.CurrentTab())
			})
			label := widget.NewLabel(v.Name)
			line := widget.NewHBox(label, layout.NewSpacer(), layout.NewSpacer(), button)
			group.Append(line)
		}

		s.Container = widget.NewVScrollContainer(group)
	} else {
		group := widget.NewGroup(fmt.Sprintf("Semesterübersicht %s", s.Course.Name))
		text := widget.NewTextGrid()
		text.SetText("Hallo Semesteruebersicht")
		button := widget.NewButton("Home", func() {
			tc.SelectTabIndex(0)
		})
		group.Append(text)

		s.Container = widget.NewVScrollContainer(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), group, layout.NewSpacer(), button))
	}

	return s
}

func addAssignments(n int, s *glabsmodel.Semester) {

	for i := 0; i < n; i++ {
		as := &glabsmodel.Assignment{
			Name: fmt.Sprintf("Assignment %d %s", i, s.Name),
			Url:  "www.google.com",
		}
		s.AddAssignmentToSemester(as)
	}
}
