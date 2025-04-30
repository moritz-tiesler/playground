package main

import (
	"fmt"
	_ "fmt"
	"iter"
	_ "playground/blog"
	_ "playground/encoding"
	_ "playground/exp"
	_ "playground/middleware"
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

	// for i := range patterns.ThreeTimes {
	// 	if i == 2 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// for i := range patterns.Range(2, 10) {
	// 	fmt.Println(i)
	// }

	// exp.Run()

	// mux := middleware.NewServeMux()

	// fmt.Println("Running server")

	// err := http.ListenAndServe(":8080", middleware.LogRequestMiddleware(middleware.SecureHeadersMiddleware(mux)))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// CallNamer(&NPC{name: "Bob"})
	// CallNamer(MakeNamer(func() string { return "Bob" }))

	// input := make(chan string)
	// stop := make(chan struct{})
	// resume := make(chan struct{})

	// go func() {
	// 	for {
	// 		select {
	// 		case in := <-input:
	// 			if in == "s" {
	// 				fmt.Println("received stop")
	// 				stop <- struct{}{}
	// 			}
	// 			if in == "c" {
	// 				fmt.Println("received continue")
	// 				resume <- struct{}{}
	// 			}
	// 		}
	// 	}
	// }()

	// go func() {
	// 	ticker := time.NewTicker(2000 * time.Millisecond)
	// 	for {
	// 		select {
	// 		case <-stop:
	// 			ticker.Stop()
	// 			<-resume
	// 			ticker = time.NewTicker(2000 * time.Millisecond)
	// 		case <-ticker.C:
	// 			fmt.Println("running")
	// 		}
	// 	}
	// }()

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	input <- strings.TrimSpace(scanner.Text())
	// }

	// go func() {
	// 	halt := make(chan struct{})
	// 	for {
	// 		select {
	// 		case <-stop:
	// 			<-halt
	// 		}
	// 	}
	// }()
	// nuketheram.Launch(10_000_000)
	// s := set.New(patterns.ThreeTimes)
	// for i := range s.Items() {
	// 	fmt.Println(i)
	// }
	// fmt.Println(s.Len())

	// singlefuncinterface.Run()
	seq1 := iter.Seq[int](func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i) {
				return
			}
		}
	})

	// seq2 := iter.Seq[int](func(yield func(int) bool) {
	// 	for i := 10; i < 13; i++ {
	// 		if !yield(i) {
	// 			return
	// 		}
	// 	}
	// })

	// take := patterns.TakeWhile(seq1, func(i int) bool { return i < 5 })
	// for i := range take {
	// 	fmt.Println(i)
	// }

	t := 0
	for v := range patterns.Cycle(seq1) {
		if t == 30 {
			break
		}
		fmt.Println(v)
		t++
	}

	v := patterns.NewSingleTon[int]()
	*v = 12
	fmt.Println(*v)

	v2 := patterns.NewSingleTon[int]()
	fmt.Println(*v2)

	// patterns.ForEach(seq1, func(i int) { fmt.Println(i * 2) })

}

type Namer interface {
	Name() string
}

type NPC struct {
	name string
}

func (n *NPC) Name() string { return n.name }

type NamerFunc func() string

func (n NamerFunc) Name() string {
	return n()
}

func CallNamer(n Namer) {
	fmt.Println(n.Name())
}

func MakeNamer(f NamerFunc) Namer {
	return NamerFunc(f)
}
