package model

import (
	"fmt"

	util "github.com/eulersexception/glabs-ui/util"
)

type Student struct {
	Id        int
	Name      string
	FirstName string
	NickName  string
	email     string
}

func NewStudent(id int, name string, firstName string) *Student {
	if name == "" || firstName == "" {
		fmt.Println("Please provide valid name or first name.")
		return nil
	}

	student := &Student{
		Id:        id,
		Name:      name,
		FirstName: firstName,
	}

	return student
}

func (s *Student) Mail(email string) bool {
	if util.IsValidMail(email) {
		s.email = email
		return true
	} else {
		return false
	}
}

func (s Student) GetMail() string {
	return s.email
}
