package db

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func init() {
	NewTestDatabase()
}

func TestCreatePress(t *testing.T) {

	p1 := Player.Create().CreateRelations().BuildInsert()
	p2 := Player.Create().CreateRelations().BuildInsert()

	recipients := []int{p1.ID, p2.ID}

	p := Press.Create().Build()
	pressId, err := PressRepo.CreatePress(p, recipients)

	if err != nil {
		t.Fatalf("Could not create Press: %s", err.Error())
	}

	newPress, err := PressRepo.FindPress(pressId)
	if err != nil {
		t.Fatalf("Could not find Press (%d): %s", pressId, err.Error())
	}

	if newPress.Subject != p.Subject {
		t.Fatalf("Subject of Press created does not match given Subject: got (%s), want (%s)", newPress.Subject, p.Subject)
	}

	want := strings.Join([]string{p1.User.Username, p2.User.Username}, ",")
	if newPress.Recipients != want {
		t.Fatalf("Recipients usernames list does not match: got (%s), want (%s)", newPress.Recipients, want)
	}

	t.Logf("Created Press with Subject (%s), ID (%d)", newPress.Subject, newPress.ID)
}

func TestFindRecipients(t *testing.T) {

	g := Game.Create().BuildInsert()

	p1 := Player.Create().CreateUser().SetGame(&g).BuildInsert()
	p2 := Player.Create().CreateUser().SetGame(&g).BuildInsert()

	recipients := []int{p1.ID, p2.ID}

	p := Press.Create().Build()
	pressId, err := PressRepo.CreatePress(p, recipients)

	if err != nil {
		t.Fatalf("Could not create Press: %s", err.Error())
	}

	newRecipients, err := PressRepo.FindRecipients(pressId)
	if err != nil {
		t.Fatalf("Could not find Press Recipients (%d): %s", pressId, err.Error())
	}

	if len(newRecipients) == 0 {
		t.Fatalf("Recipients list is empty. Want (%v)", recipients)
	}

	for _, r := range newRecipients {
		if !slices.Contains(recipients, r.PlayerID) {
			t.Fatalf("Recipients list does not contain specified recipient. Got (%d), want (%v)", r.PlayerID, recipients)
		}
	}
	t.Logf("Found Press Recipients (Press ID: %d)", pressId)
}
