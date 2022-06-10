package Timer

import (
	"time"
	"fmt"
)

func ExecuteAtInterval(fn func(func() ()) ()) {
	ticker := time.NewTicker(120 * time.Millisecond)

	done := make(chan bool) // learn channels

	go func () { // learn goroutine
		for {
			select { // learn select
				case <-done: // learn wtf symbole
				return
			case <-ticker.C:
				fn(func() {
					ticker.Stop()
					fmt.Println("END")
					done <- true // learn wtf symbole part 2
				})
			}
		}
	}()

	time.Sleep(1000*time.Minute)
}
