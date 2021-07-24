package view

import (
	"fmt"

	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type SemesterDetailView struct {
	TabContainer *widget.TabContainer
	Container    *fyne.Container
	Semester     *glabsmodel.Semester
}

func NewSemesterDetailView(semester *glabsmodel.Semester, tc *widget.TabContainer) *SemesterDetailView {

	return nil
}

func makeButtonForAssignmentOverview(tc *widget.TabContainer, s *glabsmodel.Semester) *widget.Button {
	overviewButton := widget.NewButton("Assignment-Übersicht", func() {
		assignmentOverview := NewAssignmentOverview(s, tc)
		item := widget.NewTabItem(fmt.Sprintf("Assignments %s", s.Path), assignmentOverview.Container)
		tc.Append(item)
	})

	return overviewButton
}
