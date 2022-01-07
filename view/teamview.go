package view

import (
	"fmt"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/eulersexception/glabs-ui/model"
	"github.com/eulersexception/glabs-ui/util"
)

// NewTeamView creates a window containing detailed information about a Team
// including buttons to edit or delete Team data and list Students who are members.
func NewTeamView(a *model.Assignment) fyne.Window {
	teamWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("Teams %s", a.AssignmentPath))

	names := container.NewVBox()
	inputs := container.NewVBox()
	edits := container.NewVBox()
	deletes := container.NewVBox()
	students := container.NewVBox()
	content := container.NewHBox()

	colName := widget.NewLabel("Team\t\t\t\t")
	colInput := widget.NewLabel("\t\t\t\t\t")
	colEmpty := widget.NewLabel("")

	names.Add(colName)
	inputs.Add(colInput)
	edits.Add(colEmpty)
	deletes.Add(colEmpty)
	students.Add(colEmpty)

	teams := model.GetTeamsForAssignment(a.AssignmentPath)

	for i, v := range teams {
		util.WarningLogger.Printf("%d. loop over teams - name = %s\n", i, v.Name)
	}

	sort.Slice(teams, func(i int, j int) bool { return teams[i].Name < teams[j].Name })

	for _, v := range teams {
		team := v
		name := widget.NewLabel(team.Name)
		input := widget.NewEntry()
		input.SetPlaceHolder(team.Name)
		input.SetText(team.Name)

		edit := widget.NewButton("Edit", func() {
			newName := input.Text
			team.UpdateTeam(newName)

			util.WarningLogger.Printf("Input.Value = %s, edited Team = %s", newName, team.Name)

			newTeamWindow := NewTeamView(a)
			newTeamWindow.Show()
			teamWindow.Close()
		})

		delete := widget.NewButton("Delete", func() {
			studs := model.GetStudentsForTeam(team.Name)

			if len(studs) > 0 {
				warningContent := container.NewVBox()
				warningContent.Add(widget.NewLabel("There are students organized in this team.\nRemove them first from the team before deleting team."))

				warning := widget.NewModalPopUp(
					warningContent,
					teamWindow.Canvas(),
				)

				closeWarning := widget.NewButton("OK", func() {
					warning.Hide()
				})

				warningContent.Add(closeWarning)

				warning.Show()
			} else {
				model.DeleteTeam(team.Name)
				newTeamWindow := NewTeamView(a)
				newTeamWindow.Show()
				teamWindow.Close()
			}
		})

		student := widget.NewButton("Students", func() {
			studentsView := NewStudentView(&team)
			studentsView.Show()
		})

		names.Add(name)
		inputs.Add(input)
		edits.Add(edit)
		deletes.Add(delete)
		students.Add(student)
	}

	content.Add(names)
	content.Add(inputs)
	content.Add(edits)
	content.Add(deletes)
	content.Add(students)

	teamWindow.SetContent(content)

	return teamWindow
}
