package main

import (

	// "fyne.io/fyne/v2"

	model "github.com/eulersexception/glabs-ui/model"
)

var starter = &model.StarterCode{
	Url:             "git@gitlab.lrz.de:vss/startercode/startercodeB1.git",
	FromBranch:      "ws20.1",
	ProtectToBranch: true,
}

var clone = &model.Clone{
	LocalPath: "/Users/obraun/lectures/vss/labs/gitlab/semester/ob-20ws/blatt1",
	Branch:    "develop",
}

var a = &model.Assignment{
	AssignmentPath:    "blatt1",
	SemesterPath:      "semester/ob-20ws",
	Per:               "group",
	Description:       "Blatt 1, Verteilte Softwaresysteme, WS 20/21",
	ContainerRegistry: true,
	LocalPath:         clone.LocalPath,
	StarterUrl:        starter.Url,
}

func main() {

	model.CreateTables()
	model.InitData()

}
