package patterns

import "sync"

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

type Task[T any] struct {
	f    func() (T, error)
	Res  T
	Err  error
	Done chan struct{}
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
	go func() {
		select {
		case s.work <- t:
		case <-s.done:
		}
	}()
	return t
}

func (s *taskQueue[T]) Kill() int {
	close(s.done)
	return len(s.work)
}

func NewQueue[T any]() TaskQueue[T] {
	work := make(chan *Task[T])
	workers := 4
	q := &taskQueue[T]{
		work,
		make(chan struct{}),
	}
	for range workers {
		go func() {
			for {
				select {
				case t := <-work:
					res, err := t.f()
					t.Res, t.Err = res, err
					close(t.Done)
				case <-q.done:
					return
				}
			}
		}()
	}
	return q
}
