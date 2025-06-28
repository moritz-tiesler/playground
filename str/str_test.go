package str

import (
	"testing"
)

var (
	s = "MyFancyBÄSICVar"
	e = "my_fancy_bäsic_var"
)

func BenchmarkToSnakeStringAppend(b *testing.B) {
	for b.Loop() {
		sn, _ := ToSnake(s)
		if sn != e {
			b.Fatalf("expected=%s, got=%s", e, sn)
		}
	}
}

func BenchmarkToSnakeStringBuilder(b *testing.B) {
	for b.Loop() {
		sn, _ := ToSnakeBuilder(s)
		if sn != e {
			b.Fatalf("expected=%s, got=%s", e, sn)
		}
	}
}
