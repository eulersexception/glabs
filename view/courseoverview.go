package view

import (
	"fmt"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
	glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type CourseOverview struct {
	TabContainer *widget.TabContainer
	Container    *widget.ScrollContainer
	Courses      []*glabsmodel.Course
}

func NewCourseOverview(courses []*glabsmodel.Course, tc *widget.TabContainer) *CourseOverview {
	c := &CourseOverview{
		Courses:      courses,
		TabContainer: tc,
	}

	group := widget.NewGroup("Kursübersicht")

	if courses != nil {
		for _, v := range courses {
			currentCourse := &glabsmodel.Course{
				Name:        v.Name,
				Description: v.Description,
				Url:         v.Url,
				Semesters:   v.Semesters,
			}

			addSemesters(5, currentCourse)

			button := widget.NewButton("Details", func() {
				tc.Append(widget.NewTabItem(currentCourse.Name, NewCourseDetailView(currentCourse, tc).Container))
				tc.Remove(tc.CurrentTab())
			})
			label := widget.NewLabel(v.Name)
			desc := widget.NewLabel(v.Description)
			desc.Wrapping = fyne.TextTruncate
			line := widget.NewHBox(label, desc, layout.NewSpacer(), button)
			group.Append(line)
		}

		c.Container = widget.NewVScrollContainer(group)

	} else {
		text := widget.NewLabel("Aktuell keine Kurse vorhanden")
		left := MakeButtonForCourseCreation(tc)
		right := glabsutil.MakeCloseButton(tc)
		buttons := glabsutil.MakeButtonGroup(left, right)

		c.Container = widget.NewVScrollContainer(fyne.NewContainerWithLayout(layout.NewHBoxLayout(), text, layout.NewSpacer(), buttons))
	}

	return c
}

func addSemesters(n int, course *glabsmodel.Course) {
	year := 2000

	for i := 0; i < n; i++ {

		// summer
		name := fmt.Sprintf("%s Sommersemester %d", course.Name, (year + i))
		url := "www.google.de"
		summer := &glabsmodel.Semester{Name: name, Url: url, Course: course}
		course.AddSemesterToCourse(summer)

		// winter
		name = fmt.Sprintf("%s Wintersemester %d/%d", course.Name, (year + i), (year + i + 1))
		winter := &glabsmodel.Semester{Name: name, Url: url, Course: course}
		course.AddSemesterToCourse(winter)
	}
}
