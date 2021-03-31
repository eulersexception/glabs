package model

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"
	"github.com/google/go-cmp/cmp"
)

// Team - Name is the primary key. All fields are public and
// Getter or Setter functions relate to database operations.
type Team struct {
	Assignment *Assignment
	Name       string
	Students   []*Student
}

// NewTeam creates a new team and stores the object in DB.
// String argument for name must not be empty.
// If a team with given name exists already in DB, the existing dataset will be overwritten.
// Returns a new teamo.
func NewTeam(assignment *Assignment, name string) (*Team, string) {

	if name == "" {
		res := "\n+++ Please enter a valid team name."
		return nil, res
	}

	var students []*Student

	team := &Team{
		Assignment: assignment,
		Name:       name,
		Students:   students,
	}

	err := team.setTeam()

	if err != nil {
		return nil, err.Error()
	}

	return team, ""
}

func (t *Team) AddStudent(s *Student) error {
	index := -1

	for i, v := range t.Students {
		if v.Id == s.Id {
			index = i
			if !cmp.Equal(v, s) {
				t.Students[i] = s
				err := t.UpdateTeam()

				if err != nil {
					log.Fatal(err)
					return err
				}

				s.Team = t
				err = s.UpdateStudent()

				if err != nil {
					log.Fatal(err)
					return err
				}
			}
		}
	}

	if index == -1 {
		t.Students = append(t.Students, s)
		err := t.UpdateTeam()

		if err != nil {
			log.Fatal(err)
			return err
		}

		s.Team = t
		err = s.UpdateStudent()

		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func (t *Team) RemoveStudent(s *Student) error {
	index := -1

	for i, v := range t.Students {
		if s.Id == v.Id {
			index = i
		}
	}

	if index == -1 {
		return nil
	}

	t.Students[index] = t.Students[len(t.Students)-1]
	t.Students = t.Students[:len(t.Students)-1]
	err := t.setTeam()

	if err != nil {
		log.Fatal(err)
		return err
	}

	s.Team = nil
	err = s.UpdateStudent()

	if err != nil {
		log.Fatal(err)
		return err
	}

	return err
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

// UpdateTeam changes a teams record in DB.
// Returns an error if the update fails.
func (t Team) UpdateTeam() error {
	_, err := GetTeam(t.Name)

	if err != nil {
		log.Printf("\n+++ Update of team with name %s failer while checking if team exists.\n+++ %s\n", t.Name, err.Error())
		return err
	}

	err = t.setTeam()

	return err
}

// This function updates team record in DB. An update could be a creation or edition of a record.
func (t Team) setTeam() error {
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	k := []byte(t.Name)
	v := []byte(t.encodeTeam())

	err = db.Update(func(txn *badger.Txn) error {
		e := txn.Set(k, v)

		return e
	})

	return err
}

// GetTeam fetches team from DB with an argument of type string as name.
// Returns an error if fetch fails or a pointer to the Team.
func GetTeam(name string) (*Team, error) {
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
		return nil, err
	}

	return &t, nil
}

// DeleteTeam removes a team by name (string) from DB.
// Returns an error if operation fails.
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

// PrintData outputs a human readable string for team data.
func (t Team) PrintMembers() {
	if t.Assignment != nil {
		fmt.Printf("Current Assignment = %s", t.Assignment.Name)
	}

	fmt.Printf("Team: %s\n", t.Name)

	for _, v := range t.Students {
		fmt.Printf("Member: %s\n", v.NickName)
	}
}
