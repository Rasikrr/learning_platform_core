package application

import (
	"github.com/Rasikrr/learning_platform_core/interfaces"
	"sync"
)

type Starters struct {
	mut      sync.Mutex
	starters []interfaces.Starter
}

func NewStarters() *Starters {
	return &Starters{
		starters: make([]interfaces.Starter, 0),
		mut:      sync.Mutex{},
	}
}

func (s *Starters) Add(starter interfaces.Starter) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.starters = append(s.starters, starter)
}
