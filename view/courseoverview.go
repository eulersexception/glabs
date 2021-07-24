package view

import (
	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne/widget"
)

type CourseOverview struct {
	TabContainer *widget.TabContainer
	Container    *widget.ScrollContainer
	Courses      []*glabsmodel.Course
}

func NewCourseOverview(courses []*glabsmodel.Course, tc *widget.TabContainer) *CourseOverview {

	return nil
}

func addSemesters(n int, course *glabsmodel.Course) {

}
