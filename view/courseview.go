package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/eulersexception/glabs-ui/model"
)

func NewCourseView() {
	// courses := model.GetAllCourses()
}

func createCourseAccordion(right *fyne.Container) *widget.Accordion {
	semestersByCourse := semestersByCourse(right)
	mainAccordion := widget.NewAccordion()

	for coursePath, semesterList := range semestersByCourse {
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

func createAddEditDeleteButtonsForCourses() []*widget.Button {
	courseDialog := fyne.CurrentApp().NewWindow("Create Course")
	courseDialog.Resize(fyne.NewSize(400, 100))

	add := widget.NewButton("+", func() {
		newCourseLabel := widget.NewLabel("Enter new course path: ")
		newCourseEntry := widget.NewEntry()
		newCourseEntry.SetPlaceHolder("Enter a valid course path")

		ok := widget.NewButton("Add Course", func() {
			if newCourseEntry.Text != "" {
				model.NewCourse(newCourseEntry.Text)
				done := widget.NewPopUp(widget.NewLabel(fmt.Sprintf("New course %s created", newCourseEntry.Text)), courseDialog.Canvas())
				done.Show()
			} else {
				warning := widget.NewPopUp(widget.NewLabel("Enter a valid path"), courseDialog.Canvas())
				warning.Show()
			}
		})
		cancel := widget.NewButton("               Cancel               ", func() {
			courseDialog.Close()
		})

		labels := container.NewVBox(newCourseLabel)
		entries := container.NewVBox(newCourseEntry)
		labels.Add(ok)
		entries.Add(cancel)
		content := container.NewHBox(labels, entries)
		courseDialog.SetContent(content)
		courseDialog.Show()
	})
	edit := widget.NewButton("Edit", func() {})
	delete := widget.NewButton("Delete", func() {})

	buttons := make([]*widget.Button, 0)
	buttons = append(buttons, add)
	buttons = append(buttons, edit)
	buttons = append(buttons, delete)

	return buttons
}
