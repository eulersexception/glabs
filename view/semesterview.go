package view

import (
	"fmt"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/eulersexception/glabs-ui/model"
)

func NewSemesterView() {

}

func semestersByCourse(content *fyne.Container) map[binding.String][]*widget.Button {
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

func createSemesterButtons(coursePath string, content *fyne.Container) []*widget.Button {
	semesters := model.GetAllSemestersForCourse(coursePath)
	sort.Slice(semesters, func(i int, j int) bool { return semesters[i].Path < semesters[j].Path })

	semesterButtons := make([]*widget.Button, 0)

	for _, v := range semesters {
		label := widget.NewLabel(fmt.Sprintf("Assignments für %s", v.Path))
		vBox := container.NewVBox(label)
		right := createAssignmentList(v.Path, content, vBox)
		vBox.Add(right)

		b := widget.NewButton(v.Path, func() {
			content.Objects[0] = vBox
		})

		semesterButtons = append(semesterButtons, b)
	}

	return semesterButtons
}
