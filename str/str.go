package str

import (
	"strings"
	"unicode"
)

func ToSnake(camel string) (string, bool) {

	if len(camel) == 0 {
		return "", false
	}
	out := ""
	queue := ""
	for _, r := range camel {
		if unicode.IsUpper(rune(r)) {
			// Push caps to queue, append to output later
			queue += string(unicode.ToLower(rune(r)))
			continue
		}
		if len(queue) <= 0 {
			// no caps letters to write
			out += string(r)
			continue
		}

		if len(out) > 0 {
			// avoid writing '_' to front of output
			out += "_"
		}
		// convert queued caps and write to output
		if len(queue) > 1 {
			// case 'CONSTANTVar' -> 'constant_var'
			out += queue[:len(queue)-1]
			out += "_"
			out += string(queue[len(queue)-1])
		} else {
			// case 'Var' -> 'var'
			out += queue
		}
		queue = ""
		out += string(r)
	}
	if len(queue) > 0 {
		out += "_"
		out += queue
	}
	return out, true
}

func ToSnakeBuilder(camel string) (string, bool) {

	if len(camel) == 0 {
		return "", false
	}
	out := strings.Builder{}
	out.Grow(len(camel))
	written := 0
	queue := make([]rune, len(camel))
	runesQueued := 0
	for _, r := range camel {
		if unicode.IsUpper(rune(r)) {
			// Push caps to queue, append to output later
			queue[runesQueued] = unicode.ToLower(r)
			runesQueued++
			continue
		}
		if runesQueued <= 0 {
			// no caps letters to write
			n, _ := out.WriteRune(r)
			written += n
			continue
		}

		if written > 0 {
			// avoid writing '_' to front of output
			n, _ := out.WriteRune('_')
			written += n
		}
		// convert queued caps and write to output
		if runesQueued > 1 {
			// case 'CONSTANTVar' -> 'constant_var'
			i := 0
			for ; i < runesQueued-1; i++ {
				n, _ := out.WriteRune(queue[i])
				written += n
			}
			n, _ := out.WriteRune('_')
			written += n
			out.WriteRune(queue[runesQueued-1])
			runesQueued = 0
			written += 1

		} else {
			// case 'Var' -> 'var'
			n, _ := out.WriteRune(queue[runesQueued-1])
			runesQueued = 0
			written += n
		}
		n, _ := out.WriteRune(r)
		written += n
	}
	if runesQueued > 0 {
		out.WriteRune('_')
		out.WriteRune(queue[runesQueued-1])
	}
	return out.String(), true
}
