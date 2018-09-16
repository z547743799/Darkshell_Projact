package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)

	// 写的时候啥也不能干
	go write(1)
	go read(2)
	go write(3)

	time.Sleep(2 * time.Second)
}

func read(i int) {
	println(i, "read start")

	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "read over")
}

func write(i int) {
	println(i, "write start")

	m.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	m.Unlock()

	println(i, "write over")
}
func println(i int, s string) {
	i++
	fmt.Println(i, s)
}
