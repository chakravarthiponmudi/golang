package main

import (
	"time"
)

func say(s string) {
	for i := 0; i < 50; i++ {
		time.Sleep(10 * time.Millisecond)
		// fmt.Print(s)
	}
}

func main() {
	for i := 0; i < 10000000; i++ {
		go say("world " + string(i))
	}
	// go say("world")
	// go say("world1")
	// go say("world2")
	// go say("world3")
	// go say("world4")
	// go say("world5")
	// go say("world6")
	// go say("world7")
	// go say("world7")
	// go say("world8")
	// go say("world9")
	say("completed")
}
