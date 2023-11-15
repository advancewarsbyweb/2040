package db

import (
	"testing"
)

func init() {
	NewTestDatabase()
}

func TestCreateUser(t *testing.T) {
	u := User.Create().Build()

	userId, err := UserRepo.CreateUser(u)

	if err != nil {
		t.Fatalf("Could not create User: %s", err.Error())
	}

	newUser, err := UserRepo.FindUser(userId)
	if err != nil {
		t.Fatalf("Could not find User (%d): %s", userId, err.Error())
	}

	if newUser.ID != userId {
		t.Fatalf("ID of user created does not match given User. Got (%d), want (%d)", newUser.ID, userId)
	}

	t.Logf("Created User (%s)", u.Username)
}
