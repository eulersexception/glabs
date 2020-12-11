package view

import (
	"fmt"
	"net/url"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
	glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type CourseDetailView struct {
	TabContainer *widget.TabContainer
	Container    *fyne.Container
	Course       *glabsmodel.Course
}

func NewCourseDetailView(course *glabsmodel.Course, tc *widget.TabContainer) *CourseDetailView {
	c := &CourseDetailView{
		Course:       course,
		TabContainer: tc,
	}

	// building body
	group := widget.NewGroup(fmt.Sprintf("Kurs %s", c.Course.Name))
	description := widget.NewLabel(c.Course.Description)
	description.Wrapping = fyne.TextWrapWord
	url := widget.NewHyperlink("Repo", &url.URL{Scheme: "https", Host: c.Course.Url})
	group.Append(description)
	group.Append(url)
	body := widget.NewVScrollContainer(group)
	mainWindowSize := glabsutil.GetMainWindow().Content().Size()
	body.SetMinSize(fyne.NewSize(int(float64(mainWindowSize.Width)*0.8), int(float64(mainWindowSize.Height)*0.8)))

	// buttons at the bottom
	left := makeButtonForSemesterOverview(c.TabContainer, c.Course)
	right := glabsutil.MakeCloseButton(c.TabContainer)
	buttons := glabsutil.MakeButtonGroup(left, right)

	// pack components
	c.Container = glabsutil.MakeScrollableView(body, buttons)

	return c
}

func makeButtonForSemesterOverview(tc *widget.TabContainer, c *glabsmodel.Course) *widget.Button {
	overviewButton := widget.NewButton("Semesterübersicht", func() {
		semesterOverview := NewSemesterOverview(tc, c)
		item := widget.NewTabItem(fmt.Sprintf("Semester %s", c.Name), semesterOverview.Container)
		tc.Append(item)
	})

	return overviewButton
}
