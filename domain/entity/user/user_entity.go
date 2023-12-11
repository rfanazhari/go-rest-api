package user_entity

import (
	"fmt"
	"github.com/pkg/errors"
	"rest-api/pkg/common"
)

// generate getter automaticaly
type User struct {
	id        common.ID
	firstName string
	lastName  string
	email     string
}

type UserDto struct {
	Id        common.ID
	FirstName string
	LastName  string
	Email     string
}

func NewUser(dto *UserDto) (*User, error) {
	user := &User{
		id:        dto.Id,
		firstName: dto.FirstName,
		lastName:  dto.LastName,
		email:     dto.Email,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.firstName == "" {
		return errors.New("first name required")
	}

	if u.lastName == "" {
		return errors.New("last name required")
	}

	if u.email == "" {
		return errors.New("email required")
	}

	return nil
}

func (u *User) Id() common.ID {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.firstName, u.lastName)
}

func (u *User) Email() string {
	return u.email
}
