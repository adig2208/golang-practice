package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("invalid input")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
