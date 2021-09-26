package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateHomeView(myApp fyne.App) {
	myWindow := myApp.NewWindow("Glabs")
	myWindow.Resize(fyne.NewSize(1200, 600))
	menueBar := createMenueBar()
	sepLine := canvas.NewLine(color.White)
	menue := container.NewVSplit(menueBar, sepLine)

	right := container.NewVBox(widget.NewLabel("Test"))
	changeCourseButtons := createAddEditDeleteButtonsForCourses()
	mainAccordion := createCourseAccordion(right)
	scrollableAccordion := container.NewVScroll(mainAccordion)
	themeButtons := createButtonsforDarkLightMode()
	changeCourseButtonsBox := container.NewVBox(changeCourseButtons[0], changeCourseButtons[1], changeCourseButtons[2])
	buttons := container.NewVBox(changeCourseButtonsBox, themeButtons)
	left := container.NewVSplit(scrollableAccordion, buttons)
	left.SetOffset(1)
	split := container.NewHSplit(left, right)
	split.SetOffset(0.5)
	content := container.NewBorder(menue, layout.NewSpacer(), split, layout.NewSpacer())

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func createMenueBar() *widget.Menu { //*container.Split {
	menu := widget.NewMenu(fyne.NewMenu("Menu",
		fyne.NewMenuItem("Datei", func() {

		}),
	))

	// t := widget.NewToolbar(
	// 	widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
	// 	widget.NewToolbarSeparator(),
	// 	widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	// 	widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
	// 	widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
	// 	widget.NewToolbarSpacer(),
	// 	widget.NewToolbarAction(theme.HelpIcon(), func() {}),
	// )

	//menueBar := container.NewHSplit(layout.NewSpacer(), t)
	//menueBar.SetOffset(1.0)

	return menu //menueBar
}

func createButtonsforDarkLightMode() *fyne.Container {
	var lightButton, darkButton *widget.Button

	lightButton = widget.NewButton("Light", func() {
		fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		lightButton.Disable()
		darkButton.Enable()
	})

	darkButton = widget.NewButton("Dark", func() {
		fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
		darkButton.Disable()
		lightButton.Enable()
	})

	buttonGroup := container.NewVBox(lightButton, darkButton)

	return buttonGroup
}
