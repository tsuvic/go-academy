package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *Counter) Increment(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := Counter{v: make(map[string]int)}
	for i := 0; i < 10000; i++ {
		go c.Increment("somekey")
		fmt.Println(c.Value("somekey"))
	}

	// time.Sleep(time.Second)

}
