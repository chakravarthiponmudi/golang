package library

import (
	"log"
	"testing"
	"time"
)

func TestSider(t *testing.T) {
	t.Log("Testing")
	testWindow := CreateWindow(10, 1*time.Second)
	t.Log(testWindow)
	hitTesting(testWindow)
	testWindow.print()
	hitTesting(testWindow)
	testWindow.print()
	hitTesting(testWindow)
	hitTesting(testWindow)
	hitTesting(testWindow)
	hitTesting(testWindow)
	time.Sleep(5 * time.Second)
	hitTesting(testWindow)
	hitTesting(testWindow)
	hitTesting(testWindow)
	hitTesting(testWindow)
	hitTesting(testWindow)
	hitTesting(testWindow)
	hitTesting(testWindow)
	hitTesting(testWindow)
	testWindow.print()
}

func hitTesting(w *Window) bool {
	log.Println("Hitinng API")
	return w.isLimitExceeded()
}
