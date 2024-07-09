package main

import (
	"fmt"
	_ "playground/blog"
	"playground/options"
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

	setup := options.New().Configure(options.WithA(3), options.WithB("a"))
	fmt.Println(setup)
	setup = options.New().Configure(options.WithDefaults())
	fmt.Println(setup)

}
