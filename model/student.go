package model

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"

	util "github.com/eulersexception/glabs-ui/util"
)

type Student struct {
	Id        uint32
	Team      *Team
	Name      string
	FirstName string
	NickName  string
	Email     string
}

func NewStudent(team *Team, name string, firstName string, nickName string, email string, id uint32) *Student {
	if name == "" || firstName == "" {
		fmt.Println("Please provide valid name or first name.")
		return nil
	}

	if Mail(email) == false {
		fmt.Println("Please provide valid email address.")
		return nil
	}

	student := &Student{
		Id:        id,
		Team:      team,
		NickName:  nickName,
		Email:     email,
		Name:      name,
		FirstName: firstName,
	}

	student.SetStudent()

	return student
}

func Mail(email string) bool {
	if util.IsValidMail(email) {
		return true
	} else {
		return false
	}
}

func (s Student) GetMail() string {
	return s.Email
}

func (s Student) encodeStudent() []byte {
	data, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}

	return data
}

func decodeStudent(data []byte) Student {
	var s Student
	err := json.Unmarshal(data, &s)

	if err != nil {
		panic(err)
	}

	return s
}

func (s Student) AddToTeam(t *Team) {
	s.Team = t
	s.SetStudent()
}

func (s Student) SetStudent() {
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	k := make([]byte, 4)
	binary.LittleEndian.PutUint32(k, s.Id)
	v := s.encodeStudent()

	err = db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(k), []byte(v))
		err = txn.SetEntry(e)

		return err
	})
}

func GetStudent(id uint32) Student {
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var s Student

	err = db.View(func(txn *badger.Txn) error {

		k := make([]byte, 4)
		binary.LittleEndian.PutUint32(k, id)

		item, err := txn.Get([]byte(k))

		if err != nil {
			log.Fatal(err)
		}

		err = item.Value(func(val []byte) error {
			s = decodeStudent(val)
			//fmt.Println(fmt.Sprintf("Key = %s, Value = %s", item.String(), string(val)))
			return err
		})

		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	return s
}

func DeleteStudent(id uint32) error {
	// s := GetStudent(id)
	// t := s.Team.RemoveStudentFromTeam(s)
	// t.SetTeam()

	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	k := make([]byte, 4)
	binary.LittleEndian.PutUint32(k, id)

	err = db.Update(func(txn *badger.Txn) error {
		e := txn.Delete(k)

		return e
	})

	return err
}

func (s Student) PrintData() {
	if s.Team != nil {
		fmt.Printf("-------------------\nTeam:\t\t%s\nName:\t\t%s %s\nNick:\t\t%s\nMailTo:\t\t%s\nId:\t\t%d\n",
			s.Team.Name, s.FirstName, s.Name, s.NickName, s.Email, s.Id)
	} else {
		fmt.Printf("-------------------\nName:\t\t%s %s\nNick:\t\t%s\nMailTo:\t\t%s\nId:\t\t%d\n",
			s.FirstName, s.Name, s.NickName, s.Email, s.Id)
	}
}
