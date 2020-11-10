package view

import (
	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type CourseOverview struct {
	Container *widget.ScrollContainer
	Courses   []*glabsmodel.Course
}

func NewCourseOverview(courses []*glabsmodel.Course, tc *widget.TabContainer) *CourseOverview {
	c := &CourseOverview{
		Courses: courses,
	}

	group := widget.NewGroupWithScroller("")
	group.Resize(fyne.NewSize(1100, 750))

	if courses != nil {
		for _, v := range courses {
			label := v.Name
			button := widget.NewButton("Details", func() {
				tc.Append(widget.NewTabItem(label, NewCourseDetailView(v).Container))
			})
			vBox := widget.NewHBox(widget.NewLabel(label), button, layout.NewSpacer())
			group.Append(vBox)
		}
		c.Container = widget.NewVScrollContainer(group)

	} else {
		text := widget.NewTextGrid()
		text.SetText("Hallo Kursuebersicht")
		button := widget.NewButton("Dummy", func() {
			text.SetText("Button clicked")
			text.Refresh()
		})

		c.Container = widget.NewVScrollContainer(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), text, layout.NewSpacer(), button))
	}

	return c
}
