package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Lock() {
	c.mu.Lock()
}
func (c *Counter) Unlock() {
	c.mu.Unlock()
}
func (c *Counter) Value() int {
	return c.value
}
func NewCounter() *Counter {
	return &Counter{}
}
