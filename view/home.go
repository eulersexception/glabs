package view

import (
	"fmt"
	"image/color"
	"sort"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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

func createCourseAccordion(right *fyne.Container) *widget.Accordion {
	semesterByCourse := semesterByCourse(right)
	mainAccordion := widget.NewAccordion()

	for coursePath, semesterList := range semesterByCourse {
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

func semesterByCourse(content *fyne.Container) map[binding.String][]*widget.Button {
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
	add := widget.NewButton("Add", func() {})
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

	util.WarningLogger.Printf("SemesterPath = %s", semesterPath)

	assignmentList := container.NewVBox()

	for _, v := range assignments {
		assignmentDetails := createAssignmentDetailBox2(v.AssignmentPath)
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

	buttonGroup := container.NewVBox(lightButton, darkButton)

	return buttonGroup
}

func createAssignmentDetailBox(assignmentPath string) *fyne.Container {
	a := model.GetAssignment(assignmentPath)

	aPathLabel := widget.NewLabel("AssignmentPath:")
	aPathStr := binding.NewString()
	aPathStr.Set(a.AssignmentPath)

	aPathValue := widget.NewLabelWithData(aPathStr)

	aSemPathLabel := widget.NewLabel("SemesterPath:")
	aSemPathStr := binding.NewString()
	aSemPathStr.Set(a.SemesterPath)
	aSemPathValue := widget.NewLabelWithData(aSemPathStr)

	aPerLabel := widget.NewLabel("Per:")
	aPerStr := binding.NewString()
	aPerStr.Set(a.Per)
	aPerValue := widget.NewLabelWithData(aPerStr)

	aDescLabel := widget.NewLabel("Description:")
	aDescStr := binding.NewString()
	aDescStr.Set(a.Description)
	aDescValue := widget.NewLabelWithData(aDescStr)

	aContRegLabel := widget.NewLabel("ContainerRegistry:")
	aContRegBool := binding.NewBool()
	aContRegBool.Set(a.ContainerRegistry)
	aContRegValue := widget.NewLabelWithData(binding.BoolToString(aContRegBool))

	aStarterCodeLabel := widget.NewLabel("StarterCode:")
	aStarterCodeValue := widget.NewButton(a.StarterUrl, func() {})

	aCloneLabel := widget.NewLabel("Clone:")
	aCloneValue := widget.NewButton(a.LocalPath, func() {})

	labels := container.NewVBox(aPathLabel, aSemPathLabel, aPerLabel, aDescLabel, aContRegLabel, aStarterCodeLabel, aCloneLabel)
	values := container.NewVBox(aPathValue, aSemPathValue, aPerValue, aDescValue, aContRegValue, aStarterCodeValue, aCloneValue)

	editButton := widget.NewButton("Edit", func() {
		// editLabels := container.NewVBox(aPathLabel, aSemPathLabel, aPerLabel, aDescLabel, aContRegLabel, aStarterCodeLabel, aCloneLabel)
		pathEntry := widget.NewEntryWithData(aPathStr)
		pathEntry.SetPlaceHolder(aPathValue.Text)
		editPath := widget.NewFormItem(aPathLabel.Text, pathEntry)

		semPathEntry := widget.NewEntryWithData(aSemPathStr)
		semPathEntry.SetPlaceHolder(aSemPathValue.Text)
		editSemPath := widget.NewFormItem(aSemPathLabel.Text, semPathEntry)

		perEntry := widget.NewEntryWithData(aPerStr)
		perEntry.SetPlaceHolder(aPerValue.Text)
		editPer := widget.NewFormItem(aPerLabel.Text, perEntry)

		descEntry := widget.NewEntryWithData(aDescStr)
		descEntry.SetPlaceHolder(aDescValue.Text)
		editDesc := widget.NewFormItem(aDescLabel.Text, descEntry)

		contRegRadioButtons := widget.NewRadioGroup([]string{"true", "false"}, func(choice string) {
			b, _ := strconv.ParseBool(choice)
			aContRegBool.Set(b)
		})
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
				// TODO: Add cancel logic to reset fields
				aPathStr.Set(a.AssignmentPath)
				aSemPathStr.Set(a.SemesterPath)
				aPerStr.Set(a.Per)
				aDescStr.Set(a.Description)
				aContRegBool.Set(a.ContainerRegistry)

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
				newPath := pathEntry.Text
				newSemPath := semPathEntry.Text
				newPer := perEntry.Text
				newDesc := descEntry.Text
				newContReg, _ := strconv.ParseBool(contRegRadioButtons.Selected)
				// TODO: implement value extraction logic for starter code and clone

				aPathStr.Set(newPath)
				aSemPathStr.Set(newSemPath)
				aPerStr.Set(newPer)
				aDescStr.Set(newDesc)
				aContRegBool.Set(newContReg)

				if newPath != a.AssignmentPath {
					model.UpdateAssignmentPath(a.AssignmentPath, newPath)
				}

				newAssignment := &model.Assignment{
					AssignmentPath:    newPath,
					SemesterPath:      newSemPath,
					Per:               newPer,
					Description:       newDesc,
					ContainerRegistry: newContReg,
					StarterUrl:        aStarterCodeValue.Text,
					LocalPath:         aCloneValue.Text,
				}

				newAssignment.UpdateAssignment()

				for _, v := range fyne.CurrentApp().Driver().AllWindows() {
					v.Content().Refresh()
				}

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

	labelsValues := container.NewHBox(labels, values)

	teamsButton := widget.NewButton("Teams", func() {
		w := CreateNewTeamView(a)
		w.Show()
	})

	buttons := container.NewVBox(editButton, teamsButton)

	return container.NewBorder(labelsValues, buttons, nil, nil)
}

func createAssignmentDetailBox2(assignmentPath string) *fyne.Container {
	a := model.GetAssignment(assignmentPath)
	view := NewAssignmentView(a)

	return view.Content

}
