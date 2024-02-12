package db

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func init() {
	NewTestDatabase()
}

func TestFindAllPress(t *testing.T) {
	assert := assert.New(t)

	recipients := []int{PlayerFactory.CreateRelations().BuildInsert().ID, PlayerFactory.CreateRelations().BuildInsert().ID}
	p := PlayerFactory.CreateRelations().BuildInsert()
	press1 := PressFactory.Create().SetPlayer(&p).BuildInsert(recipients)
	press2 := PressFactory.Create().SetPlayer(&p).BuildInsert(recipients)

	allPress, err := PressRepo.FindAllPress(p.ID)
	if err != nil {
		assert.FailNow(err.Error())
	}

	assert.Equal(len(allPress), 2, "Incorrect amount of press found")
	assert.Equal(allPress[0].GetSubject(), press1.GetSubject(), "Press 1 has the wrong Subject")
	assert.Equal(allPress[1].GetSubject(), press2.GetSubject(), "Press 2 has the wrong Subject")
}

func TestCreatePress(t *testing.T) {

	p1 := PlayerFactory.Create().CreateRelations().BuildInsert()
	p2 := PlayerFactory.Create().CreateRelations().BuildInsert()

	recipients := []int{p1.ID, p2.ID}

	p := PressFactory.Create().Build()
	pressId, err := PressRepo.CreatePress(p, recipients)

	if err != nil {
		t.Fatalf("Could not create Press: %s", err.Error())
	}

	newPress, err := PressRepo.FindPress(pressId)
	if err != nil {
		t.Fatalf("Could not find Press (%d): %s", pressId, err.Error())
	}

	if newPress.GetSubject() != p.GetSubject() {
		t.Fatalf("Subject of Press created does not match given Subject: got (%s), want (%s)", newPress.GetSubject(), p.GetSubject())
	}

	want := strings.Join([]string{p1.User.Username, p2.User.Username}, ",")
	if newPress.GetRecipients() != want {
		t.Fatalf("Recipients usernames list does not match: got (%s), want (%s)", newPress.GetRecipients(), want)
	}
}

func TestFindRecipients(t *testing.T) {

	g := Game.Create().BuildInsert()

	p1 := PlayerFactory.Create().CreateUser().SetGame(&g).BuildInsert()
	p2 := PlayerFactory.Create().CreateUser().SetGame(&g).BuildInsert()

	recipients := []int{p1.ID, p2.ID}

	p := PressFactory.Create().Build()
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
}
