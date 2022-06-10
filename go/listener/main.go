package main

import (
	"fmt"
	"os"
	"time"
	"os/exec"
	"bufio"
	"log"
)

const (
	LEFT = 10
	TOP = 10
	RIGHT = 10
	BOTTOM = 10
)

func main() {
	Clear()
	cha := make(chan string)


	go func (ch chan string) {
		// disable input buffering and do not display characters on screen
		// source: https://stackoverflow.com/questions/54422309/how-to-catch-keypress-without-enter-in-golang-loop
		exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
		exec.Command("stty", "-f", "/dev/tty", "-echo").Run()

		var b []byte = make([]byte, 1)

		for {
		os.Stdin.Read(b)	
		ch <- string(b)
		}
	}(cha)

	var direction string = "UP"
	count := 0

	for {
		Clear()

		count++
		fmt.Println(direction, count)

		select {
			case stdin, _ := <-cha:
				direction = stdin
			default:
				fmt.Printf("")
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func Clear() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
}

func Get(x *byte, c chan byte) {
	t(func() {
		fmt.Println("there", *x)
	})
}

func Sc(x *byte, c chan byte) {
	reader := bufio.NewReader(os.Stdin)

		p := make([]byte, 1)

		for {
			n, err := reader.Read(p)
			if err != nil {
				log.Fatal("ERROR")
			}

			*x = p[:n][0]
			c <- p[:n][0]
			fmt.Println(p[:n][0])
		}
}

func t(fn func() ()) {
	ticker := time.NewTicker(1000 * time.Millisecond)

	done := make(chan bool) // learn channels
 
	go func () { // learn goroutine
		for {
			select { // learn select
				case <-done: // learn wtf symbole
					return
				case <-ticker.C:
					fn()
			}
		}
	}()

	time.Sleep(100000*time.Minute)
	ticker.Stop()
	fmt.Println("END")
	done <- true // learn wtf symbole part 2
}
