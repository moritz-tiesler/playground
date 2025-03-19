package main

import (
	"fmt"
	_ "fmt"
	_ "playground/blog"
	_ "playground/encoding"
	_ "playground/exp"
	_ "playground/options"
	"playground/patterns"
	_ "playground/patterns"
	_ "playground/types"
)

func main() {
	//rawData1 := []string{}
	//rawData2 := []string{"a"}

	//parsed1, err := generics.NewNonEmptySlice(rawData1)
	//fmt.Printf("%v, %v\n", parsed1, err)
	//parsed2, err := generics.NewNonEmptySlice(rawData2)
	//fmt.Printf("%v, %v\n", parsed2, err)

	//nums, err := patterns.GetData()
	//if errors.Is(err, patterns.ErrTimeout) {
	//fmt.Println("we timed out")
	//}
	//fmt.Println(nums, err)

	// setup := options.New().Configure(options.WithA(3), options.WithB("a"))
	// fmt.Println(setup)
	// setup = options.New().Configure(options.WithDefaults())
	// fmt.Println(setup)
	// types.Run()
	// patterns.RunSemaphore(100)

	// fmt.Println(
	// 	patterns.First("count = 10", patterns.Yahoo(), patterns.DuckDuckGo(), patterns.Google()),
	// )

	// tString := "aaabcaabdeffabfabuab"
	// p, c := encoding.MostCommonPair(tString)
	// fmt.Printf("pair=%s, occured=%d\n", p, c)
	// aFunction := func(i int) string {
	// 	return fmt.Sprintf("Called with %d", i)
	// }

	// callMe := func(f func(i int) string, arg int) string {
	// 	return f(arg)
	// }

	// mock := exp.Fn(aFunction)

	// res := callMe(aFunction, 4)
	// fmt.Println(res)

	// testVal := 0

	for i := range patterns.Range(2, 10) {
		fmt.Println(i)
	}

}
