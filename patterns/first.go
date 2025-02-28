package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	data any
	err  error
}

func Service(url string, queryString string) Result {
	return Result{
		data: fmt.Sprintf("query=%s done, %s", queryString, url),
		err:  nil,
	}
}

func DuckDuckGo() func(string) Result {
	return func(query string) Result {
		return Service("duckduckgo", query)
	}
}

func Yahoo() func(string) Result {
	return func(query string) Result {
		return Service("yahoo", query)
	}
}

func Google() func(string) Result {
	return func(query string) Result {
		return Service("google", query)
	}
}

func First(queryString string, services ...func(string) Result) Result {
	res := make(chan Result)
	for _, s := range services {
		f := func() Result {
			time.Sleep(time.Duration(rand.Intn(2)*1000) * time.Millisecond)
			return s(queryString)
		}
		go func() {
			res <- f()
		}()

	}
	return <-res
}
