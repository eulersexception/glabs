package view

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type AssignmentView struct {
	Content    *fyne.Container
	Assignment *model.Assignment
}

// NewAssignmentView generates an AssignmentView, which contains details for a specific
// Assignment including buttons to retrieve information about participating Teams, related Clone and Starter Code.
func NewAssignmentView(path string) *AssignmentView {
	a := model.GetAssignment(path)

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

	labels := container.NewVBox(aPathLabel, aSemPathLabel, aPerLabel, aDescLabel, aContRegLabel)
	values := container.NewVBox(aPathValue, aSemPathValue, aPerValue, aDescValue, aContRegValue)

	editButton := widget.NewButton("Edit", func() {
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

		editForm := widget.NewForm(editPath, editSemPath, editPer, editDesc, editContReg)
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

				a.AssignmentPath = newPath
				a.SemesterPath = newSemPath
				a.Per = newPer
				a.Description = newDesc
				a.ContainerRegistry = newContReg

				a.UpdateAssignment()

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

	aStarterButton := widget.NewButton("Starter Code", func() {
		starterWindow := fyne.CurrentApp().NewWindow((fmt.Sprintf("Starter Code for %s", a.StarterUrl)))
		starterWindow.SetContent(NewStarterCodeView(a.StarterUrl))
		starterWindow.Show()
	})

	aCloneButton := widget.NewButton("Clone", func() {
		cloneWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("Clone for %s", a.LocalPath))
		cloneWindow.SetContent(NewCloneView(a.LocalPath))
		cloneWindow.Show()
	})

	teamsButton := widget.NewButton("Teams", func() {
		w := NewTeamView(a)
		w.Show()
	})

	labelsValues := container.NewHBox(labels, values)
	buttons := container.NewVBox(editButton, aStarterButton, aCloneButton, teamsButton)
	aView := &AssignmentView{}
	aView.Content = container.NewBorder(labelsValues, buttons, nil, nil)

	return aView
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

// NewAssignmentOverwievForClone creates a window that contains
// a list of all Assignments for a Clone.
func NewAssignmentOverwievForClone(clonePath string) fyne.Window {
	w := fyne.CurrentApp().NewWindow(fmt.Sprintf("Assignments for Clone %s", clonePath))
	assignments := model.GetAllAssignmentsForClone(clonePath)
	paths := container.NewVBox(widget.NewLabel("Assignment Paths"))
	details := container.NewVBox(layout.NewSpacer())

	for _, v := range assignments {
		a := v

		paths.Add(widget.NewLabel(a.AssignmentPath))
		detail := widget.NewButton("Details", func() {
			detailsWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("Details Assignment %s", a.AssignmentPath))
			detailsWindow.SetContent(NewAssignmentView(a.AssignmentPath).Content)
			fyne.CurrentApp().Driver().AllWindows()[0].SetContent(CreateHomeView(fyne.CurrentApp()).Content())
			fyne.CurrentApp().Driver().AllWindows()[0].Content().Refresh()
			detailsWindow.Show()
		})

		details.Add(detail)
	}

	w.SetContent(container.NewHBox(paths, details))

	return w
}

// NewAssignmentOverwievForStarterCode creates a window that contains
// a list of all Assignments for a Starter Code.
func NewAssignmentOverwievForStarterCode(starterUrl string) fyne.Window {
	w := fyne.CurrentApp().NewWindow(fmt.Sprintf("Assignments for Starter Code %s", starterUrl))
	assignments := model.GetAllAssignmentsForStarterCode(starterUrl)
	paths := container.NewVBox(widget.NewLabel("Assignment Paths"))
	details := container.NewVBox(layout.NewSpacer())

	for _, v := range assignments {
		a := v
		paths.Add(widget.NewLabel(a.AssignmentPath))
		detail := widget.NewButton("Details", func() {
			detailsWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("Details Assignment %s", a.AssignmentPath))
			detailsWindow.SetContent(NewAssignmentView(a.AssignmentPath).Content)
			fyne.CurrentApp().Driver().AllWindows()[0].SetContent(CreateHomeView(fyne.CurrentApp()).Content())
			fyne.CurrentApp().Driver().AllWindows()[0].Content().Refresh()
			detailsWindow.Show()
		})

		details.Add(detail)
	}

	w.SetContent(container.NewHBox(paths, details))

	return w
}
