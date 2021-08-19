package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
)

func NewCloneView(localPath string) *fyne.Container {
	cloneContainer := container.NewVBox()
	var clone *widget.Button

	clone = widget.NewButton(localPath, func() {
		innerClone := glabsmodel.GetClone(localPath)

		pathLabel := widget.NewLabel("Local Path:")
		path := widget.NewLabel(innerClone.LocalPath)

		branchLabel := widget.NewLabel("Branch:")
		branch := widget.NewLabel(innerClone.Branch)

		cloneLabels := container.NewVBox(pathLabel, branchLabel)
		cloneValues := container.NewVBox(path, branch)

		cloneData := container.NewHBox(cloneLabels, cloneValues)
		cloneContainer.Add(cloneData)

		var hide *widget.Button

		hide = widget.NewButton("Hide", func() {
			cloneData.Hide()
			hide.Hide()
			clone.Show()
		})

		cloneContainer.Add(hide)
		clone.Hide()
	})

	cloneContainer.Add(clone)

	return cloneContainer
}
