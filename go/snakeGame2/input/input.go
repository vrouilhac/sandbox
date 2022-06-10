package input

import (
	"os/exec"
	"os"
)

func GetInput(c chan string) {
	// source: https://stackoverflow.com/questions/54422309/how-to-catch-keypress-without-enter-in-golang-loop
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)

	for {
		os.Stdin.Read(b)	
		c <- string(b)
	}
}
