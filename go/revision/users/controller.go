package users

import (
	"fmt"

	utils "vrouilhac/revision/utils"
)

type NoUserError struct {
	message string
}

func (err NoUserError) Error() string {
	return fmt.Sprintf("Err[User]: %v", err.message)
}

func Signup(firstname, lastname, email, password string, age int32) User {
	return User {
		firstname,
		lastname,
		email,
		password,
		age,
	}
}

func BuildUser() User {
	email := utils.AskInput("Email")
	password := utils.AskInput("Password")
	return Signup("", "", email, password, 0)
}

func Login(users []User, email, password string) (User, error) {
	for _, user := range users {
		if user.email == email && user.password == password	{
			return user, nil
		}
	}

	return User{}, NoUserError{ "No user found with those credentials" }
}
