package view

import (
	"net/url"

	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type CourseDetailView struct {
	Container *widget.ScrollContainer
	Course    *glabsmodel.Course
}

func NewCourseDetailView(course *glabsmodel.Course) *CourseDetailView {
	c := &CourseDetailView{
		Course: course,
	}

	header := widget.NewLabelWithStyle(c.Course.Name, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	description := widget.NewTextGrid()
	description.SetText(c.Course.Description)
	url := widget.NewHyperlink("Repo", &url.URL{Scheme: "https", Host: c.Course.Url})
	layout := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), header, description, url)
	c.Container = widget.NewVScrollContainer(layout)

	return c
}
