package view

import (
	"fmt"

	glabsmodel "github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type StudentView struct {
	Content *fyne.Container
	Student *glabsmodel.Student
}

func NewStudentView(s *glabsmodel.Student) *StudentView {
	title := widget.NewLabel("Student Info")

	matrikelLabel := widget.NewLabel("Matrikelnummer:")
	matrikelNummer := widget.NewLabel(fmt.Sprintf("%v", s.MatrikelNr))

	nameLabel := widget.NewLabel("Name:")
	name := widget.NewLabel(s.Name)

	firstNameLabel := widget.NewLabel("Vorname:")
	firstName := widget.NewLabel(s.FirstName)

	nickNameLabel := widget.NewLabel("Nickname:")
	nickName := widget.NewLabel(s.NickName)

	emailLabel := widget.NewLabel("Email:")
	email := widget.NewLabel(s.Email)

	labels := container.NewVBox(matrikelLabel, nameLabel, firstNameLabel, nickNameLabel, emailLabel)
	values := container.NewVBox(matrikelNummer, name, firstName, nickName, email)

	data := container.NewHBox(labels, values)
	content := container.NewVBox(title, data)

	studView := &StudentView{
		Content: content,
		Student: s,
	}

	return studView
}
