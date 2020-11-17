package view

import (
	"net/url"

	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne"
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
	group := widget.NewGroup("")
	header := widget.NewLabel(c.Course.Name)
	description := widget.NewTextGrid()
	description.SetText(c.Course.Description)
	description.Resize(fyne.NewSize(300, 300))
	url := widget.NewHyperlink("Repo", &url.URL{Scheme: "https", Host: c.Course.Url})
	group.Append(header)
	group.Append(description)
	group.Append(url)
	//spacer := layout.NewSpacer()
	//layout := widget.NewVBox(header, spacer, description, spacer, url, spacer, spacer, spacer, spacer)
	c.Container = widget.NewVScrollContainer(group)

	return c
}
