package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
/*
show dbs
show tables
show table <table>
add db <name>
add table <name> [(name:string;age:int)]
insert table name ("Vincent";23;4)
insert table name (age: 32, name: "Vincent", run: 4)
*/

type Database[T any] struct {
	tables []T
}


func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		line, _, err := stdinReader.ReadLine()

		if err != nil {
			fmt.Println("Error reading line")
		}

		if end := Execute(string(line)); end {
			break
		}
	}
}

func Execute(str string) (end bool) {
	if str == "exit" {
		fmt.Println("Bye")
		return true
	}

	token, err := Tokenize(str)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(token)

	return false
}

type Token struct {
	action string
	verb string
}

type InvalidActionError string

func (err InvalidActionError) Error() string {
	return fmt.Sprintf("An error happened")
}

func Tokenize(str string) (token Token, err error) {
	strs := strings.Split(str, " ")

	action, isAction := ValidateAction(strs[0])

	if isAction == false {
		return Token{}, InvalidActionError("Action is not a valid action")
	}

	if len(strs) < 2 {
		return Token{}, InvalidActionError("Missing arguments")
	}

	verb, isVerb := ValidateVerb(strs[1])

	if isVerb == false {
		return Token{}, InvalidActionError("Verb is not a valid verb")
	}

	return Token {
		action,
		verb,
	}, nil
}

func ValidateAction(str string) (action string, isAction bool) {
	validActions := []string{"add", "insert", "show"}
	isAction = false

	for _, v := range validActions {
		if str == v {
			isAction = true
		}
	}
	return str, isAction
}

func ValidateVerb(str string) (verb string, isVerb bool) {
	validVerbs := []string{"dbs", "db", "table", "tables"}
	isVerb = false

	for _, v := range validVerbs {
		if str == v {
			isVerb = true
		}
	}
	return str, isVerb
}
