package library

import (
	"sync"
	"testing"
	"time"
)

type counter struct {
	allowed int
	blocked int
	mux     sync.Mutex
}

var result counter

var wg sync.WaitGroup

func TestSlider(t *testing.T) {
	result.allowed = 0
	result.blocked = 0

	testsize := 100000000

	wg.Add(testsize)
	testWindow := CreateWindow(1000, 1*time.Millisecond, 500000)
	for i := 0; i < testsize; i++ {
		go hitTesting(testWindow, t, i)
		if i%100000 == 0 {
			time.Sleep(1 * time.Millisecond)
		}

	}
	wg.Wait()
	testWindow.print()
	t.Log("allowed", result.allowed)
	t.Log("blocked", result.blocked)
	t.Log("Total", result.allowed+result.blocked)
}

func hitTesting(w *Window, t *testing.T, i int) {
	if w.isLimitExceeded() {
		result.mux.Lock()
		result.allowed++
		result.mux.Unlock()
	} else {
		result.mux.Lock()
		result.blocked++
		result.mux.Unlock()

	}

	wg.Done()
}
