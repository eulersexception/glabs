package view

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/eulersexception/glabs-ui/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// NewStudentView returns a window with details about a Student and
// corresponding buttons for deleting Student and editing data.
func NewStudentView(team *model.Team) fyne.Window {
	studentWindow := fyne.CurrentApp().NewWindow(fmt.Sprintf("Students in Team %s", team.Name))
	students := model.GetStudentsForTeam(team.Name)
	sort.Slice(students, func(i int, j int) bool { return students[i].MatrikelNr < students[j].MatrikelNr })

	matrikels := container.NewVBox()
	nickNames := container.NewVBox()
	names := container.NewVBox()
	firstNames := container.NewVBox()
	emails := container.NewVBox()
	edits := container.NewVBox()
	deletes := container.NewVBox()
	content := container.NewHBox()

	colMatrikel := widget.NewLabel("Matrikelnummer")
	colNickName := widget.NewLabel("Nickname\t")
	colName := widget.NewLabel("Name\t\t\t")
	colFirstName := widget.NewLabel("First Name\t")
	colEmail := widget.NewLabel("Email\t\t\t")
	colEmpty := widget.NewLabel("")

	matrikels.Add(colMatrikel)
	nickNames.Add(colNickName)
	names.Add(colName)
	firstNames.Add(colFirstName)
	emails.Add(colEmail)
	edits.Add(colEmpty)
	deletes.Add(colEmpty)

	for _, v := range students {
		stud := v
		matrikelEntry := widget.NewEntry()
		matrikelEntry.SetPlaceHolder(fmt.Sprintf("%d", stud.MatrikelNr))
		matrikelEntry.SetText(fmt.Sprintf("%d", stud.MatrikelNr))

		nickEntry := widget.NewEntry()
		nickEntry.SetPlaceHolder(stud.NickName)
		nickEntry.SetText(stud.NickName)

		nameEntry := widget.NewEntry()
		nameEntry.SetPlaceHolder(stud.Name)
		nameEntry.SetText(stud.Name)

		firstNameEntry := widget.NewEntry()
		firstNameEntry.SetPlaceHolder(stud.FirstName)
		firstNameEntry.SetText(stud.FirstName)

		emailEntry := widget.NewEntry()
		emailEntry.SetPlaceHolder(stud.Email)
		emailEntry.SetText(stud.Email)

		edit := widget.NewButton("Edit", func() {
			newMatrikelNr, err := strconv.ParseInt(matrikelEntry.Text, 10, 64)

			if err != nil {
				allWindows := fyne.CurrentApp().Driver().AllWindows()
				var warning *widget.PopUp
				warningCheck := widget.NewButton("OK", func() {
					warning.Hide()
				})
				warningContent := container.NewBorder(widget.NewLabel("Enter a valid decimal number as Matrikelnummer."), warningCheck, nil, nil)
				warning = widget.NewPopUp(warningContent, allWindows[len(allWindows)-1].Canvas())
				warning.Show()
			}

			newStud := &model.Student{
				MatrikelNr: stud.MatrikelNr,
				NickName:   nickEntry.Text,
				Name:       nameEntry.Text,
				FirstName:  firstNameEntry.Text,
				Email:      emailEntry.Text,
			}

			newStud.UpdateStudent()

			if stud.MatrikelNr != newMatrikelNr {
				model.UpdateMatrikelNummer(stud.MatrikelNr, newMatrikelNr)
			}

			newStudentWindowContent := NewStudentView(team).Content()
			studentWindow.SetContent(newStudentWindowContent)
			studentWindow.Content().Refresh()
		})

		delete := widget.NewButton("Delete", func() {
			model.DeleteStudent(stud.MatrikelNr)
			newStudentWindow := NewStudentView(team)
			newStudentWindow.Show()
			studentWindow.Close()
		})

		matrikels.Add(matrikelEntry)
		nickNames.Add(nickEntry)
		names.Add(nameEntry)
		firstNames.Add(firstNameEntry)
		emails.Add(emailEntry)
		edits.Add(edit)
		deletes.Add(delete)
	}

	content.Add(matrikels)
	content.Add(nickNames)
	content.Add(names)
	content.Add(firstNames)
	content.Add(emails)
	content.Add(edits)
	content.Add(deletes)
	studentWindow.SetContent(content)

	return studentWindow
}
