package patterns

import (
	"errors"
	"fmt"
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
	once sync.Once
	mu   sync.Mutex
	done chan struct{}
}

func (t *Task[T]) Done() <-chan struct{} {
	return t.done
}

func (t *Task[T]) Complete(result T, err error) {
	t.once.Do(func() {
		t.mu.Lock()
		defer t.mu.Unlock()
		t.Res = result
		t.Err = err
		close(t.done)
	})
}

func (t *Task[T]) Cancel() {
	t.Complete(t.Res, TaskCanceled)
}

func (t *Task[T]) CancelWith(err error) {
	t.Complete(t.Res, err)
}

type TaskQueue[T any] interface {
	Push(func() (T, error)) *Task[T]
	Kill() int
}

type taskQueue[T any] struct {
	work chan *Task[T]
	done chan struct{}
	wg   sync.WaitGroup
}

func (tq *taskQueue[T]) Push(f func() (T, error)) *Task[T] {
	var res T
	ch := make(chan struct{})
	t := &Task[T]{
		f,
		res,
		nil,
		sync.Once{},
		sync.Mutex{},
		ch,
	}
	select {
	case <-tq.done:
		t.CancelWith(TaskKilled)
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
			t.CancelWith(TaskKilled)
		}
	}()
}

func (tq *taskQueue[T]) Kill() int {
	close(tq.done)
	close(tq.work)
	tq.wg.Wait()
	return len(tq.work)
}

func NewQueue[T any](workers int) TaskQueue[T] {
	work := make(chan *Task[T])
	q := &taskQueue[T]{
		work: work,
		done: make(chan struct{}),
	}
	q.wg.Add(workers)
	for range workers {
		go func() {
			defer q.wg.Done()
			for {
				select {
				case t, ok := <-work:
					if ok {
						select {
						case <-t.Done():
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
		select {
		case <-t.Done():
		default:
			t.Complete(t.Res, TaskKilled)
		}
	}
}

func runTask[T any](t *Task[T]) {
	// TODO: panic/recover with named err return?
	select {
	case <-t.Done():
	default:
		res, err := wrapWithRecover(t.f)
		t.Complete(res, err)
	}
}

var ErrTaskPanic error = errors.New("Task panic")

func wrapWithRecover[T any](f func() (T, error)) (res T, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v: %w", r, ErrTaskPanic)
		}
	}()
	res, err = f()
	return
}
