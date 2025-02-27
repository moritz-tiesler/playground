package patterns

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func RunOrDone() {
	done := make(chan struct{})

	wg.Add(1)
	go orDone(doWork1, done)
	wg.Add(1)
	go orDone(doWork2, done)

	go func() {
		defer time.Sleep(time.Second * 2)
		close(done)
	}()

	wg.Wait()
}

func doWork1() {
	defer wg.Done()
	fmt.Println("Finished doWork1")
}

func doWork2() {
	defer wg.Done()
	fmt.Println("Finished doWork2")
}

func orDone(f func(), done <-chan struct{}) {
	go func() {
		for {
			select {
			case <-done:
				return

			default:
				select {
				case <-done:
					return
				default:
					go f()
				}
			}
		}
	}()
}

func RunSemaphore(n int) {
	fmt.Printf("starting sem with work=%d\n", n)
	work := func(id int) {
		extra := time.Duration(rand.Intn(20)*100) * time.Millisecond
		time.Sleep(time.Millisecond*1000 + extra)
		fmt.Printf("id=%d done\n", id)
	}

	sem := make(chan struct{}, 10)
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go func() {
			sem <- struct{}{}
			work(i)
			<-sem
			wg.Done()
		}()
	}
	wg.Wait()
}
