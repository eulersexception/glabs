package main

import (
	// "time"
	// "fmt"

	//"log"

	"fyne.io/fyne/v2/app"
	"github.com/eulersexception/glabs-ui/model"
	"github.com/eulersexception/glabs-ui/util"
	"github.com/eulersexception/glabs-ui/view"
)

func main() {

	util.InitLoggers()
	model.CreateTables()
	model.InitData()
	myApp := app.New()
	view.CreateHomeView(myApp)
	model.DropTables()

	//curTime := fmt.Sprintf("logs_%s", time.Now().String())
	//curTime = curTime[:24]
	// os.Rename("logs.txt", curTime)
	//os.Remove("logs.txt")
}
