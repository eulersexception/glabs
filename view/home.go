package view

import (
	"fmt"
	"image/color"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/eulersexception/glabs-ui/model"
)

func CreateHomeView(myApp fyne.App) {
	myWindow := myApp.NewWindow("Glabs")
	myWindow.Resize(fyne.NewSize(1200, 600))
	menueBar := createMenueBar()
	sepLine := canvas.NewLine(color.White)
	menue := container.NewVSplit(menueBar, sepLine)

	right := container.NewVBox(widget.NewLabel("Test"))
	changeCourseButtons := createAddEditDeleteButtonsForCourses()
	mainAccordion := createCourseAccordion(right)
	scrollableAccordion := container.NewVScroll(mainAccordion)
	themeButtons := createButtonsforDarkLightMode()
	changeCourseButtonsBox := container.NewVBox(changeCourseButtons[0], changeCourseButtons[1], changeCourseButtons[2])
	buttons := container.NewVBox(changeCourseButtonsBox, themeButtons)
	left := container.NewVSplit(scrollableAccordion, buttons)
	left.SetOffset(1)
	split := container.NewHSplit(left, right)
	split.SetOffset(0.5)
	content := container.NewBorder(menue, layout.NewSpacer(), split, layout.NewSpacer())

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func createMenueBar() *widget.Menu { //*container.Split {
	menu := widget.NewMenu(fyne.NewMenu("Menu",
		fyne.NewMenuItem("Datei", func() {

		}),
	))

	// t := widget.NewToolbar(
	// 	widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
	// 	widget.NewToolbarSeparator(),
	// 	widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	// 	widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
	// 	widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
	// 	widget.NewToolbarSpacer(),
	// 	widget.NewToolbarAction(theme.HelpIcon(), func() {}),
	// )

	//menueBar := container.NewHSplit(layout.NewSpacer(), t)
	//menueBar.SetOffset(1.0)

	return menu //menueBar
}

func createCourseAccordion(right *fyne.Container) *widget.Accordion {
	semestersByCourse := semestersByCourse(right)
	mainAccordion := widget.NewAccordion()

	for coursePath, semesterList := range semestersByCourse {
		semesterItems := container.NewVBox()

		for _, v := range semesterList {
			semesterItems.Add(v)
		}

		semesterItems.Show()
		courseTitle, _ := coursePath.Get()

		aItem := widget.NewAccordionItem(courseTitle, semesterItems)

		mainAccordion.Append(aItem)
	}

	return mainAccordion
}

func semestersByCourse(content *fyne.Container) map[binding.String][]*widget.Button {
	courses := model.GetAllCourses()
	semesterByCourses := make(map[binding.String][]*widget.Button)

	for _, v := range courses {
		semesters := createSemesterButtons(v.Path, content)
		path := binding.NewString()
		path.Set(v.Path)
		semesterByCourses[path] = semesters
	}

	return semesterByCourses
}

func createAddEditDeleteButtonsForCourses() []*widget.Button {
	add := widget.NewButton("+", func() {})
	edit := widget.NewButton("Edit", func() {})
	delete := widget.NewButton("Delete", func() {})

	buttons := make([]*widget.Button, 0)
	buttons = append(buttons, add)
	buttons = append(buttons, edit)
	buttons = append(buttons, delete)

	return buttons
}

func createSemesterButtons(coursePath string, content *fyne.Container) []*widget.Button {
	semesters := model.GetAllSemestersForCourse(coursePath)
	sort.Slice(semesters, func(i int, j int) bool { return semesters[i].Path < semesters[j].Path })

	semesterButtons := make([]*widget.Button, 0)

	for _, v := range semesters {
		label := widget.NewLabel(fmt.Sprintf("Assignments für %s", v.Path))
		vBox := container.NewVBox(label)
		right := createAssignmentList(v.Path, content, vBox)
		vBox.Add(right)

		b := widget.NewButton(v.Path, func() {
			content.Objects[0] = vBox
		})

		semesterButtons = append(semesterButtons, b)
	}

	return semesterButtons
}

func createAssignmentList(semesterPath string, content *fyne.Container, right *fyne.Container) *fyne.Container {
	assignments := model.GetAllAssignmentsForSemester(semesterPath)
	sort.Slice(assignments, func(i int, j int) bool {
		return assignments[i].AssignmentPath < assignments[j].AssignmentPath
	})

	assignmentList := container.NewVBox()

	for _, v := range assignments {
		a := v
		view := NewAssignmentView(a.AssignmentPath)
		assignmentDetails := view.Content
		b := widget.NewButton(a.AssignmentPath, func() {
			newRight := container.NewHBox(right, assignmentDetails)
			content.Objects[0] = newRight
		})
		assignmentList.Add(b)
	}

	return assignmentList
}

func createButtonsforDarkLightMode() *fyne.Container {
	var lightButton, darkButton *widget.Button

	lightButton = widget.NewButton("Light", func() {
		fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		lightButton.Disable()
		darkButton.Enable()
	})

	darkButton = widget.NewButton("Dark", func() {
		fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		darkButton.Disable()
		lightButton.Enable()
	})

	buttonGroup := container.NewVBox(lightButton, darkButton)

	return buttonGroup
}
