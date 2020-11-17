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

	group := widget.NewGroup("")

	if courses != nil {
		for _, v := range courses {
			x := &glabsmodel.Course{
				Name:        v.Name,
				Description: v.Description,
				Url:         v.Url,
				Semesters:   v.Semesters,
			}
			button := widget.NewButton("Details", func() {
				tc.Append(widget.NewTabItem(x.Name, NewCourseDetailView(x).Container))
				tc.Remove(tc.CurrentTab())
			})
			label := widget.NewLabel(v.Name)
			desc := widget.NewLabel(v.Description)
			desc.Wrapping = fyne.TextTruncate
			line := widget.NewHBox(label, desc, layout.NewSpacer(), button, layout.NewSpacer(), layout.NewSpacer())
			group.Append(line)
		}

		c.Container = widget.NewVScrollContainer(group)

	} else {
		text := widget.NewTextGrid()
		text.SetText("Hallo Kursuebersicht")
		button := widget.NewButton("Dummy", func() {
			text.SetText("Button clicked")
			text.Refresh()
		})

		c.Container = widget.NewVScrollContainer(fyne.NewContainerWithLayout(layout.NewHBoxLayout(), text, layout.NewSpacer(), button))
	}

	return c
}
