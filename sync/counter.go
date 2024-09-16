package sync

import "sync"

type Counter struct {
	/*
	   Here we could say:

	   sync.Mutex
	   value int

	   Meaning that the mutex is embedded in the struct, so you can do:

	   c.Lock()

	   in the methods (with a Counter receiver c)

	   This is inadvisable (according to LGWT) as it makes the mutex part of the public
	   interface of the struct, which you might not want
	*/
	mut   sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// Lock will make any other incoming goroutines have to wait
	c.mut.Lock()
	defer c.mut.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
