package main

//view "github.com/eulersexception/glabs-ui/view"

import (
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/badger/v3"
)

func main() {

	// s1, _ := model.NewStudent(nil, "Payne", "Max", "MaxPayne", "max@payne.com", 1)
	// s2, _ := model.NewStudent(nil, "Payne", "Martin", "MartinPayne", "martin@payne.com", 2)
	// s3, _ := model.NewStudent(nil, "Payne", "Michi", "MichiPayne", "michi@payne.com", 3)
	// s4, _ := model.NewStudent(nil, "Duck", "Tick", "TickDuck", "tick@duck.com", 4)
	// s5, _ := model.NewStudent(nil, "Duck", "Trick", "TrickDuck", "trick@duck.com", 5)
	// s6, _ := model.NewStudent(nil, "Duck", "Track", "TrackDuck", "track@duck.com", 6)
	// // s7 := model.NewStudent(nil, "Duck", "Donald", "DonaldDuck", "track@duck.com", 7)

	// s1.PrintData()
	// s2.PrintData()
	// s3.PrintData()
	// s4.PrintData()
	// s5.PrintData()
	// s6.PrintData()

	// model.DeleteStudent(1)

	// s, errS := model.GetStudent(2)

	// if errS != nil {
	// 	log.Fatal(errS)
	// } else {
	// 	s.PrintData()
	// }
	// s3_copy := model.GetStudent(1)
	// s3_copy.PrintData()

	// fmt.Printf("Current unix time = %v", time.Now().Unix())

	// time.Sleep(60 * time.Second)
	// fmt.Printf("Current unix time = %v", time.Now().Unix())

	// zeros := model.NewTeam(nil, "zeros")
	// ones := model.NewTeam(nil, "ones")

	// zeros.PrintMembers()
	// ones.PrintMembers()

	// fmt.Printf("First set of team in DB\n")
	// zeros.SetTeam()
	// ones.SetTeam()

	// fmt.Printf("First get of teams from DB\n")
	// zerros := model.GetTeam("zeros")
	// onnes := model.GetTeam("ones")

	// fmt.Printf("First output of teams (teams should be empty):\n")
	// zerros.PrintMembers()
	// onnes.PrintMembers()
	// zerros.AddStudentToTeam(s1).AddStudentToTeam(s2).AddStudentToTeam(s3)
	// onnes.AddStudentToTeam(s4).AddStudentToTeam(s5).AddStudentToTeam(s6)
	// zerros.AddStudentToTeam(s7)

	// zerros.RemoveStudentFromTeam(*s2)
	// zerros.RemoveStudentFromTeam(*s3)
	// zerros.AddStudentToTeam(s1)
	// zerros.AddStudentToTeam(s7)
	// zerros.AddStudentToTeam(s7)
	// zerros.AddStudentToTeam(s7)
	// zerros.AddStudentToTeam(s7)
	// zerros.AddStudentToTeam(s7)

	//for _, v := range zerros.Students {
	//	zerros.RemoveStudentFromTeam(*v)
	//}

	// fmt.Printf("First output of teams:\n")
	// zerros.PrintMembers()
	// onnes.PrintMembers()
	//s2.PrintData()

	//fmt.Println("Iterator section")
	//model.GetStudent(2)

	// // fmt.Printf("Data manipulation - student deletion - output of teams:\n")
	// // zerros = *zerros.RemoveStudentFromTeam(*s1)
	// // onnes = *onnes.AddStudentToTeam(s1)

	// //zerros.SetTeam()
	// //onnes.SetTeam()
	// //zerros.PrintMembers()
	// fmt.Printf("Get teams from DB after manipulation:\n")
	// model.GetTeam(zerros.Name).PrintMembers()
	// model.GetTeam(onnes.Name).PrintMembers()

	// members := onnes.Students

	// for _, v := range members {
	// 	v.PrintData()
	// }

	// starter := &model.StarterCode{
	// 	Url:             "git@gitlab.lrz.de:vss/startercode/startercodeB1.git",
	// 	FromBranch:      "ws20.1",
	// 	ProtectToBranch: true,
	// }

	// clone := &model.Clone{
	// 	LocalPath: "/Users/obraun/lectures/vss/labs/gitlab/semester/ob-20ws/blatt1",
	// 	Branch:    "develop",
	// }

	// a := model.NewAssignment("Test Ass", nil, clone, starter)

	// a.AddTeamToAssignment(&onnes)
	// a.AddTeamToAssignment(&zerros)

	// a.SetAssignment()

	// ass := model.GetAssignment(a.Name)

	// fmt.Print("First output after seeting Assignment:\n")
	// ass.PrintData()

	// fmt.Print("Assignment after local deletion of Team from Assignment:\n")
	// ass.DeleteTeamFromAssignment(onnes)
	// ass.PrintData()

	// ass.SetAssignment()
	// ass = model.GetAssignment(ass.Name)

	// fmt.Print("Second output after seeting updated Assignment in DB:\n")
	// ass.PrintData()

	// err := model.DeleteAssignment(ass.Name)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// n := ass.Name
	// ass = model.GetAssignment(n)

	// if &ass != nil {
	// 	fmt.Printf("Assignment %s deleted from DB\n", n)
	// }

	//stud := model.GetStudent(12345)

	//fmt.Printf("%s %s, Nick: %s, EMail: %s, Matrikelnr: %d, Team: %s\n", stud.FirstName, stud.Name, stud.NickName, stud.Email, stud.Id, stud.Team.Name)

	//err := model.DeleteStudent(12345)

	//if err != nil {
	//	log.Fatal(err)
	//}

	//stud = model.GetStudent(12345)

	//fmt.Printf("%s %s, Nick: %s, EMail: %s, Matrikelnr: %d\n", stud.FirstName, stud.Name, stud.NickName, stud.Email, stud.Id)

	//view.NewHomeview().Window.ShowAndRun()

	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	go cleanValues(db)

	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		log.Println("\n------------ created Iterator")
		defer it.Close()
		log.Println("\n------------ right before loop")
		for it.Rewind(); it.Valid(); it.Next() {
			log.Println("\n------------ creating item")
			item := it.Item()
			log.Println("\n------------ creating key")
			k := item.Key()

			log.Println("\n------------ getting value from item Iterator")
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%d, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal()
	}

	time.Sleep(2 * time.Minute)

	return

}

func cleanValues(db *badger.DB) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
	again:
		err := db.RunValueLogGC(0.7)
		if err == nil {
			goto again
		}
	}
}
