package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
)

func NewCloneView(localPath string) *fyne.Container {
	cloneContainer := container.NewVBox()

	clone := glabsmodel.GetClone(localPath)

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
