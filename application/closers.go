package application

import (
	"github.com/Rasikrr/learning_platform_core/interfaces"
	"sync"
)

type Closers struct {
	mut     sync.Mutex
	closers []interfaces.Closer
}

func NewClosers() *Closers {
	return &Closers{
		closers: make([]interfaces.Closer, 0),
		mut:     sync.Mutex{},
	}
}

func (c *Closers) Add(closer interfaces.Closer) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.closers = append(c.closers, closer)
}
