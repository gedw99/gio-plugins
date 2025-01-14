package internal

import (
	"sync"
)

// Scheduler is a runs a set of functions in the given an
// single runner function.
//
// That is useful for running a set of functions in a single
// goroutine, linked to the main loop.
type Scheduler struct {
	runner func(f func())

	mutex   sync.Mutex
	counter int
	tasks   map[int]int
	queue   []func()
	update  chan struct{}
}

// NewScheduler creates a new Scheduler.
func NewScheduler(runner func(f func())) (s *Scheduler) {
	s = &Scheduler{}
	s.SetRunner(runner)
	return s
}

// SetRunner sets the runner function.
func (s *Scheduler) SetRunner(r func(f func())) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if r == nil {
		return
	}

	if s.runner == nil {
		if s.update == nil {
			s.update = make(chan struct{}, 1)
		}

		go func() {
			fns := make([]func(), 0, 32)
			last := 0
			for range s.update {
			repeat:
				s.mutex.Lock()
				if last == s.counter {
					s.mutex.Unlock()
					continue
				}
				last = s.counter

				for i := range s.tasks {
					s.tasks[i] = -1
				}
				for i := range s.queue {
					fns = append(fns, s.queue[i])
				}
				s.queue = s.queue[:0]

				runner := s.runner
				s.mutex.Unlock()

				for i := range fns {
					runner(fns[i])
				}

				fns = fns[:0]

				s.mutex.Lock()
				isOutdated := s.counter != last && len(s.update) == 0
				s.mutex.Unlock()
				if isOutdated {
					goto repeat
				}
			}
		}()
	}

	s.runner = r
	s.signal()
}

// Run runs a function in the scheduler, however it only
// allow one function per method.
//
// That will replace the old declared function if the
// method is already in use.
func (s *Scheduler) Run(method int, f func()) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.tasks == nil {
		s.tasks = make(map[int]int, 32)
	}

	i, ok := s.tasks[method]
	if i < 0 || !ok {
		s.tasks[method] = len(s.queue)
		s.queue = append(s.queue, f)
	} else {
		s.queue[i] = f
	}

	s.counter++

	s.signal()
}

// MustRun runs a function in the scheduler, that is
// guaranteed to run and duplication may happen if
// misused.
func (s *Scheduler) MustRun(f func()) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.counter++
	s.queue = append(s.queue, f)

	s.signal()
}

func (s *Scheduler) signal() {
	select {
	case s.update <- struct{}{}:
	default:
	}
}
