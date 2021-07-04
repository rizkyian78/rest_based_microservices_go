package domain

import (
	"net/http"
	"rest_based_micro_go/src/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "njeng", LastName: "ok", Email: "okeeee@mail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{Message: "User Not Found", StatusCode: http.StatusBadRequest, Code: "not_found"}

}
