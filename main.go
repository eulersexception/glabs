package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/eulersexception/glabs-ui/model"
	"github.com/eulersexception/glabs-ui/view"

	util "github.com/eulersexception/glabs-ui/util"
)

func main() {
	util.InitLoggers()
	model.CreateTables()
	model.InitData()
	myApp := app.New()
	view.CreateHomeView(myApp).ShowAndRun()
	model.DropTables()
}
