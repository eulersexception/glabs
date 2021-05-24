package model

import (
	"testing"

	"github.com/eulersexception/glabs-ui/util"
	"github.com/google/go-cmp/cmp"

	DB "modernc.org/ql"
)

var testStarter = &StarterCode{
	Url:             "www.test.com",
	FromBranch:      "develop",
	ProtectToBranch: true,
}

var testClone = &Clone{
	LocalPath: "local",
	Branch:    "master",
}

var testAssignment = &Assignment{}

var team = &Team{}

var studOne = &Student{
	Name:       "Minogue",
	FirstName:  "Kylie",
	NickName:   "kymi",
	Email:      "kymi@example.com",
	MatrikelNr: 10000,
}

var studTwo = &Student{
	Name:       "Simone",
	FirstName:  "Nina",
	NickName:   "nisi",
	Email:      "nisi@example.com",
	MatrikelNr: 10001,
}

func TestNewTeamSuccess(t *testing.T) {
	want := &Team{Name: "TestTeam1"}

	NewTeam("TestTeam1")

	got := GetTeam("TestTeam1")

	if !cmp.Equal(want, got) {
		t.Errorf("NewTeam:\nName = '%s', want '%s'\n",
			got.Name, want.Name)
	}
}

func TestNewTeamFail(t *testing.T) {
	want := "\n+++ Please enter a valid team name."

	gotTeam, got := NewTeam("")

	if gotTeam != nil || want != got {
		t.Errorf("Expected error message '%s', got a value %v", want, gotTeam)
	}
}

func TestDeleteTeam(t *testing.T) {
	want := 1

	NewTeam("TestTeam2")
	DeleteTeam("TestTeam2")

	db := util.GetDB()
	defer util.FlushAndClose(db)

	got, _, err := db.Run(DB.NewRWCtx(), `
			BEGIN TRANSACTION;
				SELECT count(*) FROM Team WHERE TeamName = $1;
			COMMIT;
	`, "TestTeam2")

	if err != nil {
		panic(err)
	}

	if len(got) != want {
		t.Errorf("Expected %d row with column headers, %d rows", want, len(got))
	}
}

func TestAddExistingStudent(t *testing.T) {

}
