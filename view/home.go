package view

import (
	"fmt"

	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type HomeView struct {
	App       fyne.App
	Window    fyne.Window
	TabBar    *widget.TabContainer
	News      *widget.TextGrid
	Buttons   *widget.SplitContainer
	Container *fyne.Container
}

func NewHomeview() *HomeView {
	h := &HomeView{
		App: app.New(),
	}

	h.Window = h.App.NewWindow("GLabs")
	h.TabBar = widget.NewTabContainer()
	h.Buttons = widget.NewHSplitContainer(addButtonForCourseOverview(h.TabBar), addButtonForCourseCreation(h.TabBar))
	h.News = addNewsContent()
	h.Container = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), layout.NewSpacer(), h.News, layout.NewSpacer(), layout.NewSpacer(), h.Buttons)
	item := widget.NewTabItem("Home", h.Container)
	h.TabBar.Append(item)
	h.TabBar.SelectTab(item)
	h.Window.SetContent(h.TabBar)
	h.Window.Resize(fyne.NewSize(1200, 800))

	return h
}

func addNewsContent() *widget.TextGrid {
	newsTicker := widget.NewTextGrid()
	newsTicker.SetText("Here you can see a lot of updates for all the repos")

	return newsTicker
}

func addButtonForCourseCreation(tc *widget.TabContainer) *widget.Button {
	createCourseButton := widget.NewButton("Kurs erstellen", func() {
		courseNameEntry := widget.NewEntry()
		courseNameEntry.SetPlaceHolder("Kurs Bezeichnung")
		courseDescription := widget.NewMultiLineEntry()
		courseDescription.SetPlaceHolder("Kursbeschreibung eingeben")
		form := widget.NewForm(widget.NewFormItem("Kurs", courseNameEntry), widget.NewFormItem("Beschreibung", courseDescription))
		button := widget.NewButton("Fertig", func() {
			courseDescription.SetText(fmt.Sprintf("%s%s", courseDescription.Text, "\n\nDaten gespeichert!"))
			courseDescription.SetReadOnly(true)
			courseNameEntry.SetReadOnly(true)
		})
		button.Resize(fyne.NewSize(100, 20))
		container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), form, layout.NewSpacer(), button)
		item := widget.NewTabItem("Kurs erstellen", container)
		tc.Append(item)
	})

	return createCourseButton
}

func addButtonForCourseOverview(tc *widget.TabContainer) *widget.Button {
	overviewButton := widget.NewButton("Kursübersicht", func() {
		courseOverview := NewCourseOverview(createDummyCourses(30), tc)
		tc.Append(widget.NewTabItem("Kursübersicht", courseOverview.Container))
	})

	return overviewButton
}

func createDummyCourses(n int) []*glabsmodel.Course {
	courses := make([]*glabsmodel.Course, 0)

	for i := 0; i < n; i++ {
		name := fmt.Sprintf("AlgoDat %02d", i)
		description := fmt.Sprintf("Algorithmen pur Teil %d", i)
		course := &glabsmodel.Course{Name: name, Description: description, Url: "google.de", Semesters: nil}
		courses = append(courses, course)
	}

	return courses
}
