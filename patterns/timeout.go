package patterns

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("timed out")

type SomeData struct {
	numbers []int
}

type someDataResult struct {
	result *SomeData
	err    error
}

func GetData() (*SomeData, error) {
	result := make(chan someDataResult, 1)
	go func() {
		result <- doGetData()
	}()

	select {
	case result := <-result:
		return result.result, result.err
	case <-time.After(3 * time.Second):
		return nil, ErrTimeout
	}
}

func doGetData() someDataResult {
	res, err := getSomeThingFromSomeWhere(2)
	return someDataResult{
		result: &res,
		err:    err,
	}
}

func getSomeThingFromSomeWhere(load int64) (SomeData, error) {
	duration := time.Duration(load * int64(time.Second))
	time.Sleep(duration)
	return SomeData{
		numbers: []int{1, 2, 3},
	}, nil
}
