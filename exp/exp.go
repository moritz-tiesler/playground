package exp

import (
	"fmt"
	"reflect"
)

func ExecCallBack(cb func(int, int) int, arg int) {
	cb(arg, arg)
}

type Mock struct {
	called int
	copy   reflect.Value
	Args   [][]reflect.Value
}

func (m *Mock) ArgsToSlice() {
	for _, args := range m.Args {
		for _, arg := range args {
			fmt.Printf("%d\n", arg.Int())
		}
	}
}

func MakeMock(fptr any) *Mock {
	m := &Mock{}

	funcValue := reflect.ValueOf(fptr).Elem()
	copy := reflect.New(funcValue.Type()).Elem() // Create a value of the correct type
	copy.Set(funcValue)

	wrapCall := func(in []reflect.Value) []reflect.Value {
		m.Args = append(m.Args, in)
		res := copy.Call(in)
		m.called++
		return res
	}

	makeMock := func(fptr any) {
		fn := reflect.ValueOf(fptr).Elem()
		v := reflect.MakeFunc(fn.Type(), wrapCall)
		fn.Set(v)
	}
	makeMock(fptr)

	return m
}

func Run() {
	mockFunc := func(int, int) int {
		fmt.Println("Mocke executed")
		return 3
	}
	bruh := MakeMock(&mockFunc)
	ExecCallBack(mockFunc, 3)
	ExecCallBack(mockFunc, 3)
	fmt.Println(bruh.called)
	bruh.ArgsToSlice()
}
