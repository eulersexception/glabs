package view

import (
	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne/widget"
)

type SemesterOverview struct {
	TabContainer *widget.TabContainer
	Container    *widget.ScrollContainer
	Semesters    []*glabsmodel.Semester
	Course       *glabsmodel.Course
}

func NewSemesterOverview(tc *widget.TabContainer, course *glabsmodel.Course) *SemesterOverview {

	return nil
}

func addAssignments(n int, s *glabsmodel.Semester) {

}
