package patterns

import (
	"errors"
	"sync"
)

func AttemptWrite[T any](ch chan T, v T) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}

func AttemptRead[T any](ch chan T) (T, bool) {
	var v T
	select {
	case v, open := <-ch:
		if open {
			return v, true
		}
		return v, false
	default:
		return v, false
	}
}

func GoAndCollect[T any](res chan T, fs ...func() (T, error)) []T {
	ch := make(chan T)

	var wg sync.WaitGroup
	results := []T{}
	for _, f := range fs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err := f()
			if err != nil {
				ch <- r
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		results = append(results, r)
	}
	return results
}

var (
	TaskKilled   error = errors.New("Task killed")
	TaskCanceled error = errors.New("Task canceled")
)

type Task[T any] struct {
	f    func() (T, error)
	Res  T
	Err  error
	Done chan struct{}
}

func (t *Task[T]) Cancel() {
	stopWithError(t, TaskCanceled)
}

type TaskQueue[T any] interface {
	Push(func() (T, error)) *Task[T]
	Kill() int
}

type taskQueue[T any] struct {
	work chan *Task[T]
	done chan struct{}
}

func (s *taskQueue[T]) Push(f func() (T, error)) *Task[T] {
	var res T
	ch := make(chan struct{})
	t := &Task[T]{
		f,
		res,
		nil,
		ch,
	}
	select {
	case <-s.done:
		close(t.Done)
	default:
		go func() {
			s.enqueue(t)
		}()
	}
	return t
}

func (s *taskQueue[T]) enqueue(t *Task[T]) {
	select {
	case s.work <- t:
	case <-s.done:
		close(t.Done)
	}
}

func (s *taskQueue[T]) Kill() int {
	close(s.done)
	close(s.work)
	return len(s.work)
}

func NewQueue[T any]() TaskQueue[T] {
	work := make(chan *Task[T], 1000)
	workers := 4
	q := &taskQueue[T]{
		work,
		make(chan struct{}),
	}
	for range workers {
		go func() {
			for {
				select {
				case t, ok := <-work:
					if ok {
						select {
						case <-t.Done:
							continue
						default:
							res, err := t.f()
							t.Res, t.Err = res, err
							select {
							case <-t.Done:
							default:
								close(t.Done)
							}
						}
					} else {
						return
					}
				case <-q.done:
					for t := range work {
						stopWithError(t, TaskKilled)
					}
					return
				}
			}
		}()
	}
	return q
}

func stopWithError[T any](t *Task[T], err error) {
	select {
	case <-t.Done:
	default:
		t.Err = err
		close(t.Done)
	}
}
