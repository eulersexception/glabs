package view

import (
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	model "github.com/eulersexception/glabs-ui/model"
)

func NewCourseView() {
	// courses := model.GetAllCourses()

}

func semestersByCourseX(content *fyne.Container) map[binding.String][]*widget.Button {
	courses := model.GetAllCourses()
	sort.Slice(courses, func(i int, j int) bool { return courses[i].Path < courses[j].Path })
	semesterByCourses := make(map[binding.String][]*widget.Button)

	for _, v := range courses {
		semesters := createSemesterButtons(v.Path, content)
		path := binding.NewString()
		path.Set(v.Path)
		semesterByCourses[path] = semesters
	}

	return semesterByCourses
}

func createAddEditDeleteButtonsForCoursesX() []*widget.Button {
	add := widget.NewButton("+", func() {})
	edit := widget.NewButton("Edit", func() {})
	delete := widget.NewButton("Delete", func() {})

	buttons := make([]*widget.Button, 0)
	buttons = append(buttons, add)
	buttons = append(buttons, edit)
	buttons = append(buttons, delete)

	return buttons
}
