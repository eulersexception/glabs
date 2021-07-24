package view

import (

	// "fyne.io/fyne/v2"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	model "github.com/eulersexception/glabs-ui/model"
)

func CreateHomeView(myApp fyne.App) {

	//myApp := app.New()
	myWindow := myApp.NewWindow("Glabs")
	// myWindow.Resize(fyne.NewSize(800.0, 600.0))

	menueBar := createMenueBar()
	sepLine := canvas.NewLine(color.White)
	sepLine.StrokeWidth = 0.2
	menue := container.NewVSplit(menueBar, sepLine)

	content := container.NewBorder(menue, nil, nil, nil)

	courseButtons := createCourseButtons()
	courseAccordion := widget.NewAccordion()

	for _, v := range courseButtons {
		courseAccordion.Append(widget.NewAccordionItem(v.Text, v))
	}

	courseOne := widget.NewAccordion()

	semesterList := container.NewVScroll(courseOne)

	//content := container.NewBorder(tools, nil, semesterList, canvas.NewText("", color.White))

	right1 := canvas.NewText("Sommersemester 2016", color.White)
	right2 := canvas.NewText("Wintersemester 2016/2017", color.White)
	right3 := canvas.NewText("Sommersemester 2017", color.White)
	// right4 := canvas.NewText("Wintersemester 2017/2018", color.White)
	// right5 := canvas.NewText("Sommersemester 2018", color.White)

	testContent1 := widget.NewButton("SoSe 2016", func() {
		content.Objects[2] = right1
	})

	testContent2 := widget.NewButton("WiSe 2016/2017", func() {
		content.Objects[2] = right2
	})

	testContent3 := widget.NewButton("SoSe 2017", func() {
		content.Objects[2] = right3
	})

	testContent4 := widget.NewButton("WiSe 2017/2018", func() {
		assignment, _ := model.NewAssignment(a.AssignmentPath, a.SemesterPath, a.Per,
			a.Description, a.ContainerRegistry, a.LocalPath,
			clone.Branch, a.StarterUrl, starter.FromBranch,
			starter.ProtectToBranch)
		content.Objects[2] = NewAssignmentView(assignment).Content
	})

	testContent5 := widget.NewButton("SoSe 2018", func() {
		stud := &model.Student{
			MatrikelNr: 12345,
			Name:       "Max",
			FirstName:  "Mustermann",
			NickName:   "Maxl",
			Email:      "max@muster.mann",
		}
		content.Objects[2] = NewStudentView(stud).Content
	})

	semOne := widget.NewAccordionItem("AlgoDat1", fyne.NewContainerWithLayout(layout.NewVBoxLayout(), testContent1, testContent2, testContent3, testContent4, testContent5))
	semTwo := widget.NewAccordionItem("AlgoDat2", fyne.NewContainerWithLayout(layout.NewVBoxLayout(), testContent1, testContent2, testContent3, testContent4, testContent5))
	semThree := widget.NewAccordionItem("VSS", fyne.NewContainerWithLayout(layout.NewVBoxLayout(), testContent1, testContent2, testContent3, testContent4, testContent5))
	semFour := widget.NewAccordionItem("Softwareentwicklung 1", fyne.NewContainerWithLayout(layout.NewVBoxLayout(), testContent1, testContent2, testContent3, testContent4, testContent5))

	courseOne.Append(semOne)
	courseOne.Append(semTwo)
	courseOne.Append(semThree)
	courseOne.Append(semFour)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}

func createMenueBar() *container.Split {

	t := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	menueBar := container.NewHSplit(layout.NewSpacer(), t)
	menueBar.SetOffset(1.0)

	return menueBar
}

func createCourseButtons() []*widget.Button {

	courses := model.GetAllCourses()
	courseButtons := make([]*widget.Button, 0)

	for _, v := range courses {
		b := widget.NewButton(v.Path, func() {})
		courseButtons = append(courseButtons, b)
	}

	return courseButtons
}

func createSemesterButtons(coursePath string) []*widget.Button {

	semesters := model.GetAllSemestersForCourse(coursePath)
	semesterButtons := make([]*widget.Button, 0)

	for _, v := range semesters {
		b := widget.NewButton(v.Path, func() {})
		semesterButtons = append(semesterButtons, b)
	}

	return semesterButtons
}
