package str

import (
	"bytes"
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
	return out, true
}

func ToSnakeBuilder(camel string) (string, bool) {

	if len(camel) == 0 {
		return "", false
	}
	out := strings.Builder{}
	out.Grow(len(camel))
	written := 0
	queue := bytes.Buffer{}
	for _, r := range camel {
		if unicode.IsUpper(rune(r)) {
			// Push caps to queue, append to output later
			queue.WriteRune(unicode.ToLower(r))
			continue
		}
		if queue.Len() <= 0 {
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
		if queue.Len() > 1 {
			// case 'CONSTANTVar' -> 'constant_var'

			n, _ := out.Write(queue.Next(queue.Len() - 1))
			written += n
			n, _ = out.WriteRune('_')
			written += n
			queue.WriteTo(&out)
			written += 1

		} else {
			// case 'Var' -> 'var'
			n, _ := out.Write(queue.Next(1))
			written += n
		}
		queue.Reset()
		n, _ := out.WriteRune(r)
		written += n
	}
	return out.String(), true
}
