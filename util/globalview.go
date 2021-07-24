package util

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var dialogButtonSize = fyne.NewSize(80, 30)
var dialogWindowSize = fyne.NewSize(400, 150)

func MakeCloseButton() *widget.Button {
	closeButton := widget.NewButton("Schließen", func() {

	})

	return closeButton
}

func MakeButtonGroup(left *widget.Button, right *widget.Button) *container.Split {
	return container.NewHSplit(left, right)
}

func MakeScrollableView(body *container.Scroll, buttons *container.Scroll) *fyne.Container {
	//mainWindowSize := GetMainWindow().Content().Size()
	body.SetMinSize(fyne.NewSize(5, 10))
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), body, layout.NewSpacer(), buttons)
}

func MakeCancelButtonForDialog(mainWindow fyne.Window, currentWindow fyne.Window) *widget.Button {
	return widget.NewButton("Abbrechen", func() {
		warningMessageWindow := fyne.CurrentApp().NewWindow("Vorgang abbrechen")
		warning := widget.NewLabel("Klicken Sie \"Ok\", um die Eingabe abzubrechen.\nDie Daten werden nicht gespeichert. Fortfahren?")
		warningBox := fyne.NewContainerWithLayout(layout.NewCenterLayout(), warning)
		ok := widget.NewButton("Ok", func() {
			mainWindow.RequestFocus()
			warningMessageWindow.Close()
			currentWindow.Close()
		})
		cancel := widget.NewButton("Abbrechen", func() {
			warningMessageWindow.Close()
			currentWindow.RequestFocus()
		})
		buttonBox := container.NewHSplit(ok, cancel)
		container := container.NewVBox(layout.NewSpacer(), warningBox, layout.NewSpacer(), buttonBox)
		warningMessageWindow.SetContent(container)
		warningMessageWindow.Resize(dialogWindowSize)
		warningMessageWindow.CenterOnScreen()
		warningMessageWindow.Show()
	})
}

func GetMainWindow() fyne.Window {
	var mainWindow fyne.Window
	windows := fyne.CurrentApp().Driver().AllWindows()

	for _, v := range windows {
		if v.Title() == "GLabs" {
			mainWindow = v
		}
	}

	return mainWindow
}

func MakeMainMenu() *fyne.MainMenu {
	newItem := fyne.NewMenuItem("Neu", func() {})
	prefItem := fyne.NewMenuItem("Einstellungen", func() {})
	themeItem := fyne.NewMenuItem("Theme wechseln", func() {})
	darkItem := fyne.NewMenuItem("Dark", func() {
		fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	})
	lightItem := fyne.NewMenuItem("Light", func() {
		fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	})
	themeItem.ChildMenu = fyne.NewMenu("", darkItem, lightItem)

	menu := fyne.NewMainMenu(fyne.NewMenu("Fyne", newItem, prefItem, themeItem))

	return menu
}
