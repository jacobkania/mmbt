package auth

import (
	"mmbt/models"
	"mmbt/testhelpers"
	"testing"
)

func TestCreateToken(t *testing.T) {
	testDB := testhelpers.CreateTestDB()

	user := &models.User{
		FullName: "Test McTesterson",
		Username: "t_mctest",
		Passw:    "abc",
	}

	_, err := testDB.Model(user).Insert()
	if err != nil {
		t.Fatal(err.Error())
	}

	// tokenService call
	tokenService := &TokenService{DB: testDB}
	resultToken, err := tokenService.CreateToken(user)

	if err != nil {
		t.Fatal(err.Error())
	}

	if resultToken.ID == 0 {
		t.Fatal("resultToken was not set")
	}

	fetchedToken := &models.UserLoginToken{}
	err = testDB.Model(fetchedToken).Last()

	if err != nil {
		t.Fatal(err.Error())
	}

	if fetchedToken.ID == 0 {
		t.Fatal("fetchedToken was not set")
	}
}
