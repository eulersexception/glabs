package view

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/eulersexception/glabs-ui/model"
	"github.com/eulersexception/glabs-ui/util"
)

func CreateHomeView(myWindow fyne.Window) {

	myWindow.Resize(fyne.NewSize(800, 600))

	menueBar := createMenueBar()
	sepLine := canvas.NewLine(color.White)
	menue := container.NewVSplit(menueBar, sepLine)

	right := container.NewVBox(widget.NewLabel("Test"))
	semesterByCourses := createCourseAccordion(right)
	mainAccordion := widget.NewAccordion()

	for coursePath, semesterList := range semesterByCourses {
		semesterItems := container.NewVBox()

		for _, v := range semesterList {
			semesterItems.Add(v)
		}

		semesterItems.Show()
		mainAccordion.Append(widget.NewAccordionItem(coursePath, semesterItems))
	}

	scrollableAccordion := container.NewVScroll(mainAccordion)
	themeButtons := createButtonsforDarkLightMode()
	left := container.NewVSplit(scrollableAccordion, themeButtons)
	left.SetOffset(1)

	split := container.NewHSplit(left, right)
	split.SetOffset(0.5)
	content := container.NewBorder(menue, layout.NewSpacer(), split, layout.NewSpacer())

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func createMenueBar() *container.Split {

	t := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {}),
	)

	menueBar := container.NewHSplit(layout.NewSpacer(), t)
	menueBar.SetOffset(1.0)

	return menueBar
}

func createCourseAccordion(content *fyne.Container) map[string][]*widget.Button {
	courses := model.GetAllCourses()
	semesterByCourses := make(map[string][]*widget.Button)

	for _, v := range courses {
		semesters := createSemesterButtons(v.Path, content)
		semesterByCourses[v.Path] = semesters
	}

	return semesterByCourses
}

func createSemesterButtons(coursePath string, content *fyne.Container) []*widget.Button {
	semesters := model.GetAllSemestersForCourse(coursePath)
	semesterButtons := make([]*widget.Button, 0)

	for _, v := range semesters {
		right := createAssignmentList(v.Path)
		label := widget.NewLabel(fmt.Sprintf("Assignments für %s", v.Path))

		b := widget.NewButton(v.Path, func() {
			content.Objects[0] = container.NewVBox(label, right)
		})
		semesterButtons = append(semesterButtons, b)
	}

	return semesterButtons
}

func createAssignmentList(semesterPath string) *fyne.Container {
	assignments := model.GetAllAssignmentsForSemester(semesterPath)

	util.WarningLogger.Printf("SemesterPath = %s", semesterPath)

	assignmentList := container.NewVBox()

	for _, v := range assignments {
		b := widget.NewButton(v.AssignmentPath, func() {

		})
		assignmentList.Add(b)
	}

	return assignmentList
}

func createButtonsforDarkLightMode() *container.Split {

	lightButton := widget.NewButton("Light", func() {
		fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	})

	darkButton := widget.NewButton("Dark", func() {
		fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	})

	buttonGroup := container.NewHSplit(lightButton, darkButton)

	return buttonGroup
}
