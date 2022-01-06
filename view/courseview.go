package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/eulersexception/glabs-ui/model"
)

// rightCont is a reference to the main part on the right side of the view.
// This reference enables a correct update of the view on changes.
var rightCont *fyne.Container

// CreateCourseAccordion generates a sidebar with the main accordion items which are buttons.
// Clicking on an accordion item unfolds a list of semester for a course.
func CreateCourseAccordion(right *fyne.Container) *widget.Accordion {
	rightCont = right
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

// CreateAddEditDeleteButtonsForCourses populates the sidebar with additional buttons for
// creating and editing courses.
func CreateAddEditButtonsForCourses(acc *fyne.Container) []*widget.Button {
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
				newAcc := CreateCourseAccordion(rightCont)
				acc.Objects[0] = newAcc
				acc.Refresh()
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

	edit := widget.NewButton("Edit", func() {
		w := CreateCourseEditWindow(acc)
		w.Show()
	})

	buttons := make([]*widget.Button, 0)
	buttons = append(buttons, add)
	buttons = append(buttons, edit)

	return buttons
}

// CreateCourseEditWindow opens a new window that shows a table where
// each row contains information, entries and buttons related to a course.
func CreateCourseEditWindow(acc *fyne.Container) fyne.Window {
	w := fyne.CurrentApp().NewWindow("Course Edit")
	courses := model.GetAllCourses()
	names := container.NewVBox()
	entries := container.NewVBox()
	okButtons := container.NewVBox()
	deleteButtons := container.NewVBox()
	editSemestersButtons := container.NewVBox()

	names.Add(widget.NewLabel("Current Path"))
	entries.Add(widget.NewLabel("New Path                   "))
	okButtons.Add(widget.NewLabel(""))
	deleteButtons.Add(widget.NewLabel(""))
	editSemestersButtons.Add(widget.NewLabel(""))

	for _, v := range courses {
		c := v
		old := widget.NewLabel(c.Path)
		new := widget.NewEntry()
		new.SetPlaceHolder("Enter Path")

		okButton := widget.NewButton("OK", func() {
			model.UpdateCourse(c.Path, new.Text)
			newAcc := CreateCourseAccordion(rightCont)
			acc.Objects[0] = newAcc
			acc.Refresh()
		})

		deleteButton := widget.NewButton("Delete", func() {
			semesters := model.GetAllSemestersForCourse(c.Path)

			if len(semesters) > 0 {
				popUp := widget.NewPopUp(container.NewCenter(widget.NewLabel(fmt.Sprintf("Delete existing semesters before deleting course '%s'", c.Path))), w.Canvas())
				popUp.Show()
			} else {
				model.DeleteCourse(c.Path)
				newAcc := CreateCourseAccordion(rightCont)
				acc.Objects[0] = newAcc
				acc.Refresh()
			}
		})

		editSemesterButton := widget.NewButton("Edit Semesters", func() {
			editWindow := CreateSemesterEditWindow(&c, acc)
			editWindow.Show()
		})

		names.Add(old)
		entries.Add(new)
		okButtons.Add(okButton)
		deleteButtons.Add(deleteButton)
		editSemestersButtons.Add(editSemesterButton)
	}

	content := container.NewHBox(names, entries, okButtons, deleteButtons, editSemestersButtons)
	w.SetContent(content)

	return w
}
