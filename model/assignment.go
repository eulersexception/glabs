package model

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"
)

type StarterCode struct {
	Url             string
	FromBranch      string
	ProtectToBranch bool
}

func (s StarterCode) toString() string {
	return fmt.Sprintf("\tStarterCode:\n\t\tUrl:\t%s\n\t\tFromBranch:\t%s\n\t\tProtectToBranch:\t%v", s.Url, s.FromBranch, s.ProtectToBranch)
}

type Clone struct {
	LocalPath string
	Branch    string
}

func (c Clone) toString() string {
	return fmt.Sprintf("\tClone:\n\t\tLocalPath:\t%s\n\t\tBranch:\t%s", c.LocalPath, c.Branch)
}

type Assignment struct {
	Name              string
	Semester          *Semester
	Teams             []*Team
	LocalClone        *Clone
	Starter           *StarterCode
	ContainerRegistry bool
}

func NewAssignment(name string, sem *Semester, clone *Clone, starter *StarterCode) *Assignment {
	if name == "" {
		fmt.Println("Please enter a valid course name.")
		return nil
	}

	var teams []*Team

	assignment := &Assignment{
		Semester:   sem,
		Name:       name,
		Starter:    starter,
		LocalClone: clone,
		Teams:      teams,
	}

	return assignment
}

func (a *Assignment) AddTeamToAssignment(t *Team) *Assignment {
	if t == nil {
		fmt.Println("No valid data for team.")
		return a
	}

	a.Teams = append(a.Teams, t)
	

	return a
}

func (a *Assignment) DeleteTeamFromAssignment(t Team) *Assignment {
	index := -1

	for i, v := range a.Teams {
		if v.Name == t.Name {
			index = i
		}
	}

	if index == -1 {
		return a
	}

	a.Teams[index] = a.Teams[len(a.Teams)-1]
	a.Teams = a.Teams[:len(a.Teams)-1]

	return a
}

func (a Assignment) encodeAssignment() []byte {
	data, err := json.Marshal(a)

	if err != nil {
		panic(err)
	}

	return data
}

func decodeAssignment(data []byte) Assignment {
	var a Assignment
	err := json.Unmarshal(data, &a)

	if err != nil {
		panic(err)
	}

	return a
}

func (a Assignment) SetAssignment() {
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(a.Name), []byte(a.encodeAssignment()))
		err = txn.SetEntry(e)

		return err
	})
}

func GetAssignment(name string) Assignment {
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var a Assignment

	err = db.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(name))

		if err != nil {
			log.Fatal(err)
		}

		err = item.Value(func(val []byte) error {
			a = decodeAssignment(val)
			//fmt.Println(fmt.Sprintf("Key = %s, Value = %s", item.String(), string(val)))
			return err
		})
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	return a
}

func DeleteAssignment(name string) error {
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

func (a Assignment) PrintData() {
	fmt.Printf("-------------------\nName:\t\t%s\nLocalClone:\t%s\nStarter:\t\t%s\nContainerRegistry:\t\t%v\n",
		a.Name, a.LocalClone.toString(), a.Starter.toString(), a.ContainerRegistry)

	for _, v := range a.Teams {
		v.PrintMembers()
	}
}
