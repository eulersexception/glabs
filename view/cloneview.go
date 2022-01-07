package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	model "github.com/eulersexception/glabs-ui/model"
)

// NewCloneView provides an detail view for a Clone.
// It returns a container object that is used as content for an existing window.
func NewCloneView(localPath string) *fyne.Container {
	cloneContainer := container.NewVBox()

	clone := model.GetClone(localPath)

	pathLabel := widget.NewLabel("Local Path:")
	path := widget.NewLabel(clone.LocalPath)

	branchLabel := widget.NewLabel("Branch:")
	branch := widget.NewLabel(clone.Branch)

	cloneLabels := container.NewVBox(pathLabel, branchLabel)
	cloneValues := container.NewVBox(path, branch)

	cloneData := container.NewHBox(cloneLabels, cloneValues)
	cloneContainer.Add(cloneData)

	return cloneContainer
}

// NewCloneOverview provides an overview for all existing Clones in a new window.
func NewCloneOverview() fyne.Window {
	clones := model.GetAllClones()
	w := fyne.CurrentApp().NewWindow("Overview Clones")
	localPaths := container.NewVBox(widget.NewLabel("Local Paths"))
	branches := container.NewVBox(widget.NewLabel("Branches"))
	deleteButtons := container.NewVBox(layout.NewSpacer())
	editButtons := container.NewVBox(layout.NewSpacer())
	assignmentsButtons := container.NewVBox(layout.NewSpacer())

	for _, v := range clones {
		c := v
		localPaths.Add(widget.NewLabel(c.LocalPath))
		branches.Add(widget.NewLabel(c.Branch))

		deleteButton := widget.NewButton("Delete", func() {
			var warning fyne.Window
			warningText := widget.NewLabel(fmt.Sprintf("Caution - deleting clone %s. Related Assignments must be updated manually.", c.LocalPath))
			okDelete := widget.NewButton("Proceed", func() {
				model.DeleteClone(c.LocalPath)
				w.Close()
				w = NewCloneOverview()
				w.Show()
				warning.Close()
			})
			cancelDelete := widget.NewButton("Cancel", func() {
				warning.Close()
			})

			warning = fyne.CurrentApp().NewWindow(fmt.Sprintf("Delete Clone %s", c.LocalPath))
			buttons := container.NewHBox(okDelete, cancelDelete)
			content := container.NewVBox(warningText, buttons)
			warning.SetContent(content)
			warning.Show()
		})

		deleteButtons.Add(deleteButton)

		editButton := widget.NewButton("Edit", func() {
			editWindow := NewEditCloneWindow(&c, w)
			editWindow.Show()
		})

		editButtons.Add(editButton)

		assignmentButton := widget.NewButton("Assignments", func() {
			assignmentWindow := NewAssignmentOverwievForClone(c.LocalPath)
			assignmentWindow.Show()
		})

		assignmentsButtons.Add(assignmentButton)
	}

	content := container.NewHBox(localPaths, branches, editButtons, deleteButtons, assignmentsButtons)
	w.SetContent(content)

	return w
}

// NewEditCloneWindow creates a window to edit a Clone.
func NewEditCloneWindow(c *model.Clone, w fyne.Window) fyne.Window {
	editWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("Edit Clone %s", c.LocalPath))

	curPath := widget.NewLabel(c.LocalPath)
	newPath := widget.NewEntry()
	newPath.SetPlaceHolder("Enter new path")

	curBranch := widget.NewLabel(fmt.Sprintf("%s\t\t\t", c.Branch))
	newBranch := widget.NewEntry()
	newBranch.SetPlaceHolder("Enter new branch")

	okButton := widget.NewButton("OK", func() {
		if newBranch.Text != "Enter new branch" && newBranch.Text != c.Branch && newBranch.Text != "" {
			c.Branch = newBranch.Text
			c.UpdateClone()
		}

		if newPath.Text != "Enter new path" && newPath.Text != c.LocalPath && newPath.Text != "" {
			model.UpdateClonePath(c.LocalPath, newPath.Text)
		}

		w.Close()
		w = NewCloneOverview()
		w.Show()

		editWindow.Close()
	})

	cancelButton := widget.NewButton("Cancel", func() {
		editWindow.Close()
	})

	paths := container.NewVBox(widget.NewLabel("Local Path"), curPath, newPath, okButton)
	branches := container.NewVBox(widget.NewLabel("Branch"), curBranch, newBranch, cancelButton)
	content := container.NewHBox(paths, branches)

	editWindow.SetContent(content)

	return editWindow
}
