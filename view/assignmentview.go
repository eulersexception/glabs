package view

import (
	"fmt"
	"image/color"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
	"github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type AssignmentView struct {
	Content    *fyne.Container
	Assignment *glabsmodel.Assignment
}

func NewAssignmentView(a *glabsmodel.Assignment) *AssignmentView {

	aView := &AssignmentView{
		Assignment: a,
	}

	title := widget.NewLabel("Assignment Info")
	sepLine := canvas.NewLine(color.White)

	assignmentLabel := widget.NewLabel("Assignment Path:")
	assignment := widget.NewLabel(a.AssignmentPath)

	semesterLabel := widget.NewLabel("Semester Path:")
	semester := widget.NewLabel(a.SemesterPath)

	perLabel := widget.NewLabel("Per:")
	per := widget.NewLabel(a.Per)

	descriptionLabel := widget.NewLabel("Description:")
	description := widget.NewLabel(a.Description)

	registryLabel := widget.NewLabel("Container Registry:")
	registry := widget.NewLabel(fmt.Sprintf("%v", a.ContainerRegistry))

	starterLabel := widget.NewLabel("Starter Code:")
	starterContainer := container.NewVBox()
	var starterCode *widget.Button

	starterCode = widget.NewButton(a.StarterUrl, func() {
		starter := glabsmodel.GetStarterCode(a.StarterUrl)

		urlLabel := widget.NewLabel("Starter Code URL:")
		url := widget.NewLabel(starter.Url)

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

	cloneLabel := widget.NewLabel("Clone:")
	cloneContainer := container.NewVBox()
	var clone *widget.Button

	clone = widget.NewButton(a.LocalPath, func() {
		innerClone := glabsmodel.GetClone(a.LocalPath)

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

	labels := container.NewVBox(assignmentLabel, semesterLabel, perLabel, descriptionLabel, registryLabel, starterLabel, cloneLabel)
	values := container.NewVBox(assignment, semester, per, description, registry, starterContainer, cloneContainer)

	data := container.NewHBox(labels, values)

	left := MakeEditButton(a, *semester, *per, *description)
	right := util.MakeCloseButton()
	buttons := util.MakeButtonGroup(left, right)

	content := container.NewVBox(title, sepLine, data, layout.NewSpacer(), sepLine, buttons)

	aView.Content = content

	return aView
}

func MakeEditButton(a *glabsmodel.Assignment, labels ...widget.Label) *widget.Button {

	form := widget.NewForm()
	//entries := widget.NewFormItem()

	button := widget.NewButton("Edit", func() {
		for _, v := range labels {
			item := widget.NewFormItem(v.Text, widget.NewEntry())
			form.AppendItem(item)
		}

		newWindow := fyne.CurrentApp().NewWindow("Edit Assignment")

		newWindow.SetContent(form)
		newWindow.Show()
	})

	return button
}
