package library

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Window struct {
	head                      *Node
	tail                      *Node
	windowCount               int64
	allowedLimit              int64
	windowSlot                int16
	slotDuration              time.Duration
	lastSlotCreationTimeStamp time.Time
	wmux                      sync.Mutex
}

type Node struct {
	nextNode *Node
	prevNode *Node
	counter  int64
}

func createList(node *Node, listSize int16, tailNode **Node) *Node {
	if listSize <= 0 {
		return nil
	}

	var newNode = &Node{
		prevNode: node,
		counter:  0,
	}
	if listSize > 1 {
		newNode.nextNode = createList(newNode, listSize-1, tailNode)
	}
	if listSize == 1 {
		fmt.Println("Tila node is set")
		*tailNode = newNode
	}

	return newNode
}

func CreateWindow(numberOfSlots int16, slotDuration time.Duration) *Window {
	window := Window{
		windowCount:               0,
		windowSlot:                numberOfSlots,
		slotDuration:              slotDuration,
		lastSlotCreationTimeStamp: time.Now(),
	}

	var rootNode = &Node{
		prevNode: nil,
		counter:  0,
	}

	rootNode.nextNode = createList(rootNode, numberOfSlots-1, &window.tail)

	window.head = rootNode

	return &window
}

func addNode(window *Window) *Window {
	window.print()
	window.windowCount = window.windowCount + window.head.counter - window.tail.counter
	window.tail.prevNode.nextNode = nil
	window.tail = window.tail.prevNode

	var newNode = &Node{
		prevNode: nil,
		nextNode: window.head,
		counter:  0,
	}
	window.head.prevNode = newNode
	window.head = newNode

	return window
}

func GetCurrentLimit(window *Window) int64 {
	return window.head.counter + window.windowCount
}

func incr(node *Node) int64 {
	node.counter++
	val := node.counter
	return val
}

func (w *Window) isLimitExceeded() bool {
	result := false
	w.wmux.Lock()
	currentTime := time.Now()
	duration := currentTime.Sub(w.lastSlotCreationTimeStamp)
	if duration > w.slotDuration {
		nodesToAdd := duration / w.slotDuration
		if duration%w.slotDuration > 0 {
			nodesToAdd++
		}
		log.Println("nodes to add", int64(nodesToAdd))
		//TODO: Scope for performance improvement here. We can even  add multiple nodes in a single shot...
		for i := 0; i <= int(nodesToAdd); i++ {
			addNode(w)
		}
	}
	w.lastSlotCreationTimeStamp = currentTime
	if GetCurrentLimit(w) > w.allowedLimit {
		result = false
	}
	incr(w.head)
	result = true
	w.wmux.Unlock()

	return result
}

func (w *Window) print() {
	var windowProperties string
	windowProperties = fmt.Sprintf("Window Count %d, windowSlot %d, Slot Duration %d, Allowed Limit %d ", w.windowCount,
		w.windowSlot, w.slotDuration, w.allowedLimit)
	log.Println(windowProperties)
	log.Println("Node details")
	for i := w.head; i != nil; i = i.nextNode {
		if i == w.head {
			log.Println("HEAD NODE", i.counter)
		} else if i == w.tail {
			log.Println("TAIL NODE", i.counter)
		} else {
			log.Println("SUB NODE", i.counter)
		}

	}

}
