package model

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"

	util "github.com/eulersexception/glabs-ui/util"
)

// Student - Id (Matrikelnummer) is the primary key. All fields are public and
// Getter or Setter functions relate to database operations.
type Student struct {
	Id        uint32
	Team      *Team
	Name      string
	FirstName string
	NickName  string
	Email     string
}

type JSONStudent struct {
	Id        uint32 `json:"id"`
	TeamId    string `json:"teamid"`
	Name      string `json:"name"`
	FirstName string `json:"firstname"`
	NickName  string `json:"nickname"`
	Email     string `json:"email"`
}

func NewJSONStudent(s *Student) JSONStudent {

	var teamId string

	if s.Team == nil {
		teamId = ""
	} else {
		teamId = s.Team.Name
	}

	return JSONStudent{
		s.Id,
		teamId,
		s.Name,
		s.FirstName,
		s.NickName,
		s.Email,
	}
}

func (js JSONStudent) Student() *Student {
	//team, err := GetTeam(js.TeamId)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return &Student{
		js.Id,
		nil,
		js.Name,
		js.FirstName,
		js.NickName,
		js.Email,
	}
}

// NewStudent creates a new student and stores the object in DB.
// Arguments id of type uint32 and strings for name and firstName must not be empty and a well formed email must be provided.
// If a student with given id exists already in DB, the existing dataset will be overwritten.
// Returns a pointer to a new student and a message string. If provided arguments are invalid the message will not be empty.
func NewStudent(team *Team, name string, firstName string, nickName string, email string, id uint32) (*Student, string) {

	if name == "" || firstName == "" {
		res := "\n+++ Please provide valid name or first name.\n"
		return nil, res
	}

	if Mail(email) == false {
		res := "\n+++ Please provide valid email address.\n"
		return nil, res
	}

	stud := &Student{
		Id:        id,
		Team:      team,
		NickName:  nickName,
		Email:     email,
		Name:      name,
		FirstName: firstName,
	}

	stud.setStudent()

	return stud, ""
}

// Mail checks if the given string is a well formed email address.
// Returns true or false.
func Mail(email string) bool {
	if util.IsValidMail(email) {
		return true
	} else {
		return false
	}
}

// GetMail of student.
// Returns email string.
func (s *Student) GetMail() string {
	return s.Email
}

func (s *Student) encodeStudent() []byte {
	data, err := json.Marshal(NewJSONStudent(s))

	if err != nil {
		panic(err)
	}

	return data
}

func decodeStudent(data []byte) *Student {
	var js JSONStudent
	err := json.Unmarshal(data, &js)

	if err != nil {
		panic(err)
	}

	return js.Student()
}

// UpdateStudent changes a students record in DB.
// Returns an error if the update fails.
func (s *Student) UpdateStudent() error {
	_, err := GetStudent(s.Id)

	if err != nil {
		log.Printf("\n+++ Update of student with id %d failed while checking if student exists.\n+++ %s\n", s.Id, err.Error())
		return err
	}

	err = s.setStudent()

	return err
}

// This function updates student record in DB. An update could be a creation or edition of a record.
func (s *Student) setStudent() error {
	db := util.GetDB()
	defer db.Close()

	k := make([]byte, 4)
	binary.BigEndian.PutUint32(k, s.Id)
	v := []byte(s.encodeStudent())

	err := db.Update(func(txn *badger.Txn) error {
		e := txn.Set(k, v)

		return e
	})

	return err
}

// GetStudent fetches student from DB with an argument of type uint32 as id.
// Returns an error if fetch fails or a pointer to the student.
func GetStudent(id uint32) (*Student, error) {
	db := util.GetDB()
	defer db.Close()

	var s *Student

	err := db.View(func(txn *badger.Txn) error {

		k := make([]byte, 4)
		binary.BigEndian.PutUint32(k, id)

		item, err := txn.Get([]byte(k))

		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			s = decodeStudent(val)
			//fmt.Println(fmt.Sprintf("Key = %s, Value = %s", item.String(), string(val)))
			return err
		})

		return err
	})

	if err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteStudent removes a student by id (uint32) from DB.
// Returns an error if operation fails.
func DeleteStudent(id uint32) error {
	db := util.GetDB()
	defer db.Close()

	k := make([]byte, 4)
	binary.BigEndian.PutUint32(k, id)

	err := db.Update(func(txn *badger.Txn) error {
		e := txn.Delete(k)
		return e
	})

	return err
}

// PrintData outputs a human readable string for students data.
func (s *Student) PrintData() {
	if s.Team != nil {
		fmt.Printf("-------------------\nTeam:\t\t%s\nName:\t\t%s %s\nNick:\t\t%s\nMailTo:\t\t%s\nId:\t\t%d\n",
			s.Team.Name, s.FirstName, s.Name, s.NickName, s.Email, s.Id)
	} else {
		fmt.Printf("-------------------\nName:\t\t%s %s\nNick:\t\t%s\nMailTo:\t\t%s\nId:\t\t%d\n",
			s.FirstName, s.Name, s.NickName, s.Email, s.Id)
	}
}

func (fst *Student) Equals(scd *Student) bool {
	if scd == nil {
		return false
	}
	return fst.Name == scd.Name && fst.FirstName == scd.FirstName && fst.Id == scd.Id
}

// JoinTeam adds student to team.
// Expects the team name (string).
// Returns an error if the team doesn't exist otherwise nil (succesful operation).
func (s Student) JoinTeam(teamName string) error {
	team, err := GetTeam(teamName)

	if err != nil {
		log.Fatal(err)
	}

	s.Team = team
	s.UpdateStudent()
	e := team.AddStudent(&s)

	return e
}
