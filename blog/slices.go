package blog

import (
	"bytes"
	"fmt"
)

// https://go.dev/blog/slices

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

var buffer [256]byte

func Exec() {
	slice := buffer[100:150]
	fmt.Println("Before: len(slice) =", len(slice))
	PtrSubtractOneFromLength(&slice)
	fmt.Println("After:  len(slice) =", len(slice))
}

type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func Exec2() {
	pathName := path("/usr/bin/tso")
	pathName.TruncateAtFinalSlash()
	fmt.Printf("%s\n", pathName)
}

func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

func Exec3() {
	pathName := path("/usr/bin/tso")
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)
}

func (p *path) ToUpperWithPtr() {
	for i, b := range *p {
		if 'a' <= b && b <= 'z' {
			(*p)[i] = b + 'A' - 'a'
		}
	}
}

func Exec4() {
	pathName := path("/usr/bin/tso")
	pathName.ToUpperWithPtr()
	fmt.Printf("%s\n", pathName)
}
