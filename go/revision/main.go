package main

import (
	"fmt"
	"os"

	"vrouilhac/revision/users"
	"vrouilhac/revision/utils"
)

type UserStatus(string)
type Menu(string)

type Todo struct {
	title string
	done bool
}


var USERS []users.User
var TODOS []Todo

type TUI struct {
	status UserStatus
	menu Menu
	user users.User
}

func (tui *TUI) PrintMenu() {
	fmt.Println("What do you want to do ?")
	if tui.status == "UNLOGGED" {
		fmt.Println("\t1 ) Register")
		fmt.Println("\t2 ) Login")
		tui.menu = Menu("REGISTER")
	} else {
		fmt.Println("\t1 ) List todos")
		fmt.Println("\t2 ) Logout")
		tui.menu = Menu("TODOS")
	}
}

func (tui *TUI) GetChoice() string {
	var choice string	
	fmt.Printf("> ")
	fmt.Scan(&choice)
	return choice
}

func (tui *TUI) Quit() {
	os.Exit(1)	
}

func (tui *TUI) Register(choice string) {
	var user users.User

	switch choice {
		case "1":
			user = users.BuildUser()
			USERS = append(USERS, user)
		case "2":
			email := utils.AskInput("Email")
			password := utils.AskInput("Password")
			loggedUser, err := users.Login(USERS, email, password)

			if err != nil {
				fmt.Println("Wrong Credentials")
				tui.Quit()
			} else {
				user = loggedUser
			}
		default:
			fmt.Println("Choice not recognized")
			return
	}

	tui.user = user
	tui.status = UserStatus("LOGGED")
}

func (tui *TUI) Logout() {
	var filteredUsers []users.User

	for _, user := range USERS {
		if user.GetEmail() != tui.user.GetEmail() {
			filteredUsers = append(filteredUsers, user)	
		}
	}

	tui.user = users.User{}
	tui.status = UserStatus("UNLOGGED")
	USERS = filteredUsers
}

func (tui *TUI) ListTodos() {
	for _, todo := range TODOS {
		fmt.Println(todo)
	}
}

func (tui *TUI) DoMainMenu(choice string) {
	switch choice {
		case "1":
			tui.ListTodos()
		case "2":
			tui.Logout()
		default:
			fmt.Println("Bad choice")
	}
}

func (tui *TUI) ExecuteChoice(choice string) {
	switch tui.menu {
		case "REGISTER":
			tui.Register(choice)
		case "TODOS":
			tui.DoMainMenu(choice)
		default:
			tui.Quit()
	}
}

func main() {
	tui := TUI{ status: UserStatus("UNLOGGED"), menu: Menu("REGISTER")}

	for {
		tui.PrintMenu()
		choice := tui.GetChoice()
		tui.ExecuteChoice(choice)
	}
}
