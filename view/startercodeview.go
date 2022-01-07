package view

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/eulersexception/glabs-ui/model"
	glabsmodel "github.com/eulersexception/glabs-ui/model"
)

// NewStarterCodeView provides an detail view for a Starter Code.
// It returns a container object that is used as content for an existing window.
func NewStarterCodeView(starterUrl string) *fyne.Container {
	starterContainer := container.NewVBox()

	starter := glabsmodel.GetStarterCode(starterUrl)

	urlLabel := widget.NewLabel("Starter Code URL:")
	url := widget.NewLabel(starter.StarterUrl)

	fromBranchLabel := widget.NewLabel("From Branch:")
	fromBranch := widget.NewLabel(starter.FromBranch)

	protectLabel := widget.NewLabel("Protect to Branch:")
	protect := widget.NewLabel(fmt.Sprintf("%v", starter.ProtectToBranch))

	starterLabels := container.NewVBox(urlLabel, fromBranchLabel, protectLabel)
	starterValues := container.NewVBox(url, fromBranch, protect)

	starterData := container.NewHBox(starterLabels, starterValues)
	starterContainer.Add(starterData)

	return starterContainer
}

// NewStarterCodeOverview provides an overview for all existing Starter Codes in a new window.
func NewStarterCodeOverview() fyne.Window {
	starters := model.GetAllStarterCodes()
	w := fyne.CurrentApp().NewWindow("Overview Clones")
	starterUrls := container.NewVBox(widget.NewLabel("URL"))
	fromBranches := container.NewVBox(widget.NewLabel("From Branch"))
	protects := container.NewVBox(widget.NewLabel("Protect"))
	deleteButtons := container.NewVBox(layout.NewSpacer())
	editButtons := container.NewVBox(layout.NewSpacer())
	assignmentsButtons := container.NewVBox(layout.NewSpacer())

	for _, v := range starters {
		s := v
		starterUrls.Add(widget.NewLabel(s.StarterUrl))
		fromBranches.Add(widget.NewLabel(s.FromBranch))
		protects.Add(widget.NewLabel(strconv.FormatBool(s.ProtectToBranch)))

		deleteButton := widget.NewButton("Delete", func() {
			var warning fyne.Window
			warningText := widget.NewLabel(fmt.Sprintf("Caution - deleting starter code %s. Related Assignments must be updated manually.", s.StarterUrl))
			okDelete := widget.NewButton("Proceed", func() {
				model.DeleteClone(s.StarterUrl)
				w.Close()
				w = NewStarterCodeOverview()
				w.Show()
				warning.Close()
			})
			cancelDelete := widget.NewButton("Cancel", func() {
				warning.Close()
			})

			warning = fyne.CurrentApp().NewWindow(fmt.Sprintf("Delete Starter Code %s", s.StarterUrl))
			buttons := container.NewHBox(okDelete, cancelDelete)
			content := container.NewVBox(warningText, buttons)
			warning.SetContent(content)
			warning.Show()
		})

		deleteButtons.Add(deleteButton)

		editButton := widget.NewButton("Edit", func() {
			editWindow := NewEditStarterCodeWindow(&s, w)
			editWindow.Show()
		})

		editButtons.Add(editButton)

		assignmentButton := widget.NewButton("Assignments", func() {
			assignmentWindow := NewAssignmentOverwievForStarterCode(s.StarterUrl)
			assignmentWindow.Show()
		})

		assignmentsButtons.Add(assignmentButton)
	}

	content := container.NewHBox(starterUrls, fromBranches, protects, editButtons, deleteButtons, assignmentsButtons)
	w.SetContent(content)

	return w
}

// NewEditStarterCodeWindow creates a window to edit a Starter Code.
func NewEditStarterCodeWindow(s *model.StarterCode, w fyne.Window) fyne.Window {
	editWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("Edit Starter Code %s", s.StarterUrl))

	curStarterUrl := widget.NewLabel(s.StarterUrl)
	newStarterUrl := widget.NewEntry()
	newStarterUrl.SetPlaceHolder("Enter new path")

	curBranch := widget.NewLabel(fmt.Sprintf("%s\t\t\t", s.FromBranch))
	newBranch := widget.NewEntry()
	newBranch.SetPlaceHolder("Enter new from branch")

	var t string
	var f string

	if s.ProtectToBranch {
		t = strconv.FormatBool(s.ProtectToBranch)
		f = strconv.FormatBool(!s.ProtectToBranch)
	} else {
		t = strconv.FormatBool(!s.ProtectToBranch)
		f = strconv.FormatBool(s.ProtectToBranch)
	}

	selectGroup := widget.NewRadioGroup([]string{t, f}, func(choice string) {})
	selectGroup.SetSelected(strconv.FormatBool(s.ProtectToBranch))

	okButton := widget.NewButton("OK", func() {
		if newBranch.Text != "Enter new from branch" && newBranch.Text != s.FromBranch && newBranch.Text != "" {
			s.FromBranch = newBranch.Text
		}

		s.ProtectToBranch, _ = strconv.ParseBool(selectGroup.Selected)
		s.UpdateStarterCode()

		if newStarterUrl.Text != "Enter new path" && newStarterUrl.Text != s.StarterUrl && newStarterUrl.Text != "" {
			model.UpdateStarterUrl(s.StarterUrl, newStarterUrl.Text)
		}

		w.Close()
		w = NewStarterCodeOverview()
		w.Show()

		editWindow.Close()
	})

	cancelButton := widget.NewButton("Cancel", func() {
		editWindow.Close()
	})

	paths := container.NewVBox(widget.NewLabel("Starter Urls"), curStarterUrl, newStarterUrl, okButton)
	branches := container.NewVBox(widget.NewLabel("From Branch"), curBranch, newBranch, cancelButton)
	protects := container.NewVBox(widget.NewLabel("Protect to Branch"), selectGroup)
	content := container.NewHBox(paths, branches, protects)

	editWindow.SetContent(content)

	return editWindow
}
