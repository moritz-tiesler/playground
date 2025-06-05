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
	done chan struct{}
}

func (t *Task[T]) Done() <-chan struct{} {
	return t.done
}

func (t *Task[T]) Cancel() {
	t.stopWithError(TaskCanceled)
}

func (t *Task[T]) CancelWith(err error) {
	t.stopWithError(err)
}

func (t *Task[T]) stopWithError(err error) {
	select {
	case <-t.done:
	default:
		t.Err = err
		close(t.done)
	}
}

type TaskQueue[T any] interface {
	Push(func() (T, error)) *Task[T]
	Kill() int
}

type taskQueue[T any] struct {
	work chan *Task[T]
	done chan struct{}
}

func (tq *taskQueue[T]) Push(f func() (T, error)) *Task[T] {
	var res T
	ch := make(chan struct{})
	t := &Task[T]{
		f,
		res,
		nil,
		ch,
	}
	select {
	case <-tq.done:
		close(t.done)
	case <-t.done:
	default:
		tq.enqueue(t)
	}
	return t
}

func (tq *taskQueue[T]) enqueue(t *Task[T]) {
	go func() {
		select {
		case tq.work <- t:
		case <-tq.done:
			t.stopWithError(TaskKilled)
		}
	}()
}

func (tq *taskQueue[T]) Kill() int {
	close(tq.done)
	close(tq.work)
	return len(tq.work)
}

func NewQueue[T any](workers int) TaskQueue[T] {
	work := make(chan *Task[T])
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
						case <-t.done:
							continue
						default:
							runTask(t)
						}
					} else {
						return
					}
				case <-q.done:
					q.cancelWork()
					return
				}
			}
		}()
	}
	return q
}

func (tq *taskQueue[T]) cancelWork() {
	for t := range tq.work {
		t.stopWithError(TaskKilled)
	}
}

func runTask[T any](t *Task[T]) {
	res, err := t.f()
	t.Res, t.Err = res, err
	select {
	case <-t.done:
	default:
		close(t.done)
	}
}
