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
	t.complete(result, err)
}

func (t *Task[T]) complete(result T, err error) {
	t.once.Do(func() {
		t.mu.Lock()
		defer t.mu.Unlock()
		t.Res = result
		t.Err = err
		close(t.done)
	})
}

func (t *Task[T]) Cancel() {
	var zero T
	t.complete(zero, TaskCanceled)
}

func (t *Task[T]) CancelWith(err error) {
	var zero T
	t.complete(zero, err)
}

type TaskQueue[T any] interface {
	Push(func() (T, error)) *Task[T]
	Kill() int
	Start()
}

type opts[T any] struct {
	panicDefer func(any, *T, *error)
	numWorkers int
	queueBuf   int
}

type option[T any] func(*opts[T])

func WithWorkers[T any](n int) option[T] {
	return func(o *opts[T]) {
		o.numWorkers = n
	}
}

func WithQueueBuffer[T any](n int) option[T] {
	return func(o *opts[T]) {
		o.queueBuf = n
	}
}

func WithPanicDefer[T any](f func(any, *T, *error)) option[T] {
	return func(o *opts[T]) {
		o.panicDefer = f
	}
}

func defaultOpts[T any]() *opts[T] {
	return &opts[T]{
		nil,
		1,
		0,
	}
}

type taskQueue[T any] struct {
	work    chan *Task[T]
	bouncer chan *Task[T]
	done    chan struct{}
	wg      sync.WaitGroup
	opts    *opts[T]
}

// TODO: return true if task was pushed successfully
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
	tq.tryEnqueue(t)
	return t
}

func (tq *taskQueue[T]) tryEnqueue(t *Task[T]) {
	go func() {
		select {
		case tq.bouncer <- t:
		case <-tq.done:
			t.CancelWith(TaskKilled)
		}
	}()
}

func (tq *taskQueue[T]) Kill() int {
	close(tq.done)
	tq.wg.Wait()
	return len(tq.work)
}

func NewQueue[T any](options ...option[T]) TaskQueue[T] {
	opts := defaultOpts[T]()
	for _, o := range options {
		o(opts)
	}
	work := make(chan *Task[T], opts.queueBuf)
	bouncer := make(chan *Task[T])
	q := &taskQueue[T]{
		work:    work,
		done:    make(chan struct{}),
		opts:    opts,
		bouncer: bouncer,
	}

	return q
}

func (tq *taskQueue[T]) Start() {
	tq.wg.Add(tq.opts.numWorkers)
	for range tq.opts.numWorkers {
		go func() {
			defer tq.wg.Done()
			for t := range tq.work {
				tq.runTask(t)
			}
		}()
	}
	go func() {
		for {
			select {
			case guest := <-tq.bouncer:
				tq.work <- guest
			case <-tq.done:
				close(tq.work)
				tq.cancelWork(tq.work)
				return
			}
		}
	}()
}

func (tq *taskQueue[T]) cancelWork(ch chan *Task[T]) {
	for t := range ch {
		select {
		case <-t.Done():
		default:
			t.CancelWith(TaskKilled)
		}
	}
}

func (tq *taskQueue[T]) runTask(t *Task[T]) {
	select {
	case <-t.Done():
	default:
		res, err := tq.wrapWithRecover(t.f)
		t.Complete(res, err)
	}
}

var ErrTaskPanic error = errors.New("Task panic")

func WrapPanic[T any](rec any, _ *T, err *error) {
	*err = fmt.Errorf("%v: %w", rec, ErrTaskPanic)
}

func (tq *taskQueue[T]) wrapWithRecover(f func() (T, error)) (res T, err error) {
	if tq.opts.panicDefer != nil {
		defer func() {
			if r := recover(); r != nil {
				tq.opts.panicDefer(r, &res, &err)
			}
		}()
	}
	res, err = f()
	return
}
