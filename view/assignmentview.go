package view

import (
	"fmt"
	"strconv"

	"github.com/eulersexception/glabs-ui/model"
	glabsmodel "github.com/eulersexception/glabs-ui/model"

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
	Assignment *glabsmodel.Assignment
}

func NewAssignmentView(assignment *glabsmodel.Assignment) *AssignmentView {

	aView := &AssignmentView{
		Assignment: assignment,
	}

	a := model.GetAssignment(assignment.AssignmentPath)

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

	aView.Content = container.NewBorder(labelsValues, buttons, nil, nil)

	return aView
}
