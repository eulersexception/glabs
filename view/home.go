package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// CreateHomeView is the entrypoint for the application where higher view elements
// like the sidebar are displayed.
func CreateHomeView(myApp fyne.App) fyne.Window {
	myWindow := myApp.NewWindow("glabs")
	myWindow.Resize(fyne.NewSize(1200, 600))

	right := container.NewVBox(widget.NewLabel("Welcome to glabs"))
	mainAccordion := CreateCourseAccordion(right)
	accordion := container.NewVBox(mainAccordion)
	changeCourseButtons := CreateAddEditButtonsForCourses(accordion)
	themeButtons := CreateButtonsforDarkLightMode()
	changeCourseButtonsBox := container.NewVBox(
		changeCourseButtons[0],
		changeCourseButtons[1],
	)

	cloneOverViewButton := widget.NewButton("Clones", func() {
		cloneOverview := NewCloneOverview()
		cloneOverview.Show()
	})

	starterCodeOverViewButton := widget.NewButton("Starter Codes", func() {
		starterCodeOverView := NewStarterCodeOverview()
		starterCodeOverView.Show()
	})

	buttons := container.NewVBox(changeCourseButtonsBox, themeButtons, cloneOverViewButton, starterCodeOverViewButton)
	left := container.NewVBox(accordion, layout.NewSpacer(), buttons)
	split := container.NewHBox(left, right)
	content := container.NewBorder(layout.NewSpacer(), layout.NewSpacer(), split, layout.NewSpacer())

	myWindow.SetContent(content)

	return myWindow
}

// CreateButtonsForDarkLightMode contains logic to switch between light and dark mode.
// Default is dark mode. Light mode is not working properly.
func CreateButtonsforDarkLightMode() *fyne.Container {
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

	darkButton.Disable()
	buttonGroup := container.NewVBox(lightButton, darkButton)

	return buttonGroup
}
