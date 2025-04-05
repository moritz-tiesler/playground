package nuketheram

import (
	"sync"
	"time"
)

func Launch(n int) {
	var wg sync.WaitGroup
	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			WaitForSignal()
		}()
	}

	wg.Wait()
}

func WaitForSignal() {
	time.Sleep(7 * time.Second)
}
