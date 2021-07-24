package view

import (
	glabsmodel "github.com/eulersexception/glabs-ui/model"
	//glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type StudentDetailView struct {
	TabContainer *widget.TabContainer
	Container    *fyne.Container
	Student      *glabsmodel.Student
}

func NewStudentDetailView(student *glabsmodel.Student, tc *widget.TabContainer) *StudentDetailView {
	return nil
}
