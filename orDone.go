package main

import (
	"fmt"
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
