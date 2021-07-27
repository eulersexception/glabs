package view

import (
	"fmt"
	"image/color"
	"sort"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/eulersexception/glabs-ui/model"
	"github.com/eulersexception/glabs-ui/util"
)

func CreateHomeView(myApp fyne.App) {
	myWindow := myApp.NewWindow("Glabs")
	myWindow.Resize(fyne.NewSize(1200, 600))
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

	util.WarningLogger.Printf("SemesterPath = %s", semesterPath)

	assignmentList := container.NewVBox()

	for _, v := range assignments {
		assignmentDetails := createAssignmentDetailBox(v.AssignmentPath)
		b := widget.NewButton(v.AssignmentPath, func() {
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

	buttonGroup := container.NewHBox(lightButton, darkButton)

	return buttonGroup
}

func createAssignmentDetailBox(assignmentPath string) *fyne.Container {
	a := model.GetAssignment(assignmentPath)

	aPathLabel := widget.NewLabel("AssignmentPath:")
	aPathValue := widget.NewLabel(a.AssignmentPath)

	aSemPathLabel := widget.NewLabel("SemesterPath:")
	aSemPathValue := widget.NewLabel(a.SemesterPath)

	aPerLabel := widget.NewLabel("Per:")
	aPerValue := widget.NewLabel(a.Per)

	aDescLabel := widget.NewLabel("Description:")
	aDescValue := widget.NewLabel(a.Description)

	aContRegLabel := widget.NewLabel("ContainerRegistry:")
	aContRegValue := widget.NewLabel(strconv.FormatBool(a.ContainerRegistry))

	aStarterCodeLabel := widget.NewLabel("StarterCode:")
	aStarterCodeValue := widget.NewButton(a.StarterUrl, func() {})

	aCloneLabel := widget.NewLabel("Clone:")
	aCloneValue := widget.NewButton(a.LocalPath, func() {})

	labels := container.NewVBox(aPathLabel, aSemPathLabel, aPerLabel, aDescLabel, aContRegLabel, aStarterCodeLabel, aCloneLabel)
	values := container.NewVBox(aPathValue, aSemPathValue, aPerValue, aDescValue, aContRegValue, aStarterCodeValue, aCloneValue)

	editButton := widget.NewButton("Edit", func() {
		// editLabels := container.NewVBox(aPathLabel, aSemPathLabel, aPerLabel, aDescLabel, aContRegLabel, aStarterCodeLabel, aCloneLabel)
		pathEntry := widget.NewEntry()
		pathEntry.SetPlaceHolder(aPathValue.Text)
		editPath := widget.NewFormItem(aPathLabel.Text, pathEntry)

		semPathEntry := widget.NewEntry()
		semPathEntry.SetPlaceHolder(aSemPathValue.Text)
		editSemPath := widget.NewFormItem(aSemPathLabel.Text, semPathEntry)

		perEntry := widget.NewEntry()
		perEntry.SetPlaceHolder(aPerValue.Text)
		editPer := widget.NewFormItem(aPerLabel.Text, perEntry)

		descEntry := widget.NewMultiLineEntry()
		descEntry.SetPlaceHolder(aDescValue.Text)
		editDesc := widget.NewFormItem(aDescLabel.Text, descEntry)

		contRegRadioButtons := widget.NewRadioGroup([]string{"true", "false"}, func(string) {})
		contRegRadioButtons.SetSelected(aContRegValue.Text)
		editContReg := widget.NewFormItem(aContRegLabel.Text, contRegRadioButtons)

		starterEntry := widget.NewEntry()
		starterEntry.SetPlaceHolder(aStarterCodeValue.Text)
		editStarterCode := widget.NewFormItem(aStarterCodeLabel.Text, starterEntry)

		cloneEntry := widget.NewEntry()
		cloneEntry.SetPlaceHolder(aContRegValue.Text)
		editClone := widget.NewFormItem(aCloneLabel.Text, cloneEntry)

		editForm := widget.NewForm(editPath, editSemPath, editPer, editDesc, editContReg, editStarterCode, editClone)
		editWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("Edit %s", aPathValue.Text))

		cancelButton := widget.NewButton("Abbrechen", func() {

			var editModal *widget.PopUp
			warning := canvas.NewText("Eingabe abbrechen? Erfasste Daten werden nicht gespeichert", theme.TextColor())
			no := widget.NewButton("Nein", func() {
				editModal.Hide()
			})
			yes := widget.NewButton("Ja", func() {
				editWindow.Close()
			})

			buttonGroup := container.NewHBox(yes, no)
			cancelContent := container.NewBorder(warning, buttonGroup, nil, nil)
			editModal = widget.NewModalPopUp(cancelContent, editWindow.Canvas())
			editModal.Show()
		})

		saveButton := widget.NewButton("Speichern", func() {
			var editModal *widget.PopUp
			warning := canvas.NewText("Eingabe sichern? Erfasste Daten werden gespeichert", theme.TextColor())
			no := widget.NewButton("Nein", func() {
				editModal.Hide()
			})
			yes := widget.NewButton("Ja", func() {
				editWindow.Close()
			})

			buttonGroup := container.NewHBox(yes, no)
			cancelContent := container.NewBorder(warning, buttonGroup, nil, nil)
			editModal = widget.NewModalPopUp(cancelContent, editWindow.Canvas())
			editModal.Show()
		})

		buttons := container.NewHBox(layout.NewSpacer(), cancelButton, saveButton, layout.NewSpacer())
		editContent := container.NewBorder(editForm, buttons, nil, nil)

		editWindow.SetContent(editContent)
		editWindow.Resize(fyne.NewSize(800, 400))
		editWindow.Show()
	})

	editLabelsValues := container.NewHBox(labels, values)

	return container.NewBorder(editLabelsValues, editButton, nil, nil)
}
