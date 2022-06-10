package users

type User struct {
	firstname, lastname, email string
	password string
	age int32
}

func (user *User) GetEmail() string {
	return user.email
}
