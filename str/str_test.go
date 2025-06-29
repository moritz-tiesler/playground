package str

import (
	"reflect"
	"strings"
	"testing"
)

var (
	s = "MyFancyBÄSICVarÉ"
	e = "my_fancy_bäsic_var_é"
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

func TestStrSplitFunc(t *testing.T) {
	v := "aaa-bbb-ccc-"
	chunks := SplitFunc(v, func(r rune) bool {
		return r == '-'
	})

	expected := strings.Split(v, "-")
	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("expected=%v got=%v", expected, chunks)
	}

	v = "aaa-bbb-ccc-"
	chunks = SplitFunc(v, func(r rune) bool {
		return r == '_'
	})

	expected = strings.Split(v, "_")
	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("expected=%v got=%v", expected, chunks)
	}
}
func TestStrSplitAfterFunc(t *testing.T) {
	v := "aaa-bbb-ccc-"
	chunks := SplitAfterFunc(v, func(r rune) bool {
		return r == '-'
	})

	expected := strings.SplitAfter(v, "-")
	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("expected=%v got=%v", expected, chunks)
	}

	v = "aaa-bbb-ccc-"
	chunks = SplitAfterFunc(v, func(r rune) bool {
		return r == '_'
	})

	expected = strings.SplitAfter(v, "_")
	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("expected=%v got=%v", expected, chunks)
	}
}
func TestStrSplitBeforeFunc(t *testing.T) {
	v := "aaa-bbb-ccc-"
	chunks := SplitBeforeFunc(v, func(r rune) bool {
		return r == '-'
	})

	expected := []string{"aaa", "-bbb", "-ccc", "-"}
	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("expected=%v got=%v", expected, chunks)
	}

	v = "aaa-bbb-ccc-"
	chunks = SplitBeforeFunc(v, func(r rune) bool {
		return r == '_'
	})

	expected = []string{v}
	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("expected=%v got=%v", expected, chunks)
	}
}

func TestStrSplitCamel(t *testing.T) {
	v := "MyCoolVar"
	chunks := SplitCamel(v)

	expected := []string{"My", "Cool", "Var"}
	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("expected=%v got=%v", expected, chunks)
	}
}

func TestToSnakePrimitives(t *testing.T) {
	v := "MyCoolVarVa"
	snake, _ := ToSnakePrimitives(v)
	expected := "my_cool_var_va"
	if !reflect.DeepEqual(snake, expected) {
		t.Errorf("expected=%v got=%v", expected, snake)
	}
}

func BenchmarkToSnakePrimitives(b *testing.B) {
	s := "MyFancyBasicVar"
	e := "my_fancy_basic_var"
	for b.Loop() {
		sn, _ := ToSnake(s)
		if sn != e {
			b.Fatalf("expected=%s, got=%s", e, sn)
		}
	}
}
