package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
)

func NewStarterCodeView(starterUrl string) *fyne.Container {
	starterContainer := container.NewVBox()
	var starterCode *widget.Button

	starterCode = widget.NewButton(starterUrl, func() {
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

		var hide *widget.Button

		hide = widget.NewButton("Hide", func() {
			starterData.Hide()
			hide.Hide()
			starterCode.Show()
		})

		starterContainer.Add(hide)
		starterCode.Hide()
	})

	starterContainer.Add(starterCode)

	return starterContainer
}
