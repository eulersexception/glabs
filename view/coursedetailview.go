package view

import (
	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type CourseDetailView struct {
	TabContainer *widget.TabContainer
	Container    *fyne.Container
	Course       *glabsmodel.Course
}

func NewCourseDetailView(course *glabsmodel.Course, tc *widget.TabContainer) *CourseDetailView {
	return nil
}

func makeButtonForSemesterOverview(tc *widget.TabContainer, c *glabsmodel.Course) *widget.Button {
	return nil
}
