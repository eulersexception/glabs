package view

import (
	"fmt"
	"net/url"

	glabsmodel "github.com/eulersexception/glabs-ui/model"
	glabsutil "github.com/eulersexception/glabs-ui/util"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type StudentOverview struct {
	TabContainer *widget.TabContainer
	Container    *widget.ScrollContainer
	Students     []*glabsmodel.Student
	Team         *glabsmodel.Team
}

func NewStudentOverview(team *glabsmodel.Team, tc *widget.TabContainer) *StudentOverview {
	s := &StudentOverview{
		TabContainer: tc,
		Team:         team,
		Students:     team.Students,
	}

	group := widget.NewGroup(fmt.Sprintf("Mitgliederübersichet %s", s.Team.Name))

	if s.Students != nil {
		for _, v := range s.Students {
			currentStudent := &glabsmodel.Student{
				Id:        v.Id,
				Team:      v.Team,
				FirstName: v.FirstName,
				Name:      v.Name,
				NickName:  v.NickName,
				Email:     v.Email,
			}
			currentStudent.Mail(v.GetMail())

			button := widget.NewButton("Details", func() {
				tc.Append(widget.NewTabItem(currentStudent.NickName, NewStudentDetailView(currentStudent, tc).Container))
				tc.Remove(tc.CurrentTab())
			})
			label := widget.NewLabel(fmt.Sprintf("%s %s, Nick: %s", currentStudent.FirstName, currentStudent.Name, currentStudent.NickName))
			url := widget.NewHyperlink(currentStudent.GetMail(), &url.URL{Scheme: "smtp", Host: currentStudent.GetMail()})
			url.Wrapping = fyne.TextTruncate
			line := widget.NewHBox(label, url, layout.NewSpacer(), button)
			group.Append(line)
		}

		s.Container = widget.NewVScrollContainer(group)
	} else {
		text := widget.NewLabel("Aktuell keine Mitglieder vorhanden")
		left := MakeButtonForCourseCreation(tc)
		right := glabsutil.MakeCloseButton(tc)
		buttons := glabsutil.MakeButtonGroup(left, right)

		s.Container = widget.NewVScrollContainer(fyne.NewContainerWithLayout(layout.NewHBoxLayout(), text, layout.NewSpacer(), buttons))
	}

	return s
}
