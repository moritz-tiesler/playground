package str

import (
	"testing"
)

func BenchmarkToSnakeStringAppend(b *testing.B) {
	s := "MyFancyBÄSICVar"
	e := "my_fancy_bäsic_var"
	for b.Loop() {
		sn, _ := ToSnake(s)
		if sn != e {
			b.Fatalf("uh oh")
		}
	}
}

func BenchmarkToSnakeStringBuilder(b *testing.B) {
	s := "MyFancyBÄSICVar"
	e := "my_fancy_bäsic_var"
	for b.Loop() {
		sn, _ := ToSnakeBuilder(s)
		if sn != e {
			b.Fatalf("expected=%s, got=%s", e, sn)
		}
	}
}
