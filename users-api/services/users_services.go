package services

import (
	"github.com/bookstore/users-api/domain/users"
	"github.com/bookstore/users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
