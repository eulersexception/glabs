package model

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"
	"github.com/google/go-cmp/cmp"
)

type Team struct {
	Assignment *Assignment
	Name       string
	Students   []*Student
}

func NewTeam(assignment *Assignment, name string) *Team {
	if name == "" {
		fmt.Println("Please enter a valid team name.")
		return nil
	}

	var students []*Student

	team := &Team{
		Assignment: assignment,
		Name:       name,
		Students:   students,
	}

	team.SetTeam()

	return team
}

func (t *Team) AddStudentToTeam(s *Student) *Team {
	index := -1

	for i, v := range t.Students {
		if v.Id == s.Id {
			if cmp.Equal(v, s) {
				return t
			}
		} else {
			index = i
		}
	}

	if index == -1 {
		t.Students = append(t.Students, s)
		err := DeleteTeam(t.Name)

		if err != nil {
			panic(err)
		}

		t.SetTeam()
		s.AddToTeam(t)
	}

	return t
}

func (t *Team) RemoveStudentFromTeam(s Student) *Team {
	index := -1

	for i, v := range t.Students {
		if s.Id == v.Id {
			index = i
		}
	}

	if index == -1 {
		return t
	}

	t.Students[index] = t.Students[len(t.Students)-1]
	t.Students = t.Students[:len(t.Students)-1]
	s.Team = nil
	s.SetStudent()
	t.SetTeam()

	return t
}

func (t Team) encodeTeam() []byte {
	data, err := json.Marshal(t)

	if err != nil {
		panic(err)
	}

	return data
}

func decodeTeam(data []byte) Team {
	var t Team
	err := json.Unmarshal(data, &t)

	if err != nil {
		panic(err)
	}

	return t
}

func (t Team) SetTeam() {
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(t.Name), []byte(t.encodeTeam()))
		err = txn.SetEntry(e)

		return err
	})
}

func GetTeam(name string) Team {
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var t Team

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(name))

		if err != nil {
			log.Fatal(err)
		}

		err = item.Value(func(val []byte) error {
			t = decodeTeam(val)
			//fmt.Println(fmt.Sprintf("Key = %s, Value = %s", item.String(), string(val)))
			return err
		})
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	return t
}

func DeleteTeam(name string) error {
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		e := txn.Delete([]byte(name))

		return e
	})

	return err
}

func (t Team) PrintMembers() {
	if t.Assignment != nil {
		fmt.Printf("Current Assignment = %s", t.Assignment.Name)
	}

	fmt.Printf("Team: %s\n", t.Name)

	for _, v := range t.Students {
		fmt.Printf("Member: %s\n", v.NickName)
	}
}
