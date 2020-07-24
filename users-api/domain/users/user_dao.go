//This file is an access layer between application and a database.
package users

import (
	"fmt"

	"github.com/bookstore/users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

//Get returns pointer to User, when he exists
//or appropriate error when user cannot be found.
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	usersDB[user.ID] = user
	return nil
}
