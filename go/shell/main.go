package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"os/exec"
	"errors"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		ps1 := getPS1()
		fmt.Printf("%v", ps1)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execCommand(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func getPS1() string {
		hostname, err := os.Hostname()

		if err != nil {
			hostname = ""
		}

		pwd, err := os.Getwd()

		if err != nil {
			pwd = ""
		}

		iprompt := os.Getenv("IPROMPT")

		if iprompt == "" {
			return fmt.Sprintf("(%v) %v > ", hostname, pwd)
		} else {
			prompt := buildPrompt(iprompt)
			return fmt.Sprintf(prompt)
		}

}

func buildPrompt(str string) string {
	str = strings.Replace(str, "%H", getHostname(), -1)
	str = strings.Replace(str, "%P", getPWD(), -1)

	return str
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}
	return hostname
}

func getPWD() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

func execCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
		case "cd":
			if len(args) < 2 {
				return errors.New("path required")
			}

			return os.Chdir(args[1])
		case "exit":
			os.Exit(0)
	}

	if strings.Contains(args[0], "=") {
		splits := strings.Split(input, "=")
		key := splits[0]
		value := strings.Join(splits[1:], "")

		return os.Setenv(key, value)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
