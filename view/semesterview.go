package view

import (
	"fmt"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/eulersexception/glabs-ui/model"
)

func semestersByCourse(content *fyne.Container) map[binding.String][]*widget.Button {
	courses := model.GetAllCourses()
	sort.Slice(courses, func(i int, j int) bool { return courses[i].Path < courses[j].Path })
	semesterByCourses := make(map[binding.String][]*widget.Button)

	for _, v := range courses {
		c := v
		semesters := createSemesterButtons(c.Path, content)
		path := binding.NewString()
		path.Set(c.Path)
		semesterByCourses[path] = semesters
	}

	return semesterByCourses
}

func createSemesterButtons(coursePath string, content *fyne.Container) []*widget.Button {
	semesters := model.GetAllSemestersForCourse(coursePath)
	sort.Slice(semesters, func(i int, j int) bool { return semesters[i].Path < semesters[j].Path })

	semesterButtons := make([]*widget.Button, 0)

	for _, v := range semesters {
		s := v
		label := widget.NewLabel(fmt.Sprintf("Assignments für %s", s.Path))
		vBox := container.NewVBox(label)
		right := createAssignmentList(s.Path, content, vBox)
		vBox.Add(right)

		b := widget.NewButton(s.Path, func() {
			content.Objects[0] = vBox
		})

		semesterButtons = append(semesterButtons, b)
	}

	return semesterButtons
}

// CreateSemesterEditWindow contains logic to display a list and edit options
// for Semesters related to a specific Course within a new window.
func CreateSemesterEditWindow(c *model.Course, acc *fyne.Container) fyne.Window {
	w := fyne.CurrentApp().NewWindow(fmt.Sprintf("Edit semesters for course %s", c.Path))
	semesters := model.GetAllSemestersForCourse(c.Path)

	// column headers
	labels := container.NewVBox(widget.NewLabel("Semester\t\t"))
	semPaths := container.NewVBox(widget.NewLabel("Semester Path\t\t"))
	coursePaths := container.NewVBox(widget.NewLabel("Course Path\t"))
	okButtons := container.NewVBox(widget.NewLabel(""))
	deleteButtons := container.NewVBox(widget.NewLabel(""))

	// content per row
	for _, v := range semesters {
		s := v

		sLabel := widget.NewLabel(s.Path)
		sPath := widget.NewEntry()
		sPath.SetText(s.Path)
		sPath.SetPlaceHolder(s.Path)

		cPath := widget.NewEntry()
		cPath.SetText(s.CoursePath)
		cPath.SetPlaceHolder(s.CoursePath)

		okButton := widget.NewButton("OK", func() {
			if sPath.Text != s.Path {
				sem := *model.GetSemester(sPath.Text)

				if sem == (model.Semester{}) {
					var warn widget.PopUp
					warnLabel := widget.NewLabel(fmt.Sprintf("Semester %s doesn't exist. Do you want to create it?", sPath.Text))
					warnOk := widget.NewButton("OK", func() {
						model.NewSemester(s.CoursePath, sPath.Text)
						s.UpdateSemesterPath(sPath.Text)
						s.UpdateSemesterPath(sPath.Text)
						newAcc := CreateCourseAccordion(rightCont)
						acc.Objects[0] = newAcc
						acc.Refresh()
						warn.Hide()
					})
					warnCancel := widget.NewButton("Cancel", func() {
						sPath.Text = s.Path
						sPath.SetPlaceHolder(s.Path)
						warn.Hide()
					})

					warnButtons := container.NewHBox(warnOk, warnCancel)
					warnMain := container.NewVBox(warnLabel, warnButtons)
					warn = *widget.NewPopUp(warnMain, w.Canvas())
					warn.Show()

				} else {
					s.UpdateSemesterPath(sPath.Text)
					newAcc := CreateCourseAccordion(rightCont)
					acc.Objects[0] = newAcc
					acc.Refresh()
				}
			}
			if cPath.Text != s.CoursePath {
				course := *model.GetCourse(cPath.Text)

				if course == (model.Course{}) {
					warn := widget.NewPopUp(widget.NewLabel(fmt.Sprintf("Course %s doesn't exist. Please create course before assigning semesters to it", cPath.Text)), w.Canvas())
					warn.Show()
				} else {
					s.UpdateSemester(cPath.Text)
					newAcc := CreateCourseAccordion(rightCont)
					acc.Objects[0] = newAcc
					acc.Refresh()
				}
			}
		})

		deleteButton := widget.NewButton("Delete", func() {
			assignments := model.GetAllAssignmentsForSemester(s.Path)

			if len(assignments) > 0 {
				warn := widget.NewModalPopUp(widget.NewLabel(fmt.Sprintf("Delete existing assignments before deleting semester %s", s.Path)), w.Canvas())
				warn.Show()
			} else {
				model.DeleteSemester(s.Path)
				newAcc := CreateCourseAccordion(rightCont)
				acc.Objects[0] = newAcc
				acc.Refresh()
			}
		})

		labels.Add(sLabel)
		semPaths.Add(sPath)
		coursePaths.Add(cPath)
		okButtons.Add(okButton)
		deleteButtons.Add(deleteButton)
	}

	createLabel := widget.NewLabel("New semester:")
	createSemPath := widget.NewEntry()
	createSemPath.SetPlaceHolder("new semester path")
	createButton := widget.NewButton("Create", func() {
		if createSemPath.Text == "new semester path" || createSemPath.Text == "" {
			warn := widget.NewPopUp(widget.NewLabel("Enter a valid path for semester."), w.Canvas())
			warn.Show()
		} else {
			model.NewSemester(c.Path, createSemPath.Text)
			newAcc := CreateCourseAccordion(rightCont)
			acc.Objects[0] = newAcc
			acc.Refresh()
		}
	})

	labels.Add(layout.NewSpacer())
	semPaths.Add(layout.NewSpacer())
	okButtons.Add(layout.NewSpacer())

	labels.Add(createLabel)
	semPaths.Add(createSemPath)
	okButtons.Add(createButton)

	table := container.NewHBox(labels, semPaths, coursePaths, okButtons, deleteButtons)
	w.SetContent(table)

	return w
}
