//This file is an access layer between application and a database.
package users

import (
	"fmt"
	"strings"

	"github.com/bookstore/users-api/datasources/mysql/users_db"
	"github.com/bookstore/users-api/utils/date_utils"
	"github.com/bookstore/users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)

//Get returns pointer to User, when he exists
//or appropriate error when user cannot be found.
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s already exists", user.Email),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save an user: %s", err),
		)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err),
		)
	}
	user.ID = userID
	return nil
}
