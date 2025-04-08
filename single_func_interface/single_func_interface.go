package singlefuncinterface

import (
	"fmt"
	"strings"
)

type Shouter interface {
	Shout(string) string
}

type CountingShouter struct {
	count int
}

func (cs *CountingShouter) Shout(s string) string {
	cs.count++
	return strings.ToUpper(s)
}

type ShouterFunc func(string) string

func (sf ShouterFunc) Shout(s string) string {
	return sf(s)
}

func WithExclamationMarks(n int) ShouterFunc {
	return func(s string) string {
		for range n {
			s += "!"
		}
		return s
	}
}

func AllCaps() ShouterFunc {
	return func(s string) string {
		return strings.ToUpper(s)
	}
}

func ShoutHi(ss ...Shouter) {
	phrase := "hi"
	for _, s := range ss {
		phrase = s.Shout(phrase)
	}
	fmt.Println(phrase)
}

func Run() {
	ShoutHi(ShouterFunc(func(s string) string {
		return s + "!!!"
	}))

	shouter := CountingShouter{}
	ShoutHi(&shouter)

	ShoutHi(AllCaps(), WithExclamationMarks(4))
}
