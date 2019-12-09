package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//This is a Unit test in go
func TestGetUserNotUserFound(t *testing.T) {
	//Inicialization

	//Execution
	user, err := GetUser(0)


	//Validation
	assert.Nil(t, user, "we are not expecting a user with id 0")
	/*if(user != nil) {
		t.Error("we are not expecting a user with id 0")
	}*/
	assert.NotNil(t, err, "we were expecting an error when user id is 0")
	/*if err == nil{
		t.Error("we were expecting an error when user id is 0")
	}*/
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	/*if err.StatusCode != http.StatusNotFound {
		t.Error("we are expecting 404 when user is n ot found")
	}*/
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Lio", user.FirstName)
	assert.EqualValues(t, "Mazza", user.LastName)
	assert.EqualValues(t, "myemail@gmail.com", user.Email)
}
