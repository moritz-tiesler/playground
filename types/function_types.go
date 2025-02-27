package types

type myFunc func(int) int

var lookup map[int]myFunc = make(map[int]myFunc)

type FuncContainer struct{}

func (fc FuncContainer) containerFunc(arg int) int {
	return arg
}

func Run() {
	a_func := func(arg int) int {
		return arg
	}

	cont := FuncContainer{}
	lookup[1] = a_func
	lookup[2] = cont.containerFunc
}
