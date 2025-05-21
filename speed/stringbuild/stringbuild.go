package stringbuild

import (
	"strings"
)

func WithPlus(stringList []string) string {
	concat := ""
	for _, s := range stringList {
		concat += s
	}
	return concat
}

func WithBuilder(stringList []string) string {
	builder := strings.Builder{}
	size := 0
	for _, s := range stringList {
		size += len(s)
	}
	builder.Grow(size)
	for _, s := range stringList {
		builder.WriteString(s)
	}
	return builder.String()
}

func RunConcat(stringList []string, fun func([]string) string) string {
	return fun(stringList)
}
