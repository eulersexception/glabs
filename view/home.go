package view

import (
	"fmt"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
	glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var dialogButtonSize = fyne.NewSize(80, 30)
var dialogWindowSize = fyne.NewSize(400, 150)

type HomeView struct {
	App       fyne.App
	Window    fyne.Window
	TabBar    *widget.TabContainer
	News      *widget.TextGrid
	Buttons   *widget.SplitContainer
	Container *fyne.Container
	MainMenu  *fyne.MainMenu
}

func NewHomeview() *HomeView {
	h := &HomeView{
		App: app.New(),
	}

	h.Window = h.App.NewWindow("GLabs")
	h.MainMenu = glabsutil.MakeMainMenu()
	h.Window.SetMainMenu(h.MainMenu)
	h.TabBar = widget.NewTabContainer()

	h.Buttons = widget.NewHSplitContainer(makeButtonForCourseOverview(h.TabBar), MakeButtonForCourseCreation(h.TabBar))
	h.News = makeNewsContent()
	h.Container = fyne.NewContainerWithLayout(layout.NewVBoxLayout(), layout.NewSpacer(), h.News, layout.NewSpacer(), layout.NewSpacer(), h.Buttons)
	item := widget.NewTabItem("Home", h.Container)
	h.TabBar.Append(item)
	h.TabBar.SelectTab(item)
	h.Window.SetContent(h.TabBar)
	h.Window.Resize(fyne.NewSize(1200, 800))

	return h
}

func makeNewsContent() *widget.TextGrid {
	newsTicker := widget.NewTextGrid()
	newsTicker.SetText("Here you can see a lot of updates for all the repos")

	return newsTicker
}

func MakeButtonForCourseCreation(tc *widget.TabContainer) *widget.Button {

	createCourseButton := widget.NewButton("Kurs erstellen", func() {
		mainWindow := glabsutil.GetMainWindow()

		// initializing new window, creating entries and form
		w := fyne.CurrentApp().NewWindow("Kurs erstellen")
		courseNameEntry := widget.NewEntry()
		courseNameEntry.SetPlaceHolder("Kurs Bezeichnung")
		courseDescription := widget.NewMultiLineEntry()
		courseDescription.SetPlaceHolder("Kursbeschreibung eingeben")
		form := widget.NewForm(widget.NewFormItem("Kurs", courseNameEntry), widget.NewFormItem("Beschreibung", courseDescription))

		// ok button
		doneButton := widget.NewButton("Fertig", func() {
			doneWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("Kurs \"%s\" erstellen?", courseNameEntry.Text))
			message := widget.NewLabel(fmt.Sprintf("Soll der Kurs \"%s\" erstellt werden?", courseNameEntry.Text))
			messageBox := fyne.NewContainerWithLayout(layout.NewCenterLayout(), message)
			ok := widget.NewButton("OK", func() {
				courseDescription.SetReadOnly(true)
				courseNameEntry.SetReadOnly(true)
				mainWindow.Content().Refresh()
				mainWindow.RequestFocus()
				doneWindow.Close()
				w.Close()
			})
			cancel := widget.NewButton("Abbrechen", func() {
				doneWindow.Close()
				w.RequestFocus()
			})

			buttonBox := widget.NewHSplitContainer(ok, cancel)
			doneWindow.SetContent(widget.NewVBox(layout.NewSpacer(), messageBox, layout.NewSpacer(), buttonBox))
			doneWindow.Resize(dialogWindowSize)
			doneWindow.CenterOnScreen()
			doneWindow.Show()
		})

		// cancel button
		cancelButton := glabsutil.MakeCancelButtonForDialog(mainWindow, w)

		// wrapping widgets in container
		buttons := widget.NewHSplitContainer(doneButton, cancelButton)
		container := widget.NewVScrollContainer(widget.NewVBox(form, layout.NewSpacer(), buttons))
		w.SetContent(container)
		w.Resize(fyne.NewSize(800, 600))

		// displaying on center of screen
		w.CenterOnScreen()
		w.Show()
	})

	return createCourseButton
}

func makeButtonForCourseOverview(tc *widget.TabContainer) *widget.Button {
	overviewButton := widget.NewButton("Kursübersicht", func() {
		courseOverview := NewCourseOverview(createDummyCourses(30), tc)
		item := widget.NewTabItem("Kursübersicht", courseOverview.Container)
		tc.Append(item)
		item.Content.Refresh()
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
