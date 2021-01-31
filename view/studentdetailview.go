package view

import (
	"fmt"
	"net/url"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
	glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type StudentDetailView struct {
	TabContainer *widget.TabContainer
	Container    *fyne.Container
	Student      *glabsmodel.Student
}

func NewStudentDetailView(student *glabsmodel.Student, tc *widget.TabContainer) *StudentDetailView {
	s := &StudentDetailView{
		TabContainer: tc,
		Student:      student,
	}

	group := widget.NewGroup(fmt.Sprintf("Details %s", s.Student.Name))
	id := widget.NewLabel(fmt.Sprintf("Matrikelnummer: %d", s.Student.Id))
	name := widget.NewLabel(fmt.Sprintf("Name: %s", s.Student.Name))
	firstName := widget.NewLabel(fmt.Sprintf("Vorname:%s", s.Student.FirstName))
	nickName := widget.NewLabel(fmt.Sprintf("Nickname: %s", s.Student.NickName))
	email := widget.NewHyperlink(s.Student.GetMail(), &url.URL{Scheme: "smtp", Host: s.Student.GetMail()})
	group.Append(id)
	group.Append(name)
	group.Append(firstName)
	group.Append(nickName)
	group.Append(email)
	body := widget.NewVScrollContainer(group)

	left := glabsutil.MakeCloseButton(tc)
	right := glabsutil.MakeCloseButton(tc)
	buttons := glabsutil.MakeButtonGroup(left, right)

	s.Container = glabsutil.MakeScrollableView(body, buttons)

	return s
}
