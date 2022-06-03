package main

import (
	"fmt"
)

type User struct {
	username, email string
	age int
}

func main() {
	usersByCity := make(map[string][]User)

	usersByCity["Caen"] = []User{
			{"Vincent", "crys@gmail.com", 23},
			{"Julius", "julius@gmail.com", 41},
	}

	usersByCity["Paris"] = []User{
			{"Maris", "marie@gmail.com", 30},
			{"Mark", "mark@gmail.com", 12},
	}
	fmt.Println(usersByCity["Caen"])

	employees := CompanyEmployee()

	isFilled, employee := IsJobFilled(employees, "lead_developer")
	fmt.Println(isFilled, employee)

	isFilled2, employee2 := IsJobFilled(employees, "developer")
	
	fmt.Println(isFilled2, employee2)

	fmt.Println(employees)
}

func CompanyEmployee() map[string]User {
	employee := map[string]User {
		"lead_developer": {
			"Vincent", "vincent@gmail.com", 44,
		},
	}

	return employee
}

func IsJobFilled(employees map[string]User, job string) (bool, User) {
	elem, isFilled := employees[job]
	return isFilled, elem
}
